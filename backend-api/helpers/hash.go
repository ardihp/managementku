package helpers

import "golang.org/x/crypto/bcrypt"

func HashPassword(password string) string {
	hashes, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(hashes)
}
