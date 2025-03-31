package auth

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/rohitpandeydev/microservices/pkg/logger"
)

var secretKey = []byte("secret-key")

// this function generates jwt toker after login authorisation and is valide for 24 hours
func CreateToken(username string, log *logger.Logger) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": username,
		"exp":      time.Now().Add(time.Hour * 24).Unix(),
	})

	tokenString, err := token.SignedString(secretKey)
	if err != nil {
		log.Error("Error when creating jwt toker for the username %s , err message: %v", username, err)
		return "", err
	}

	return tokenString, nil
}

func VerifyToken(tokenString string, log *logger.Logger) error {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (any, error) {
		return secretKey, nil
	})
	if err != nil {
		log.Error("Failure when verifying the token %v", err)
		return err
	}
	if !token.Valid {
		return fmt.Errorf("invalid token")
	}
	return nil
}

type Claims struct {
	UserID   int32  `json:"user_id"`
	Username string `json:"username"`
	jwt.RegisteredClaims
}

func generateToken(claims *Claims, log *logger.Logger) (string, error) {
	claims.RegisteredClaims = jwt.RegisteredClaims{
		ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
		IssuedAt:  jwt.NewNumericDate(time.Now()),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(secretKey)
	if err != nil {
		log.Error("Failed to generate token: %v", err)
		return "", fmt.Errorf("failed to generate token: %w", err)
	}

	return tokenString, nil
}
