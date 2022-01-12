package v1

import (
	"github.com/klovercloud-ci/dependency"
	"github.com/labstack/echo/v4"
)

func Router(g *echo.Group) {
	RoleRouter(g.Group("/roles"))
	ResourceRouter(g.Group("/resources"))
	PermissionRouter(g.Group("/permissions"))
	UserRouter(g.Group("/users"))
	OauthRouter(g.Group("/oauth"))
}

func ResourceRouter(g *echo.Group) {
	resourceApi := NewResourceApi(dependency.GetV1ResourceService())
	g.POST("", resourceApi.Store)
	g.GET("", resourceApi.Get)
	g.GET("/:resourceName", resourceApi.GetByName)
	g.DELETE("/:resourceName", resourceApi.Delete)
}

func PermissionRouter(g *echo.Group) {
	permissionApi := NewPermissionApi(dependency.GetV1PermissionService())
	g.POST("", permissionApi.Store)
	g.GET("", permissionApi.Get)
	g.DELETE("", permissionApi.Delete)
}

func RoleRouter(g *echo.Group) {
	roleApi := NewRoleApi(dependency.GetV1RoleService())
	g.POST("", roleApi.Store)
	g.GET("", roleApi.Get)
	g.GET("/:roleName", roleApi.GetByName)
	g.DELETE("/:roleName", roleApi.Delete)
	g.POST("/:roleName", roleApi.Update)
}

func UserRouter(g *echo.Group) {
	userApi := NewUserApi(dependency.GetV1UserService(), dependency.GetV1UserResourcePermissionService(),dependency.GetV1OtpService(),dependency.GetV1JwtService())
	g.POST("", userApi.Store)
	g.GET("", userApi.Get)
	g.GET("/:id", userApi.GetByID)
	g.DELETE("/:id", userApi.Delete)
	g.PUT("", userApi.Update)
	g.PUT("/:id/userResourcePermission", userApi.UpdateUserResourcePermission)
}

func OauthRouter(g *echo.Group) {
	oauthApi := NewOauthApi(dependency.GetV1UserService(),dependency.GetV1JwtService(),dependency.GetV1UserResourcePermissionService(),dependency.GetV1TokenService())
	g.POST("/login", oauthApi.Login)
}
