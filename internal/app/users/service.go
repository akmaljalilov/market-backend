package users

type Service struct {
	repo Repository
}

// NewService constructor
func NewService(repo Repository) *Service {
	return &Service{repo: repo}
}

// Register a new user
func (s *Service) Register(name, email, password string, phoneNumbers []string, role Role, data string) (*User, error) {
	user, err := s.repo.GetByUsername(name)
	if err != nil {
		resp, err := s.repo.Create(&NewUser{
			Username:     name,
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
