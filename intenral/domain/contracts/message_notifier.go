package contracts

import (
	"context"
	"go-kafka-chat/intenral/domain/entities"
)

//go:generate mockery --name=MessageNotifier --inpackage-suffix --inpackage --case underscore
type MessageNotifier interface {
	Notify(ctx context.Context, chat *entities.Chat, msg *entities.Message) error
}
