package utils

import (
	"math/rand"
	"time"
)

func randPassword(length int) (res string) {
	const charset = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	password := make([]byte, length)
	seededRand := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := range password {
		password[i] = charset[seededRand.Intn(len(charset))]
	}
	return string(password)
}
