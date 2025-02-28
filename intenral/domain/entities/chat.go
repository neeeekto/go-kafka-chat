package entities

type ChatID string
type Chat struct {
	Id   ChatID `json:"id"`
	Name string `json:"name"`
}
