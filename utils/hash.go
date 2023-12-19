package utils

import "golang.org/x/crypto/bcrypt"

func HasPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}