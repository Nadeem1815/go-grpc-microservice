package utils

import "golang.org/x/crypto/bcrypt"

// we have two functions to encode a password with bcrypt, and then we have a validation function as well.
func HashPassword(password string) string {
	bytes, _ := bcrypt.GenerateFromPassword([]byte(password), 5)

	return string(bytes)
}

func CheckPasswordHash(password string, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))

	return err == nil
}
