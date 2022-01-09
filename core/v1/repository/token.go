package repository

import v1 "github.com/klovercloud-ci/core/v1"

type Token interface {
	Store(token v1.Token) error
	IsValid(token string) (bool, error)
	Delete(uid string) error
	Update(token string, refreshToken string, existingToken string) error
	GetByToken(token string) v1.Token
	GetByUID(uid string) v1.Token
}
