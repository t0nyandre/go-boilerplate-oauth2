package user

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"go.uber.org/zap"
)

func NewRoutes(service UserService, logger *zap.SugaredLogger) chi.Router {
	res := resource{service: service, logger: logger}
	r := chi.NewRouter()
	r.Get("/me", res.me)
	return r
}

type resource struct {
	service UserService
	logger  *zap.SugaredLogger
}

func (r *resource) me(w http.ResponseWriter, req *http.Request) {
	panic("not implemented")
}
