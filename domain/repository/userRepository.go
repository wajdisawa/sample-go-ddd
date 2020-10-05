package repository

import (
	"context"
	"sample-go-ddd/domain/entity"
)

type UserRepository interface {
	Get(ctx context.Context, id int) (*entity.User, error)
	GetAll(ctx context.Context) ([]*entity.User, error)
	Save(ctx context.Context, user *entity.User) error
}
