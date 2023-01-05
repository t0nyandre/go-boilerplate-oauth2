package utils

import (
	"crypto/rand"
	"errors"
	"fmt"
	"os"
	"vendor/golang.org/x/crypto/chacha20poly1305"
)

func NewEncryption(incoming string) ([]byte, error) {
	nonce := make([]byte, 12)
	if _, err := rand.Read(nonce); err != nil {
		return nil, errors.New(fmt.Sprintf("Failed to generate nonce: %s", err))
	}

	block, err := chacha20poly1305.New([]byte(os.Getenv("SESSION_SECRET")))
	if err != nil {
		return nil, errors.New(fmt.Sprintf("Failed to generate block: %s", err))
	}

	incomingString := []byte(incoming)
	encryptedString := block.Seal(nil, nonce, incomingString, nil)

	return encryptedString, nil
}

func Decrypt(encryptedString []byte) (string, error) {
	block, err := chacha20poly1305.New([]byte(os.Getenv("SESSION_SECRET")))
	if err != nil {
		return "", errors.New(fmt.Sprintf("Failed to generate block: %s", err))
	}

	nonceSize := block.NonceSize()
	nonce, cipherText := encryptedString[:nonceSize], encryptedString[nonceSize:]
	decryptedString, err := block.Open(nil, nonce, cipherText, nil)
	if err != nil {
		return "", errors.New(fmt.Sprintf("Failed to decrypt string: %s", err))
	}

	return string(decryptedString), nil
}
