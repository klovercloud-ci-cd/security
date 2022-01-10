package dependency

import (
	"github.com/klovercloud-ci/config"
	"github.com/klovercloud-ci/core/v1/logic"
	"github.com/klovercloud-ci/core/v1/service"
	"github.com/klovercloud-ci/enums"
	"github.com/klovercloud-ci/repository/v1/mongo"
)

func GetV1ResourceService() service.Resource {
	var resource service.Resource
	resource = logic.NewResourceService(mongo.NewResourceRepository(3000))
	return resource
}

func GetV1PermissionService() service.Permission {
	var permission service.Permission
	permission = logic.NewPermissionService(mongo.NewPermissionRepository(3000))
	return permission
}

func GetV1UserService() service.User {
	if config.RunMode==string(enums.PRODCUTION){
		return logic.NewUserService(mongo.NewUserRepository(3000), mongo.NewUserResourcePermissionRepository(3000))
	}
	return logic.NewUserMock()
}

func GetV1JwtService() service.Jwt {
	return  logic.NewJwtService()
}

func GetV1UserResourcePermissionService() service.UserResourcePermission {
	if config.RunMode==string(enums.PRODCUTION) {
		return logic.NewUserResourcePermissionService(mongo.NewUserResourcePermissionRepository(3000))
	}
	return  logic.NewMockUserResourcePermissionService()
}

func GetV1TokenService() service.Token {
	if config.RunMode==string(enums.PRODCUTION) {
		return logic.NewTokenService(mongo.NewTokenRepository(3000), GetV1JwtService())
	}
	return logic.NewTokenMock()
}
