// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.15.0

package sqlc

import (
	"time"
)

type Account struct {
	ID           int64     `json:"id"`
	Firstname    string    `json:"firstname"`
	Lastname     string    `json:"lastname"`
	Emailaddress string    `json:"emailaddress"`
	Username     string    `json:"username"`
	Password     int64     `json:"password"`
	CreatedAt    time.Time `json:"created_at"`
}