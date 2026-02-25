package users

type Repository interface {
	Create(user *NewUser) (*User, error)
	GetByUsername(username string) (*User, error)
}
