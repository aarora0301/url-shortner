package models

/**
Enumerates available keys
*/
type Key struct {
	Key string `json:"key" db:"id,pk"`
}

/**
Enumerates used keys
*/
type UsedKey struct {
	UsedKey string `json:"key" db:"id,pk"`
}
