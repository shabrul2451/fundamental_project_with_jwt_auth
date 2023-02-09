package api

import (
	v1 "awesomeProject/api/v1"
	"github.com/labstack/echo/v4"
	"github.com/swaggo/echo-swagger"
	"net/http"
)

func Routes(e *echo.Echo) {
	//Index Page
	e.GET("/", index)

	// Health Page
	e.GET("/health", health)

	//Swagger
	e.GET("/swagger/*", echoSwagger.WrapHandler)
	v1.Rourter(e.Group("/awesome_project/api/v1"))
}

func index(c echo.Context) error {
	return c.String(http.StatusOK, "This is awesomeProject")
}

func health(c echo.Context) error {
	return c.String(http.StatusOK, "awesomeProject is alive")
}
