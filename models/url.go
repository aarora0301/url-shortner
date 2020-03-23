package models

import "time"

type Url struct {
	Hash           string    `json:"hash" db:"hash,pk"`
	OriginalUrl    string    `json:"original_url" db:"original_url"`
	CreationDate   time.Time `json:"creation_date" db:"creation_date"`
	ExpirationDate time.Time `json:"expiration_date" db:"expiration_date"`
	UserID         int       `json:"user_id" db:"user_id"`
}
