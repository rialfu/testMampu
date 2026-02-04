package helpers

import (
	"errors"

	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(bytes), err
}

func CheckPassword(hashPassword string, plainPassword string) (bool, error) {
	hashPW := []byte(hashPassword)
	plain := []byte(plainPassword)
	if err := bcrypt.CompareHashAndPassword(hashPW, plain); err != nil {
		if errors.Is(err, bcrypt.ErrMismatchedHashAndPassword) {
			return false, nil
		}
		return false, err
	}
	return true, nil
}
