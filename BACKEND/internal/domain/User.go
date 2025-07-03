package domain

import "time"

const (
	SELLER = "seller"
	BUYER  = "buyer"
)

type User struct {
	ID        uint      `json:"id" gorm:"primaryKey"`
	FirstName string    `json:"first_name"`
	LastName  string    `json:"last_name"`
	Email     string    `json:"email" gorm:"index;unique;not null"`
	Phone     string    `json:"phone" gorm:"index;unique;not null"`
	Password  string    `json:"password"`
	Code      string    `json:"code"`
	Expire    time.Time `json:"expire"`
	Verified  bool      `json:"verify" gorm:"default:false"`
	UserType  string    `json:"user_type" gorm:"default:'buyer';not null"`
	CreatedAt time.Time `json:"created_at" gorm:"currentTime"`
	UpdatedAt time.Time `json:"updated_at" gorm:"currentTime"`
}
