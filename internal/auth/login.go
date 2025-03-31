package auth

import (
	"errors"
	"fmt"

	"github.com/rohitpandeydev/microservices/internal/db"
	"github.com/rohitpandeydev/microservices/internal/types"
	"github.com/rohitpandeydev/microservices/pkg/logger"
	"golang.org/x/crypto/bcrypt"
)

var (
	ErrInvalidCredentials = errors.New("invalid credentials")
	ErrEmptyCredentials   = errors.New("username or password cannot be empty")
)

// Login authenticates a user and returns a JWT token if successful
func Login(creds types.UserCredentials, dbClient *db.DB, log *logger.Logger) types.LoginResult {
	// Validate input
	if creds.Username == "" || creds.Password == "" {
		return types.LoginResult{
			Success: false,
			Error:   ErrEmptyCredentials,
		}
	}

	log.Debug("Attempting login for user: %s", creds.Username)

	hashedPassword, err := dbClient.Login(creds.Username)
	if err != nil {
		log.Error("Database login error: %v", err)
		return types.LoginResult{
			Success: false,
			Error:   fmt.Errorf("authentication failed: %w", err),
		}
	}

	// Verify password
	valid, err := checkPasswordHash(creds.Password, hashedPassword, log)
	if err != nil {
		log.Error("Password verification error: %v", err)
		return types.LoginResult{
			Success: false,
			Error:   fmt.Errorf("password verification failed: %w", err),
		}
	}

	if !valid {
		log.Warn("Invalid password attempt for user: %s", creds.Username)
		return types.LoginResult{
			Success: false,
			Error:   ErrInvalidCredentials,
		}
	}

	// Get user details
	user, err := dbClient.GetUser(creds.Username)
	if err != nil {
		log.Error("Failed to get user details: %v", err)
		return types.LoginResult{
			Success: false,
			Error:   fmt.Errorf("failed to get user details: %w", err),
		}
	}

	// Generate JWT token
	claims := &Claims{
		UserID:   user.ID,
		Username: user.Name,
	}

	token, err := generateToken(claims, log)
	if err != nil {
		log.Error("Token generation error: %v", err)
		return types.LoginResult{
			Success: false,
			Error:   fmt.Errorf("token generation failed: %w", err),
		}
	}

	log.Info("Successful login for user: %s", creds.Username)
	return types.LoginResult{
		Success: true,
		Token:   token,
		User:    user.ToResponse(),
		Error:   nil,
	}
}

func checkPasswordHash(password, hash string, log *logger.Logger) (bool, error) {
	if password == "" || hash == "" {
		return false, ErrEmptyCredentials
	}

	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	if err != nil {
		if errors.Is(err, bcrypt.ErrMismatchedHashAndPassword) {
			return false, nil // Password doesn't match but not an error
		}
		log.Error("Password comparison failed: %v", err)
		return false, fmt.Errorf("password comparison failed: %w", err)
	}

	return true, nil
}
