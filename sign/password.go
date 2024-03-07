package sign

import (
	"github.com/zeromicro/go-zero/core/logx"
	"golang.org/x/crypto/bcrypt"
)

func GenerateHashFromPassword(password string) (hash string, err error) {
	// Hashing the password with the default cost of 10
	bcryptHash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(bcryptHash), nil
}

func VerifyPassword(password, bcryptHash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(bcryptHash), []byte(password))
	if err != nil {
		logx.Error(err)
		return false
	} else {
		return true
	}
}
