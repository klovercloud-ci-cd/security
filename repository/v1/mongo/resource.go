package mongo

import (
	"context"
	"errors"
	v1 "github.com/klovercloud-ci-cd/security/core/v1"
	"github.com/klovercloud-ci-cd/security/core/v1/repository"
	"go.mongodb.org/mongo-driver/bson"
	"log"
	"time"
)

// ResourceCollection collection name
var (
	ResourceCollection = "resourceCollection"
)

type resourceRepository struct {
	manager *dmManager
	timeout time.Duration
}

func (r resourceRepository) Get() []v1.Resource {
	var resources []v1.Resource
	coll := r.manager.Db.Collection(ResourceCollection)
	result, err := coll.Find(r.manager.Ctx, bson.M{})
	if err != nil {
		log.Println(err.Error())
	}
	for result.Next(context.TODO()) {
		elemValue := new(v1.Resource)
		err := result.Decode(elemValue)
		if err != nil {
			log.Println("[ERROR]", err)
			return nil
		}
		resources = append(resources, *elemValue)
	}
	return resources
}

func (r resourceRepository) GetByName(name string) v1.Resource {
	elemValue := new(v1.Resource)
	filter := bson.M{"name": name}
	coll := r.manager.Db.Collection(ResourceCollection)
	result := coll.FindOne(r.manager.Ctx, filter)
	err := result.Decode(elemValue)
	if err != nil {
		log.Println("[ERROR]", err)
		return *elemValue
	}
	return *elemValue
}

func (r resourceRepository) Delete(resourceName string) error {
	coll := r.manager.Db.Collection(ResourceCollection)
	filter := bson.M{"name": resourceName}
	data, err := coll.DeleteOne(r.manager.Ctx, filter)
	if err != nil {
		log.Println("[ERROR]", err)
		return err
	}
	if data.DeletedCount == 0 {
		log.Println("No data found to delete!")
		return errors.New("No data found to delete!")
	}
	return err
}

func (r resourceRepository) Store(resource v1.Resource) error {
	if r.GetByName(resource.Name).Name == "" {
		coll := r.manager.Db.Collection(ResourceCollection)
		_, err := coll.InsertOne(r.manager.Ctx, resource)
		if err != nil {
			log.Println("[ERROR] Insert document:", err.Error())
		}
	}
	return nil
}

// NewResourceRepository returns repository.Resource type repository
func NewResourceRepository(timeout int) repository.Resource {
	return &resourceRepository{
		manager: GetDmManager(),
		timeout: time.Duration(timeout) * time.Second,
	}
}
