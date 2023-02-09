package repository

import (
	"awesomeProject/core/v1"
)

type User interface {
	Store(user v1.User) error
	Get() []v1.User
	GetByEmail(email string) v1.User
}
