package main

import (
	"awesomeProject/api"
	"awesomeProject/config"
	_ "awesomeProject/docs"
	"github.com/labstack/echo/v4/middleware"
	"net/http"
)

func main() {
	// @title Swagger API Documentation
	// @version v1
	// @description This is a sample server with oauth.
	// @termsOfService http://swagger.io/terms/

	// @contact.name API Support
	// @contact.url http://www.swagger.io/support
	// @contact.email support@swagger.io

	// @license.name Apache 2.0
	// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

	// @host localhost:8000/
	// @BasePath awesome_project/api/v1
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
