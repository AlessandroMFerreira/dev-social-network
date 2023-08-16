package utils

import "golang.org/x/crypto/bcrypt"

func HashPassWord(password string) ([]byte, error) {
	return bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
}
