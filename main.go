package main

import (
	"github.com/google/uuid"
	"github.com/klovercloud-ci-cd/security/api"
	"github.com/klovercloud-ci-cd/security/config"
	"github.com/klovercloud-ci-cd/security/core/v1"
	"github.com/klovercloud-ci-cd/security/dependency"
	_ "github.com/klovercloud-ci-cd/security/docs"
	"github.com/klovercloud-ci-cd/security/enums"
	"github.com/labstack/echo/v4/middleware"
	"net/http"
	"time"
)

// @title Klovercloud-ci-security API
// @description Klovercloud-security API
func main() {
	e := config.New()
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{http.MethodGet, http.MethodPut, http.MethodPost, http.MethodDelete},
	}))
	initSuperAdmin()
	go initResources()
	initPermissions()
	initRoles()
	api.Routes(e)
	e.Logger.Fatal(e.Start(":" + config.ServerPort))
}

func initResources() {
	resourceService := dependency.GetV1ResourceService()
	for _, each := range config.Resources {
		resourceService.Store(v1.Resource{Name: each})
	}
}

func initPermissions() {
	permissionService := dependency.GetV1PermissionService()
	for _, each := range config.Permissions {
		permissionService.Store(v1.Permission{Name: each})
	}
}

func initUserResourcePermission() v1.UserResourcePermission {
	roleService := dependency.GetV1RoleService()
	resourceService := dependency.GetV1ResourceService()
	userResourcePermissionDto := v1.UserResourcePermission{}
	var resourceWiseRoles []v1.ResourceWiseRoles
	existingResources := resourceService.Get()
	adminRole := roleService.GetByName(string(enums.ADMIN_ROLE))
	for _, each := range existingResources {
		resourceWiseRole := v1.ResourceWiseRoles{
			Name:  each.Name,
			Roles: []v1.Role{{Name: adminRole.Name}},
		}
		resourceWiseRoles = append(resourceWiseRoles, resourceWiseRole)
	}
	userResourcePermissionDto.Resources = resourceWiseRoles

	return userResourcePermissionDto
}

func initSuperAdmin() {
	userService := dependency.GetV1UserService()
	userResourcePermissionDto := initUserResourcePermission()
	if config.Email != "" {
		companyId := uuid.New().String()
		userRegistrationDto := v1.UserRegistrationDto{
			Metadata:           v1.UserMetadata{CompanyId: companyId},
			FirstName:          config.FirstName,
			LastName:           config.LastName,
			Email:              config.Email,
			Phone:              config.PhoneNumber,
			Password:           config.Password,
			AuthType:           enums.AUTH_TYPE(config.AuthType),
			CreatedDate:        time.Now().UTC(),
			UpdatedDate:        time.Now().UTC(),
			Status:             enums.ACTIVE,
			ID:                 uuid.New().String(),
			ResourcePermission: userResourcePermissionDto,
		}
		err := userService.Store(userRegistrationDto)
		if err == nil {
			userService.InitCompany(v1.Company{
				Id:   companyId,
				Name: config.CompanyName,
			})
		}
	}
}

func initRoles() {
	permissions := dependency.GetV1PermissionService().Get()
	role := v1.RoleDto{
		Name:        string(enums.ADMIN_ROLE),
		Permissions: permissions,
	}
	dependency.GetV1RoleService().Store(role)
	role = v1.RoleDto{
		Name: string(enums.USER_ROLE),
		Permissions: []v1.Permission{
			{
				Name: "READ",
			},
		},
	}
	dependency.GetV1RoleService().Store(role)
}

//swag init --parseDependency --parseInternal
