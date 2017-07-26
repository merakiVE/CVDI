package auth

import (
	"golang.org/x/crypto/bcrypt"
)

func VerifyPassword(_hashedPassword, _password []byte) (bool) {
	err := bcrypt.CompareHashAndPassword(_hashedPassword, _password)

	if err != nil {
		return false
	}
	return true
}
