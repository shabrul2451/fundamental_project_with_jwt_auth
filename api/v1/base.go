package v1

import (
	"awesomeProject/dependency"
	"github.com/labstack/echo/v4"
)

func Rourter(g *echo.Group) {
	UserRouter(g.Group("/users"))
	OauthRouter(g.Group("/oauth"))
}

func UserRouter(g *echo.Group) {
	userApi := NewUserApi(dependency.GetV1UserService())
	g.POST("", userApi.Save)
	g.GET("", userApi.Get)
}

func OauthRouter(g *echo.Group) {
	oauthApi := NewOauthApi(dependency.GetV1JwtService(), dependency.GetV1UserService())
	g.POST("/login", oauthApi.Login)
}
