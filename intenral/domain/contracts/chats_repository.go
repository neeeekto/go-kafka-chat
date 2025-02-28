package contracts

import (
	"context"

	"go-kafka-chat/intenral/domain/entities"
)

//go:generate mockery --name=ChatsRepository --inpackage-suffix --inpackage --case underscore
type ChatsRepository interface {
	Save(ctx context.Context, chat *entities.Chat) error
	AddMessage(ctx context.Context, message *entities.Message, chat *entities.Chat) error
}
