package db

import (
	"context"
	"fmt"

	pgx "github.com/jackc/pgx/v5"
	"github.com/rohitpandeydev/microservices/internal/config"
	"github.com/rohitpandeydev/microservices/pkg/logger"
)

// for connection with database
type DB struct {
	conn   *pgx.Conn
	logger *logger.Logger
}

// DBConfig is getting imported from internal/config/config.go
// this function return database connection it take config input and logger as input for logging
func NewDB(cfg *config.DBConfig, log *logger.Logger) (*DB, error) {
	connStr := fmt.Sprintf("postgres://%s:%s@%s:%s/%s", cfg.User, cfg.Password, cfg.Host, cfg.Port, cfg.DBName)
	log.Info("Connecting to database at %s:%s", cfg.Host, cfg.Port)
	conn, err := pgx.Connect(context.Background(), connStr)
	if err != nil {
		return nil, err
	}

	return &DB{conn: conn, logger: log}, nil
}
