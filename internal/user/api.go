package user

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"go.uber.org/zap"
)

func RegisterHandlers(service Service, logger *zap.SugaredLogger) chi.Router {
	res := resource{service: service, logger: logger}
	r := chi.NewRouter()
	r.Post("/", res.create)

	return r
}

type resource struct {
	service Service
	logger  *zap.SugaredLogger
}

func (r resource) create(w http.ResponseWriter, req *http.Request) {
	panic("not implemented")
}
