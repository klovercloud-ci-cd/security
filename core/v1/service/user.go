package service

import v1 "github.com/klovercloud-ci/core/v1"

type User interface {
	Store(user v1.UserRegistrationDto) error
	Get() []v1.User
	GetByID(id string) (v1.User, error)
	Delete(id string) error
	UpdateToken(token, refreshToken, existingToken string) error
}
