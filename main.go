package main

import (
	"github.com/klovercloud-ci-cd/security/api"
	"github.com/klovercloud-ci-cd/security/config"
	"github.com/klovercloud-ci-cd/security/core/v1"
	"github.com/klovercloud-ci-cd/security/dependency"
	_ "github.com/klovercloud-ci-cd/security/docs"
	"github.com/klovercloud-ci-cd/security/enums"
	"github.com/labstack/echo/v4/middleware"
	"net/http"
)

// @title Klovercloud-ci-security API
// @description Klovercloud-security API
func main() {
	e := config.New()
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{http.MethodGet, http.MethodPut, http.MethodPost, http.MethodDelete},
	}))
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

func initRoles() {
	permissions := dependency.GetV1PermissionService().Get()
	role := v1.RoleDto{
		Name:        string(enums.ADMIN),
		Permissions: permissions,
	}
	dependency.GetV1RoleService().Store(role)
}

//swag init --parseDependency --parseInternal
