package contracts

import (
	"context"
	"go-kafka-chat/intenral/domain/entities"
)

//go:generate mockery --name=UsersRepository --inpackage-suffix --inpackage --case underscore
type UsersRepository interface {
	Save(ctx context.Context, user *entities.User) error
	CheckExists(ctx context.Context, userID ...entities.UserID) (bool, error)
}
