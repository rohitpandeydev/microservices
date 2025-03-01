package types

import (
	"time"
)

type User struct {
	ID       int32  `json:"id"`
	Name     string `json:"name"`
	Password string `json:"password"`
}

type Books struct {
	ID        int32     `json:"id"`
	Name      string    `json:"name"`
	Genre     string    `json:"genre"`
	Available bool      `json:"available"`
	DueDate   time.Time `json:"date"`
	User      User      `json:"user"`
}
