package database

import (
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
)

func encrypt(text string, key string) (string, error) {
	block, err := aes.NewCipher([]byte(key))
	if err != nil {
		return "", err
	}

	nonce := make([]byte, 16)
	ciphertext := make([]byte, len(text))

	stream := cipher.NewCFBEncrypter(block, nonce)
	stream.XORKeyStream(ciphertext, []byte(text))

	return base64.StdEncoding.EncodeToString(ciphertext), nil
}

func decrypt(encryptedText string, key string) (string, error) {
	block, err := aes.NewCipher([]byte(key))
	if err != nil {
		return "", err
	}

	nonce := make([]byte, 16)
	ciphertext, _ := base64.StdEncoding.DecodeString(encryptedText)

	stream := cipher.NewCFBDecrypter(block, nonce)
	stream.XORKeyStream(ciphertext, ciphertext)

	return string(ciphertext), nil
}