package healthcheck

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/t0nyandre/go-rest-boilerplate/internal/config"
)

func RegisterHandlers(cfg *config.Config) chi.Router {
	res := resource{cfg}
	r := chi.NewRouter()
	r.Get("/", res.healthcheck)
	return r
}

type resource struct {
	cfg *config.Config
}

func (r resource) healthcheck(w http.ResponseWriter, req *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(fmt.Sprintf("Healthcheck:\n\t%s is OK ..", r.cfg.AppName)))
}
