package models

import "time"

// House модель для таблицы houses
type House struct {
	ID        int       `json:"id" db:"id"`
	Address   string    `json:"address" db:"address"`
	Year      int       `json:"year" db:"year"`
	Developer string    `json:"developer,omitempty" db:"developer"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
	UpdatedAt time.Time `json:"updated_at" db:"updated_at"`
}

// Flat модель для таблицы flats
type Flat struct {
	ID        int       `json:"id" db:"id"`
	HouseID   int       `json:"house_id" db:"house_id"`
	Price     int       `json:"price" db:"price"`
	Rooms     int       `json:"rooms" db:"rooms"`
	Status    string    `json:"status" db:"status"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
	UpdatedAt time.Time `json:"updated_at" db:"updated_at"`
}

// Subscription модель для таблицы subscriptions
type Subscription struct {
	ID        int       `json:"id" db:"id"`
	UserID    int       `json:"user_id" db:"user_id"`
	HouseID   int       `json:"house_id" db:"house_id"`
	Email     string    `json:"email" db:"email"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
}
