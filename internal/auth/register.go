package auth

import (
	"fmt"

	"github.com/rohitpandeydev/microservices/pkg/logger"
	"golang.org/x/crypto/bcrypt"
)

// need a handler for registering the user so like username,password,email,dob//
// also a db method to write the user, and username should be unique else give error that username already exist this is done
// so have to call this method hashPassword to store the password
func hashPassword(password string, log *logger.Logger) (string, error) {
	if password == "" {
		return "", ErrEmptyCredentials
	}

	log.Debug("Generating password hash")
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		log.Error("Password hashing failed: %v", err)
		return "", fmt.Errorf("password hashing failed: %w", err)
	}
	return string(bytes), nil
}
