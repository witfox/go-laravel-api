package hash

import (
	"gohub/pkg/logger"

	"golang.org/x/crypto/bcrypt"
)

func BcryptHash(password string) string {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	logger.LogIf(err)

	return string(bytes)
}

func BcryptCheck(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))

	return err == nil
}

func BcryptIsHash(hash string) bool {
	return len(hash) == 60
}
