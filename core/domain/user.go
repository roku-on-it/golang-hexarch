package domain

type User struct {
	ID       string
	Username string
	FullName string
	Age      int
}

type CreateUserInput struct {
	Username string
	FullName string
	Age      int
}
