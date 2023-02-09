package v1

import (
	"awesomeProject/api/common"
	v1 "awesomeProject/core/v1"
	"awesomeProject/core/v1/api"
	"awesomeProject/core/v1/service"
	"github.com/labstack/echo/v4"
	"log"
)

type userApi struct {
	userService service.UserService
}

func (u userApi) Save(ctx echo.Context) error {
	formData := v1.User{}
	if err := ctx.Bind(&formData); err != nil {
		log.Println("Input Error:", err.Error())
		return common.GenerateErrorResponse(ctx, nil, "Failed to Bind Input!")
	}
	err := u.userService.Store(formData)
	if err != nil {
		return common.GenerateErrorResponse(ctx, err, "failed")
	}
	return common.GenerateSuccessResponse(ctx, "Stored successfully", nil, "success")
}

func (u userApi) Get(ctx echo.Context) error {
	return common.GenerateSuccessResponse(ctx, u.userService.Get(), nil, "successful")
}

func NewUserApi(userService service.UserService) api.UserApi {
	return &userApi{
		userService: userService,
	}
}
