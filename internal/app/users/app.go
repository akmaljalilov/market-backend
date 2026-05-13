package users

import (
	"market/internal/utils"
	"market/internal/utils/security"
)

type App struct {
	repo          Repository
	notifyService NotifyService
}

// New constructor
func New(repo Repository, notifyService NotifyService) *App {
	return &App{repo, notifyService}
}

// Register a new user
func (a *App) Register(name, email, password string, phoneNumbers []string, role Role, data string) (*User, error) {
	username := utils.NormalizeUsername(name)
	user, err := a.repo.GetByUsername(username)
	if err != nil {
		resp, err := a.repo.Create(&NewUser{
			Username:     username,
			Name:         name,
			Email:        email,
			Password:     password,
			Data:         data,
			PhoneNumbers: phoneNumbers,
			Role:         role,
		})
		if err != nil {
			return nil, err
		}
		return resp, nil
	}
	return user, nil

}

// SignIn
func (a *App) SignIn(name, password string) (*SignInResponse, error) {
	user, err := a.repo.SignIn(name, password)
	if err != nil {
		return nil, err
	}
	resp := &SignInResponse{User: *user}
	token, err := security.GetToken(user.ID)
	if err != nil {
		return nil, err
	}
	resp.Token = token
	return resp, nil
}

// GetUserById
func (a *App) GetUserById(id string) (*User, error) {
	return a.repo.GetById(id)

}
func (a *App) RegisterDealer(name string, phoneNumbers []string, data string) (string, error) {
	username := utils.NormalizeUsername(name)
	user, err := a.repo.GetByUsername(username)
	if err != nil {
		password, _ := security.GeneratePassword(5)
		resp, err := a.repo.Create(&NewUser{
			Username:     username,
			Name:         name,
			Password:     password,
			Data:         data,
			PhoneNumbers: phoneNumbers,
			Role:         RoleDealer,
		})
		if err != nil {
			return "", err
		}
		return resp.ID, nil
	}
	return user.ID, nil
}
