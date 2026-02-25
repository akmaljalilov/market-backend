package users

type Service struct {
	repo Repository
}

// NewService constructor
func NewService(repo Repository) *Service {
	return &Service{repo: repo}
}

// Register a new user
func (s *Service) Register(name, email, password string, phoneNumbers []string, role Role, data string) (string, error) {
	id, err := s.repo.Create(&NewUser{
		Name:         name,
		Email:        email,
		Password:     password,
		Data:         data,
		PhoneNumbers: phoneNumbers,
		Role:         role,
	})
	if err != nil {
		return "", err
	}
	return id, nil
}
