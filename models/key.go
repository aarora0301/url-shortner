package models


type Key struct {
	Key string `json:"key" db:"id,pk"`
}

type UsedKey struct{
	UsedKey string `json:"key" db:"id,pk"`
}