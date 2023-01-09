package auth

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/t0nyandre/go-boilerplate-oauth2/pkg/session"
)

func logout(w http.ResponseWriter, r *http.Request) {
	if err := session.ClearSession(w, r); err != nil {
		http.Error(w, fmt.Sprintf("Error clearing session: %v", err), http.StatusInternalServerError)
		return
	}

	fmt.Println("User logged out!")
	http.Redirect(w, r, "/", http.StatusFound)
}

func NewRoutes() chi.Router {
	r := chi.NewRouter()
	r.Get("/logout", logout)
	return r
}
