package utils

import (
	"errors"

	"golang.org/x/crypto/bcrypt"
)

// PasswordMatch is a method that takes a plain text password and matches
// it with hash password and returns a boolean and an error.
func PasswordMatch(plainTextPassword, userPassword string) (bool, error) {
	err := bcrypt.CompareHashAndPassword([]byte(userPassword), []byte(plainTextPassword))

	if err != nil {
		switch {
		case errors.Is(err, bcrypt.ErrMismatchedHashAndPassword):
			// Invalid Password
			return false, nil
		default:
			return false, err
		}
	}

	return true, nil
}
