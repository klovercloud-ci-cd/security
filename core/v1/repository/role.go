package repository

import v1 "github.com/klovercloud-ci/core/v1"

type Role interface {
	Store(role v1.Role) error
	Get() ([]v1.Role, int64)
	GetByName(name string) (v1.Role, error)
	Delete(name string) error
	Update(name string, permissions []v1.Permission) error
	AppendPermissions(name string, permissions []v1.Permission) error
	RemovePermissions(name string, permission []v1.Permission) error
}
