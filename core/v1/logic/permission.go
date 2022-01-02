package logic

import (
	v1 "github.com/klovercloud-ci/core/v1"
	"github.com/klovercloud-ci/core/v1/repository"
	"github.com/klovercloud-ci/core/v1/service"
)

type permissionService struct {
	repo repository.Permission
}

func (p permissionService) Store(permission v1.Permission) error {
	//TODO implement me
	panic("implement me")
}

func (p permissionService) Get() ([]v1.Permission, error) {
	//TODO implement me
	panic("implement me")
}

func (p permissionService) Delete(permissionName string) error {
	//TODO implement me
	panic("implement me")
}

func NewPermissionService(repo repository.Permission) service.Permission {
	return &permissionService{
		repo: repo,
	}
}
