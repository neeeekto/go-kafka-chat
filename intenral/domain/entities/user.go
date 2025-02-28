package entities

type UserID string
type User struct {
	Id   UserID `json:"id"`
	Name string `json:"name"`
}
