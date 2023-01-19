package user

import "go.uber.org/zap"

type UserService interface {
	CreateUser(user *User) (*User, error)
}

type service struct {
	r      UserRepository
	logger *zap.SugaredLogger
}

// CreateUser implements UserService
func (s *service) CreateUser(incoming *User) (*User, error) {
	user, err := s.r.Create(incoming)
	if err != nil {
		s.logger.Errorw("error creating user", "error", err)
		return nil, err
	}
	return user, nil
}

func NewService(r UserRepository, logger *zap.SugaredLogger) UserService {
	return &service{r: r, logger: logger}
}
