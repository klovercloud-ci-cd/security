package v1

import (
	"encoding/json"
	"github.com/golang-jwt/jwt"
	"github.com/klovercloud-ci/api/common"
	"github.com/klovercloud-ci/config"
	v1 "github.com/klovercloud-ci/core/v1"
	"github.com/klovercloud-ci/core/v1/api"
	"github.com/klovercloud-ci/core/v1/service"
	"github.com/klovercloud-ci/enums"
	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
	"log"
	"strconv"
)

type oauthApi struct {
	userService            service.User
	jwtService             service.Jwt
	userResourcePermission service.UserResourcePermission
	tokenService           service.Token
}

func (o oauthApi) Login(context echo.Context) error {
	if context.QueryParam("grant_type") == "password" {
		return o.handlePasswordGrant(context)
	}else if context.QueryParam("grant_type") == "refresh_token"{
	return o.handleRefreshTokenGrant(context)
	}
	return common.GenerateForbiddenResponse(context, nil, "Please provide a valid grant_type")
}

func  (o oauthApi) handleRefreshTokenGrant(context echo.Context) error{
	refreshTokenDto := new(v1.RefreshTokenDto)
	if err := context.Bind(&refreshTokenDto); err != nil {
		log.Println("Input Error:", err.Error())
		return common.GenerateErrorResponse(context, "[ERROR]: Failed bind payload from context", err.Error())
	}

	if !o.jwtService.IsTokenValid(refreshTokenDto.RefreshToken){
		return common.GenerateForbiddenResponse(context, "[ERROR]: Token is expired!","Please login again to get token!")
	}
	claims := jwt.MapClaims{}
	jwt.ParseWithClaims(refreshTokenDto.RefreshToken, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(config.Publickey), nil
	})
	jsonbody, err := json.Marshal(claims["data"])
	if err != nil {
		log.Println(err)
	}
	usersPermission := v1.UserResourcePermission{}
	if err := json.Unmarshal(jsonbody, &usersPermission); err != nil {
		log.Println(err)
	}

	tokenLifeTime, err := strconv.ParseInt(config.RegularTokenLifetime, 10, 64)
		if err != nil {
			log.Println(err.Error())
			return common.GenerateForbiddenResponse(context, "[ERROR]: failed to read regular token lifetime from env!", err.Error())
		}
	token, refreshToken, err := o.jwtService.GenerateToken(usersPermission.UserId, tokenLifeTime, usersPermission)
	if err != nil {
		log.Println(err.Error())
		return common.GenerateForbiddenResponse(context, "[ERROR]: failed to create token!", err.Error())
	}

	err = o.tokenService.Store(v1.Token{usersPermission.UserId, token, refreshToken, enums.REGULAR_TOKEN})
	if err != nil {
		log.Println(err.Error())
		return common.GenerateForbiddenResponse(context, "[ERROR]: failed to store token!", err.Error())
	}
	return common.GenerateSuccessResponse(context, v1.JWTPayLoad{token, refreshToken}, nil, "")
}

func (o oauthApi) handlePasswordGrant(context echo.Context) error {
	token_type := context.QueryParam("token_type")
	if token_type == "" {
		token_type = string(enums.REGULAR_TOKEN)
	} else if token_type != string(enums.REGULAR_TOKEN) && token_type != string(enums.CTL_TOKEN) {
		return common.GenerateErrorResponse(context, "No valid token token_type provided!", "Please provide a valid token_type!")
	}
	loginDto := new(v1.LoginDto)
	if err := context.Bind(&loginDto); err != nil {
		log.Println("Input Error:", err.Error())
		return common.GenerateErrorResponse(context, "[ERROR]: Failed bind payload from context", err.Error())
	}

	existingUser := o.userService.GetByEmail(loginDto.Email)
	if existingUser.ID == "" {
		return common.GenerateForbiddenResponse(context, "[ERROR]: No User found!", "Please login with actual user email!")
	}
	err := bcrypt.CompareHashAndPassword([]byte(existingUser.Password), []byte(loginDto.Password))
	if err != nil {
		return common.GenerateForbiddenResponse(context, "[ERROR]: Password not matched!", "Please login with due credential!"+err.Error())
	}
	userResourcePermission := o.userResourcePermission.GetByUserID(existingUser.ID)
	userResourcePermission.Metadata.CompanyId=existingUser.Metadata.CompanyId
	var tokenLifeTime int64
	if token_type == string(enums.REGULAR_TOKEN) {
		i, err := strconv.ParseInt(config.RegularTokenLifetime, 10, 64)
		if err != nil {
			log.Println(err.Error())
			return common.GenerateForbiddenResponse(context, "[ERROR]: failed to read regular token lifetime from env!", err.Error())
		}
		tokenLifeTime = i
	} else {
		i, err := strconv.ParseInt(config.CTLTokenLifetime, 10, 64)
		if err != nil {
			log.Println(err.Error())
			return common.GenerateForbiddenResponse(context, "[ERROR]: failed to read ctl token lifetime from env!", err.Error())
		}
		tokenLifeTime = i
	}
	token, refreshToken, err := o.jwtService.GenerateToken(userResourcePermission.UserId, tokenLifeTime, userResourcePermission)
	if err != nil {
		log.Println(err.Error())
		return common.GenerateForbiddenResponse(context, "[ERROR]: failed to create token!", err.Error())
	}

	err = o.tokenService.Store(v1.Token{userResourcePermission.UserId, token, refreshToken, enums.TOKEN_TYPE(token_type)})
	if err != nil {
		log.Println(err.Error())
		return common.GenerateForbiddenResponse(context, "[ERROR]: failed to store token!", err.Error())
	}
	return common.GenerateSuccessResponse(context, v1.JWTPayLoad{token, refreshToken}, nil, "")
}

func NewOauthApi(userService service.User, jwtService service.Jwt,userResourcePermission service.UserResourcePermission,tokenService service.Token) api.Oauth {
	return &oauthApi{
		userService: userService,
		jwtService:  jwtService,
		userResourcePermission:userResourcePermission,
		tokenService: tokenService,
	}
}
