package users

type Repository interface {
	Create(user *NewUser) (*User, error)
	GetByUsername(username string) (*User, error)
	GetById(id string) (*User, error)
	SignIn(username, password string) (*User, error)
}
