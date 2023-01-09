package github

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/t0nyandre/go-boilerplate-oauth2/pkg/session"
	"golang.org/x/oauth2"
)

func (p *Provider) Login(w http.ResponseWriter, r *http.Request) {
	_, err := session.GetSession(r)
	if err == nil {
		http.Redirect(w, r, "/", http.StatusFound)
		return
	}
	url := p.Config.AuthCodeURL(p.State)
	http.Redirect(w, r, url, http.StatusFound)
}

func (p *Provider) Callback(w http.ResponseWriter, r *http.Request) {
	tokenData, err := session.GetSession(r)
	if err == nil {
		http.Redirect(w, r, "/", http.StatusFound)
		return
	}

	code := r.URL.Query().Get("code")

	tokenData, err = p.getAccessToken(code)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error exchanging token: %v", err), http.StatusInternalServerError)
		return
	}

	client := p.Config.Client(context.Background(), &oauth2.Token{AccessToken: tokenData.Token})
	userInfo, err := client.Get("https://api.github.com/user")
	if err != nil {
		http.Error(w, fmt.Sprintf("Error getting userinfo: %v", err), http.StatusInternalServerError)
		return
	}
	defer userInfo.Body.Close()

	var user struct {
		Login string `json:"login"`
	}

	fmt.Println(userInfo)

	if err := json.NewDecoder(userInfo.Body).Decode(&user); err != nil {
		http.Error(w, fmt.Sprintf("Error decoding userinfo: %v", err), http.StatusInternalServerError)
		return
	}
	tokenData.User = user.Login
	// TODO: Save userinfo to database

	if err := session.SetSession(w, *tokenData); err != nil {
		http.Error(w, fmt.Sprintf("Error setting session: %v", err), http.StatusInternalServerError)
		return
	}

	fmt.Printf("User %s is logged in!\n", tokenData.User)
	http.Redirect(w, r, "/", http.StatusFound)
}

func (p *Provider) NewRoutes() chi.Router {
	r := chi.NewRouter()
	r.Get("/login", p.Login)
	r.Get("/callback", p.Callback)
	return r
}
