package state

import (
	"crypto/rand"
	"encoding/base64"
)

func generateRandomBytes(length int) ([]byte, error) {
	bytes := make([]byte, length)
	_, err := rand.Read(bytes)
	if err != nil {
		return nil, err
	}
	return bytes, nil
}

func generateRandomString(length int) (string, error) {
	bytes, err := generateRandomBytes(length)
	return base64.URLEncoding.EncodeToString(bytes), err
}

func GenerateRandomState() (string, error) {
	return generateRandomString(32)
}
