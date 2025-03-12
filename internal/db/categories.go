package db

import (
	"context"

	"github.com/rohitpandeydev/microservices/internal/types"
)

func (db *DB) GetCategories() ([]types.Categories, error) {
	db.logger.Debug("Fetching categories from database")
	categories := make([]types.Categories, 0, 10)
	rows, err := db.conn.Query(context.Background(), "SELECT id,name,numberofbook FROM library.categories")
	if err != nil {
		db.logger.Error("Failed to execute category query %v", err)
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var cat types.Categories
		err := rows.Scan(&cat.ID, &cat.Name, &cat.NumberOfBooks)
		if err != nil {
			db.logger.Error("failed to scan category : %v", err)
			return nil, err
		}
		categories = append(categories, cat)
	}
	db.logger.Info("Successfully fetched %d categories", len(categories))
	return categories, nil
}
