package db

import (
	"context"
	"fmt"

	pgx "github.com/jackc/pgx/v5"
	"github.com/rohitpandeydev/microservices/internal/types"
)

// this function is to get userdetials from databser using username
func (db *DB) GetUser(name string) (types.User, error) {
	db.logger.Debug("Fetching userdetails from database for user: %s", name)
	var user types.User

	err := db.conn.QueryRow(context.Background(),
		"SELECT id, username, email, dob, slots FROM library.user WHERE username = $1",
		name).Scan(&user.ID, &user.Name, &user.Email, &user.DOB, &user.Slots)

	if err != nil {
		if err == pgx.ErrNoRows {
			return types.User{}, fmt.Errorf("no user found with name %s", name)
		}
		return types.User{}, fmt.Errorf("database error: %w", err)
	}

	db.logger.Info("Successfully fetched user details for: %s", user.Name)
	return user, nil
}

func (db *DB) Login(name string) (string, error) {
	db.logger.Debug("Fetching login credentials for user: %s", name)
	var password string

	err := db.conn.QueryRow(context.Background(),
		"SELECT password FROM library.user WHERE username = $1",
		name).Scan(&password)

	if err != nil {
		if err == pgx.ErrNoRows {
			return "", fmt.Errorf("no user found with name %s", name)
		}
		return "", fmt.Errorf("database error: %w", err)
	}

	return password, nil
}
