package types

import (
	"time"
)

type UserCredentials struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type UserDetail struct {
	ID    int32     `json:"id"`
	Name  string    `json:"name"`
	Email string    `json:"email"`
	DOB   time.Time `json:"dob"`
	Slots int32     `json:"slots"`
}

type User struct {
	ID       int32     `json:"id"`
	Name     string    `json:"name"`
	Password string    `json:"password"`
	Email    string    `json:"email"`
	DOB      time.Time `json:"dob"`
	Slots    int32     `json:"slots"`
}

// LoginResult represents the possible outcomes of a login attempt
type LoginResult struct {
	Success bool
	Token   string
	User    *UserResponse
	Error   error
}

// type for books
type Books struct {
	ID        int32     `json:"id"`
	Name      string    `json:"name"`
	Genre     string    `json:"genre"`
	Available bool      `json:"available"`
	DueDate   time.Time `json:"date"`
	User      User      `json:"user"`
}

// type for categories
type Categories struct {
	ID            int32  `json:"id"`
	Name          string `json:"name"`
	NumberOfBooks int64  `json:"numberofbooks"`
}

// UserResponse represents the user data that's safe to send to clients
type UserResponse struct {
	ID    int32     `json:"id"`
	Name  string    `json:"name"`
	Email string    `json:"email"`
	DOB   time.Time `json:"dob"`
	Slots int32     `json:"slots"`
}

// ToResponse converts a User to UserResponse
func (u *User) ToResponse() *UserResponse {
	return &UserResponse{
		ID:    u.ID,
		Name:  u.Name,
		Email: u.Email,
		DOB:   u.DOB,
		Slots: u.Slots,
	}
}
