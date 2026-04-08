package auth

import (
	"fmt"

	"github.com/alexedwards/argon2id"
)

// hash a password using argon2id.CreateHash
func HashPassword(password string) (string, error) {
	hash, err := argon2id.CreateHash(password, argon2id.DefaultParams)
	if err != nil {
		return "", fmt.Errorf("error hashing password: %v", err)
	}
	return hash, nil
}

// compare an entered password with its hash in the db
func CheckPasswordHash(password, hash string) (bool, error) {
	match, err := argon2id.ComparePasswordAndHash(password, hash)
	if err != nil {
		return false, fmt.Errorf("error comparing hashed passwords: %v", err)
	}
	return match, nil
}
