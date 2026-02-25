package users

type Repository interface {
	Create(user *NewUser) (string, error)
}
