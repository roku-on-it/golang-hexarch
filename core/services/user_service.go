package services

import (
	"hexarch/core/domain"
	"hexarch/core/ports/outbound"
)

type userService struct {
	userRepository outbound.UserRepository
}

func NewUserService(userRepository outbound.UserRepository) *userService {
	return &userService{userRepository: userRepository}
}

func (s *userService) GetUserByID(id string) (*domain.User, error) {
	return s.userRepository.FindUserByID(id)
}

func (s *userService) CreateUser(input domain.CreateUserInput) (*domain.User, error) {
	return s.userRepository.CreateUser(input)
}

func (s *userService) GetUserByUsername(u string) (*domain.User, error) {
	return s.userRepository.FindUserByUsername(u)
}
