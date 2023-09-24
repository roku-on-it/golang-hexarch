package driven

import "hexarch/core/domain"

// Example implementation

type userPostgresAdapter struct{}

func NewUserPostgresAdapter(connAddr string, port int) *userPostgresAdapter {
	return &userPostgresAdapter{}
}

func (a *userPostgresAdapter) FindUserByID(id string) (*domain.User, error) {
	panic("implement me!")
}
func (a *userPostgresAdapter) FindUserByUsername(u string) (*domain.User, error) {
	panic("implement me!")
}
func (a *userPostgresAdapter) CreateUser(input domain.CreateUserInput) (*domain.User, error) {
	panic("implement me!")
}
