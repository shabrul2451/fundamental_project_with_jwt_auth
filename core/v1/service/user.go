package service

import v1 "awesomeProject/core/v1"

type UserService interface {
	Store(user v1.User) error
	Get() []v1.User
	GetByEmail(email string) v1.User
}
