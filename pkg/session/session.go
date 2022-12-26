package session

import (
	"context"
	"crypto/rand"
	"encoding/hex"
	"errors"
	"fmt"
	"net/http"
	"os"
	"time"
	"vendor/golang.org/x/crypto/chacha20poly1305"
)

type TokenData struct {
	User    string
	Token   string
	Expires time.Time
}

// TODO: Create a function to delete session

// Encrypt access token and store in cookie with expires time
// TODO: Use ctx to get datastore to store session information
func SetSession(ctx context.Context, w http.ResponseWriter, tokenData TokenData) error {
	nonce := make([]byte, 12)
	if _, err := rand.Read(nonce); err != nil {
		return errors.New("Failed to generate nonce")
	}

	block, err := chacha20poly1305.New([]byte(os.Getenv("SESSION_SECRET")))
	if err != nil {
		return err
	}

	accessToken := []byte(tokenData.Token)
	encryptedAccessToken := block.Seal(nil, nonce, accessToken, nil)

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
func GetSession(ctx context.Context, r *http.Request) (*TokenData, error) {
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

	block, err := chacha20poly1305.New([]byte(os.Getenv("SESSION_SECRET")))
	if err != nil {
		return nil, err
	}

	nonceSize := block.NonceSize()
	nonce, cipherText := encryptedAccessToken[:nonceSize], encryptedAccessToken[nonceSize:]
	decryptedAccessToken, err := block.Open(nil, nonce, cipherText, nil)
	if err != nil {
		return nil, err
	}

	tokenData.User = user.Value
	tokenData.Token = string(decryptedAccessToken)
	tokenData.Expires = cookie.Expires

	return tokenData, nil
}
