package inbound

import (
	"hexarch/core/domain"
)

type UserService interface {
	GetUserByID(id string) (*domain.User, error)
	GetUserByUsername(u string) (*domain.User, error)
	CreateUser(input domain.CreateUserInput) (*domain.User, error)
}
