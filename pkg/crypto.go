package pkg

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"io"
	"strings"
)

func EncryptString(value string, password string) (string, error) {
	key := hashTo32Bytes(password)
	plaintext := []byte(value)

	block, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}

	nonce := make([]byte, 12)
	if _, err = io.ReadFull(rand.Reader, nonce); err != nil {
		return "", err
	}

	aesgcm, err := cipher.NewGCM(block)
	if err != nil {
		return "", err
	}

	ciphertext := aesgcm.Seal(nil, nonce, plaintext, nil)
	result := base64.StdEncoding.EncodeToString([]byte(fmt.Sprintf("%x:%x", nonce, ciphertext)))
	return fmt.Sprintf("%s", result), nil
}

func DecryptString(value string, password string) (string, error) {
	key := hashTo32Bytes(password)

	data, err := base64.StdEncoding.DecodeString(value)
	if err != nil {
		return "", err
	}

	payload := strings.Split(string(data), ":")
	if len(payload) != 2 {
		return "", fmt.Errorf("Invalid encrypted value supplied")
	}

	ciphertext, err := hex.DecodeString(payload[1])
	if err != nil {
		return "", err
	}

	nonce, err := hex.DecodeString(payload[0])
	if err != nil {
		return "", err
	}

	block, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}

	aesgcm, err := cipher.NewGCM(block)
	if err != nil {
		return "", err
	}

	plaintext, err := aesgcm.Open(nil, nonce, ciphertext, nil)
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("%s", plaintext), nil
}

func DecryptValues(values map[string]interface{}, password string) (map[string]interface{}, error) {
	for k, v := range values {
		switch v.(type) {
		case string:
			if strings.HasPrefix(v.(string), "ENC|") {
				value := strings.TrimPrefix(v.(string), "ENC|")
				decryptedValue, err := DecryptString(value, password)
				if err != nil {
					return values, err
				}
				values[k] = decryptedValue
			}
		}
	}
	return values, nil
}

func hashTo32Bytes(input string) []byte {
	data := sha256.Sum256([]byte(input))
	return data[0:]
}
