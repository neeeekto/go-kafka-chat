package entities

import "github.com/google/uuid"

type UserID string

func NewUserID() UserID {
	return UserID(uuid.New().String())
}

type User struct {
	Id   UserID `json:"id"`
	Name string `json:"name"`
}
