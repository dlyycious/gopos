package usermodule

import "time"

type UserDtos struct {
	Username  string    `json:"username"`
	Fullname  string    `json:"fullname"`
	Email     string    `json:"email"`
	Goid      uint      `json:"goid"`
	IsVerfied bool      `json:"is_verified"`
	CreatedAt time.Time `json:"created_at"`
	Role      string    `json:"role"`
	Token     string    `json:"token"`
}
