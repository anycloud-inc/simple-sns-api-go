package auth

import "golang.org/x/crypto/bcrypt"

func EncryptPassword(s string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(s), 12)
	if err != nil {
		return "", err
	}
	return string(hash), nil
}
