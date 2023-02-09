package api

import "github.com/labstack/echo/v4"

type UserApi interface {
	Save(ctx echo.Context) error
	Get(ctx echo.Context) error
}
