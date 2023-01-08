package encryption

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/hmac"
	"crypto/rand"
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"fmt"
	"io"
	"os"
	"regexp"
)

func isHex(s []byte) bool {
	hexRegexp := regexp.MustCompile("^[0-9a-fA-F]+$")
	return hexRegexp.MatchString(string(s))
}

func getSessionSecret() ([]byte, error) {
	var err error
	secret := []byte(os.Getenv("SESSION_SECRET"))
	if secret == nil {
		panic("SESSION_SECRET is not set")
	}
	if isHex(secret) {
		secret, err = hex.DecodeString(string(secret))
		if err != nil {
			return nil, errors.New(fmt.Sprintf("Failed to decode hex: %s", err))
		}
		return secret, nil
	}
	return []byte(secret), nil
}

func Encrypt(incoming string) ([]byte, error) {
	incomingString := []byte(incoming)

	key, err := getSessionSecret()
	if err != nil {
		return nil, err
	}

	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("Failed to generate block: %s", err))
	}

	ciphertext := make([]byte, aes.BlockSize+len(incomingString))
	iv := ciphertext[:aes.BlockSize]
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		return nil, errors.New(fmt.Sprintf("Failed to generate nonce: %s", err))
	}

	stream := cipher.NewCFBEncrypter(block, iv)
	stream.XORKeyStream(ciphertext[aes.BlockSize:], incomingString)

	// Authenticate the message
	mac := hmac.New(sha256.New, key)
	mac.Write(ciphertext)
	expectedMAC := mac.Sum(nil)

	// Append the MAC to the cipherText
	ciphertext = append(ciphertext, expectedMAC...)

	return ciphertext, nil
}

func Decrypt(encryptedString []byte) (string, error) {
	key, err := getSessionSecret()
	if err != nil {
		return "", err
	}

	macSize := sha256.Size
	if len(encryptedString) < macSize {
		return "", errors.New("Encrypted string is too short")
	}

	expectedMAC := encryptedString[len(encryptedString)-macSize:]
	ciphertext := encryptedString[:len(encryptedString)-macSize]

	// Authenticate the message
	mac := hmac.New(sha256.New, key)
	mac.Write(ciphertext)
	if !hmac.Equal(expectedMAC, mac.Sum(nil)) {
		return "", errors.New("Invalid MAC")
	}

	block, err := aes.NewCipher(key)
	if err != nil {
		return "", errors.New(fmt.Sprintf("Failed to generate block: %s", err))
	}

	iv := ciphertext[:aes.BlockSize]
	ciphertext = ciphertext[aes.BlockSize:]

	stream := cipher.NewCFBDecrypter(block, iv)

	plaintext := make([]byte, len(ciphertext))
	stream.XORKeyStream(plaintext, ciphertext)

	return fmt.Sprintf("%s", plaintext), nil
}
