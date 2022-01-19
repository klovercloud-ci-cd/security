package logic

import (
	v1 "github.com/klovercloud-ci/core/v1"
	"github.com/klovercloud-ci/core/v1/service"
)

type mockRoleService struct {

}

func (m mockRoleService) Store(role v1.Role) error {
	panic("implement me")
}

func (m mockRoleService) Get() []v1.Role {
	panic("implement me")
}

func (m mockRoleService) GetByName(name string) v1.Role {
	panic("implement me")
}

func (m mockRoleService) Delete(name string) error {
	panic("implement me")
}

func (m mockRoleService) Update(name string, permissions []v1.Permission, option v1.RoleUpdateOption) error {
	panic("implement me")
}

func NewMockRoleService() service.Role {
	return &mockRoleService{
	}
}