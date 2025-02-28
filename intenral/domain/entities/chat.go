package entities

import "github.com/google/uuid"

type ChatID string

func NewChatID() ChatID {
	return ChatID(uuid.New().String())
}

type Chat struct {
	Id    ChatID   `json:"id" bson:"id"`
	Name  string   `json:"name" bson:"name"`
	Users []UserID `json:"users" bson:"users"`
}
