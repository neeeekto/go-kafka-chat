package services

import (
	"context"
	"go-kafka-chat/intenral/domain"
	"go-kafka-chat/intenral/domain/contracts"
	"go-kafka-chat/intenral/domain/entities"
	"go-kafka-chat/intenral/domain/validators"
	"time"
)

type ChatsService struct {
	chatRepo contracts.ChatsRepository
	notifier contracts.MessageNotifier
	userRepo contracts.UsersRepository
}

func NewChatService(chatRepo contracts.ChatsRepository, notifier contracts.MessageNotifier, userRepo contracts.UsersRepository) *ChatsService {
	return &ChatsService{
		chatRepo: chatRepo,
		notifier: notifier,
		userRepo: userRepo,
	}
}

func (cs *ChatsService) Create(ctx context.Context, name string, users []entities.UserID) (*entities.Chat, error) {
	allUserExists, err := cs.userRepo.CheckExists(ctx, users...)
	if err != nil {
		return nil, err
	}

	if !allUserExists {
		return nil, domain.NewDomainValidationError("not all users exists")
	}

	chat := &entities.Chat{
		Id:    entities.NewChatID(),
		Users: users,
		Name:  name,
	}
	err = cs.chatRepo.Save(ctx, chat)
	if err != nil {
		return nil, err
	}

	return chat, nil
}

func (cs *ChatsService) Send(ctx context.Context, chat *entities.Chat, initiator entities.UserID,
	content string) (*entities.Message,
	error) {
	err := validators.ValidateMessageContent(content)
	if err != nil {
		return nil, err
	}

	msg := &entities.Message{
		Id:      entities.NewMessageId(),
		ChatID:  chat.Id,
		Date:    time.Now(),
		From:    initiator,
		Content: content,
	}

	err = cs.chatRepo.AddMessage(ctx, msg, chat)
	if err != nil {
		return nil, err
	}

	err = cs.notifier.Notify(ctx, chat, msg)
	if err != nil {
		return nil, err
	}
	return msg, nil
}
