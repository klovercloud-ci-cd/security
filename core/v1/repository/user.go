package repository

import v1 "github.com/klovercloud-ci/core/v1"

type User interface {
	Store(user v1.User) error
	Get() []v1.User
	GetByID(id string) (v1.User, error)
	Delete(id string) error
	GetByEmail(email string) v1.User
	GetByToken(token string) v1.User
	UpdateToken(user v1.User) error
}
