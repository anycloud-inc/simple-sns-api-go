package auth

import "golang.org/x/crypto/bcrypt"

func EncryptPassword(s string) ([]byte, error) {
	return bcrypt.GenerateFromPassword([]byte(s), 12)
}

func ComparePassword(encrypted string, s string) error {
	return bcrypt.CompareHashAndPassword([]byte(encrypted), []byte(s))
}
