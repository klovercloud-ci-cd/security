package mongo

import (
	v1 "github.com/klovercloud-ci/core/v1"
	"github.com/klovercloud-ci/core/v1/repository"
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

func (r resourceRepository) Store(resource v1.Resource) error {
	coll := r.manager.Db.Collection(ResourceCollection)
	_, err := coll.InsertOne(r.manager.Ctx, resource)
	if err != nil {
		log.Println("[ERROR] Insert document:", err.Error())
	}
	return nil
}

func NewResourceRepository(timeout int) repository.ResourceRepository {
	return &resourceRepository{
		manager: GetDmManager(),
		timeout: time.Duration(timeout) * time.Second,
	}
}
