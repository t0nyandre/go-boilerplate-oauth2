package session

import (
	"encoding/hex"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/t0nyandre/go-boilerplate-oauth2/pkg/utils/encryption"
)

type TokenData struct {
	User    string
	Token   string
	Expires time.Time
}

// Encrypt access token and store in cookie with expires time
func SetSession(w http.ResponseWriter, tokenData TokenData) error {
	encryptedAccessToken, err := encryption.Encrypt(tokenData.Token)
	if err != nil {
		return err
	}

	http.SetCookie(w, &http.Cookie{
		Name:     os.Getenv("SESSION_NAME"),
		Value:    fmt.Sprintf("%x", encryptedAccessToken),
		Path:     "/",
		Expires:  tokenData.Expires,
		HttpOnly: true,
	})

	http.SetCookie(w, &http.Cookie{
		Name:     fmt.Sprintf("%s_user", os.Getenv("SESSION_NAME")),
		Value:    tokenData.User,
		Path:     "/",
		Expires:  tokenData.Expires,
		HttpOnly: true,
	})

	return nil
}

// Decrypt access token and return token data
func GetSession(r *http.Request) (*TokenData, error) {
	tokenData := &TokenData{}
	cookie, err := r.Cookie(os.Getenv("SESSION_NAME"))
	if err != nil {
		return nil, err
	}

	user, err := r.Cookie(fmt.Sprintf("%s_user", os.Getenv("SESSION_NAME")))
	if err != nil {
		return nil, err
	}

	encryptedAccessToken, err := hex.DecodeString(cookie.Value)
	if err != nil {
		return nil, err
	}

	accessToken, err := encryption.Decrypt(encryptedAccessToken)
	if err != nil {
		return nil, err
	}

	tokenData.User = user.Value
	tokenData.Token = accessToken
	tokenData.Expires = cookie.Expires

	return tokenData, nil
}

// Delete session Cookie
func ClearSession(w http.ResponseWriter, r *http.Request) error {
	cookie, err := r.Cookie(os.Getenv("SESSION_NAME"))
	if err != nil {
		return err
	}

	user, err := r.Cookie(fmt.Sprintf("%s_user", os.Getenv("SESSION_NAME")))
	if err != nil {
		return err
	}

	http.SetCookie(w, &http.Cookie{
		Name:     cookie.Name,
		Value:    "",
		Path:     "/",
		MaxAge:   -1,
		HttpOnly: true,
	})

	http.SetCookie(w, &http.Cookie{
		Name:     user.Name,
		Value:    "",
		Path:     "/",
		MaxAge:   -1,
		HttpOnly: true,
	})

	return nil
}
