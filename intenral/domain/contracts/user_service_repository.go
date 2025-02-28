package contracts

import (
	"go-kafka-chat"
)

type UserServiceRepository interface {
	Save(user entities.User) error
}
