package v1

import (
	v1 "awesomeProject/core/v1"
	"awesomeProject/core/v1/repository"
	"awesomeProject/core/v1/service"
	"awesomeProject/enums"
	"errors"
	"github.com/google/uuid"
)

type user struct {
	userRepo repository.User
}

func (u user) GetByEmail(email string) v1.User {
	return u.userRepo.GetByEmail(email)
}

func (u user) Store(user v1.User) error {
	isUserExist := u.userRepo.GetByEmail(user.Email)
	if isUserExist.Email != "" && isUserExist.Status != enums.DELETED {
		return errors.New("email is already registered")
	}
	user.Id = uuid.NewString()
	err := u.userRepo.Store(user)
	if err != nil {
		return err
	}
	return nil
}

func (u user) Get() []v1.User {
	return u.userRepo.Get()
}

func NewUserService(userRepo repository.User) service.UserService {
	return &user{
		userRepo: userRepo,
	}
}
