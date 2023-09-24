package outbound

import (
	"hexarch/core/domain"
)

type UserRepository interface {
	FindUserByID(id string) (*domain.User, error)
	FindUserByUsername(u string) (*domain.User, error)
	CreateUser(input domain.CreateUserInput) (*domain.User, error)
}
