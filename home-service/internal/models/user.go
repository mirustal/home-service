package models

type User struct {
	ID    string `json:"id" db:"id"`
	Token string `json:"token"`
}
