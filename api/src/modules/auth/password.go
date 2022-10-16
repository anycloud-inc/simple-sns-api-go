package auth

import "golang.org/x/crypto/bcrypt"

func EncryptPassword(s string) ([]byte, error) {
	return bcrypt.GenerateFromPassword([]byte(s), 12)
}
