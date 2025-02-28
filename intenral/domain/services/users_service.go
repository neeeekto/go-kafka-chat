package services

import (
	"context"
	"go-kafka-chat/intenral/domain/contracts"
	"go-kafka-chat/intenral/domain/entities"
)

type UsersService struct {
	repository contracts.UsersRepository
}

func NewUsersService(repository contracts.UsersRepository) *UsersService {
	return &UsersService{
		repository: repository,
	}
}

func (us *UsersService) Add(ctx context.Context, name string) (*entities.User, error) {
	user := &entities.User{
		Id:   entities.NewUserID(),
		Name: name,
	}
	err := us.repository.Save(ctx, user)
	if err != nil {
		return nil, err
	}
	return user, nil
}
