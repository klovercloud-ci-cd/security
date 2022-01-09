package logic

import (
	v1 "github.com/klovercloud-ci/core/v1"
	"github.com/klovercloud-ci/core/v1/repository"
	"github.com/klovercloud-ci/core/v1/service"
	"github.com/klovercloud-ci/enums"
)

type tokenService struct {
	tokenRepo repository.Token
	jwtService service.Jwt
}

func (t tokenService) Store(token v1.Token) error {
	if token.Type == enums.REGULAR_TOKEN {
		oldToken := t.tokenRepo.GetByUID(token.Uid)

		oldToken.Token = token.Token
		oldToken.RefreshToken = token.RefreshToken
		return t.tokenRepo.Store(oldToken)
	}
	return t.tokenRepo.Store(token)
}

func (t tokenService) IsValid(token string) bool {
	res,_:=t.jwtService.IsTokenValid(token)
	if res{
		if t.tokenRepo.GetByToken(token).Uid==""{
			return false
		}
		return true
	}
	return false
}

func (t tokenService) Delete(uid string) error {
	return t.tokenRepo.Delete(uid)
}

func (t tokenService) Update(token string, refreshToken string, existingToken string) error {
	return t.tokenRepo.Update(token, refreshToken, existingToken)
}

func NewTokenService(tokenRepo repository.Token,jwtService service.Jwt) service.Token {
	return &tokenService{
		tokenRepo: tokenRepo,
		jwtService: jwtService,
	}
}