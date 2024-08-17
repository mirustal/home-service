package models

import "time"

type Flat struct {
	ID        int       `json:"id" db:"id"`
	HouseID   int       `json:"house_id" db:"house_id"`
	Price     int       `json:"price" db:"price"`
	Rooms     int       `json:"rooms" db:"rooms"`
	Status    string    `json:"status" db:"status"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
	UpdatedAt time.Time `json:"updated_at" db:"updated_at"`
}