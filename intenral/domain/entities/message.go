package entities

import "time"

type Message struct {
	Id      string
	ChatID  ChatID    `json:"chat_id"`
	From    UserID    `json:"from"`
	To      UserID    `json:"to"`
	Content string    `json:"content"`
	Date    time.Time `json:"date"`
}
