package domain

import "time"

type User struct {
	ID        uint      `json:"id"`
	FirstName string    `json:"first_name"`
	LastName  string    `json:"last_name"`
	Email     string    `json:"email"`
	Phone     string    `json:"phone"`
	Password  string    `json:"password"`
	Code      string    `json:"code"`
	Expire    time.Time `json:"expire"`
	Verify    bool      `json:"verify"`
	UserType  string    `json:"user_type"`
}
