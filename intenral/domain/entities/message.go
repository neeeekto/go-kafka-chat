package entities

import (
	"github.com/google/uuid"
	"time"
)

type MessageId string

func NewMessageId() MessageId {
	return MessageId(uuid.New().String())
}

type Message struct {
	Id      MessageId
	ChatID  ChatID    `json:"chat_id" bson:"chat_id"`
	From    UserID    `json:"from" bson:"from"`
	Content string    `json:"content"`
	Date    time.Time `json:"date"`
}
