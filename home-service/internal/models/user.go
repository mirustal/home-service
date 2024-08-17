package models

import "time"

type User struct {
    ID        string `json:"id" db:"id"`
    Email     string    `json:"email" db:"email"`
    HashPass  []byte    `json:"password" db:"password"`
    Type      string    `json:"user_type" db:"user_type"`
    CreatedAt time.Time `json:"created_at" db:"created_at"`
    UpdatedAt time.Time `json:"updated_at" db:"updated_at"`
}
