package user

import (
	"github.com/gookit/validate"
	"github.com/t0nyandre/go-rest-template/internal/entity"
	"go.uber.org/zap"
)

type Service interface {
	Create(incoming *entity.User) (*entity.User, error)
}

type service struct {
	repo   Repository
	logger *zap.SugaredLogger
}

// Create implements Service
func (service) Create(incoming *entity.User) (*entity.User, error) {
	panic("unimplemented")
}

func NewService(repo Repository, logger *zap.SugaredLogger) Service {
	return service{repo: repo, logger: logger}
}

type CreateUserRequest struct {
	Email    string `json:"email" validate:"required|email"`
	Password string `json:"password" validate:"required|minLen:6"`
}

func (r *CreateUserRequest) Validate() error {
	v := validate.Struct(r)
	if v.Validate() {
		return nil
	}
	return v.Errors
}
