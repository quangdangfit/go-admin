package utils

import (
	"github.com/quangdangfit/gosdk/utils/logger"
	"golang.org/x/crypto/bcrypt"

	"github.com/quangdangfit/go-admin/pkg/errors"
)

// HashPassword
func HashPassword(pass []byte) (string, error) {
	hashed, err := bcrypt.GenerateFromPassword(pass, bcrypt.MinCost)
	if err != nil {
		logger.Error("Failed to generate password: ", err)
		return "", errors.Wrap(err, "utils.HashPassword")
	}

	return string(hashed), nil
}
