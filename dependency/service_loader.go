package dependency

import (
	"awesomeProject/core/v1/service"
	v1 "awesomeProject/logic/v1"
	"awesomeProject/repository/v1/mongo"
)

func GetV1UserService() service.UserService {
	return v1.NewUserService(mongo.NewUserRepository(3000))
}

// GetV1JwtService returns service.Jwt
func GetV1JwtService() service.Jwt {
	return v1.NewJwtService()
}
