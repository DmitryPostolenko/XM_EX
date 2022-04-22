// Package hasher implements hashing string password and password vs hash comparison
package hasher

import (
	"golang.org/x/crypto/bcrypt"
)

// HashPassword hashes string password using bcrypt, returns hash, error
func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

// CheckPasswordHash compares password with its hash, returns bool
func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
