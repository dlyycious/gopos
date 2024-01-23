package authmodule

import (
	"time"
)

// UserDtos represents the data transfer object for user information.
type UserDtos struct {
	Username  string    `json:"username"`
	Fullname  string    `json:"fullname"`
	Email     string    `json:"email"`
	Goid      uint      `json:"goid"`
	IsVerfied bool      `json:"is_verified"`
	CreatedAt time.Time `json:"created_at"`
	Role      string    `json:"role"`
}

// LoginDtos represents the data transfer object for login credentials.
type LoginDtos struct {
	Username  string    `json:"username"`
	Password  string    `json:"password"`
	Fullname  string    `json:"fullname"`
	Email     string    `json:"email"`
	Goid      uint      `json:"goid"`
	IsVerfied bool      `json:"is_verified"`
	CreatedAt time.Time `json:"created_at"`
	Role      string    `json:"role"`
	Token     string    `json:"token"`
}
