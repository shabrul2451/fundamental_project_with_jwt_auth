package main

import (
	"awesomeProject/api"
	"awesomeProject/config"
	"github.com/labstack/echo/v4/middleware"
	"net/http"
)

func main() {
	e := config.New()

	//handle cors origin
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowHeaders: []string{"*"},
		AllowMethods: []string{http.MethodGet, http.MethodPut, http.MethodPost, http.MethodDelete},
	}))

	api.Routes(e)
	e.Logger.Fatal(e.Start(":" + config.ServerPort))
}
