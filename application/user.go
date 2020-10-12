package application

import (
	"context"
	"sample-go-ddd/domain/entity"
	"sample-go-ddd/domain/repository"
)

type UserRepo struct {
	Repo repository.UserRepository
}

func (ur UserRepo) GetUser(ctx context.Context, id int) (*entity.User, error) {
		return ur.Repo.Get(ctx, id)
}
func (ur UserRepo) Users(ctx context.Context) ([]*entity.User, error) {
	return ur.Repo.GetAll(ctx)
}
func (ur UserRepo) Save(ctx context.Context, name string, email string) error {
	usr, err := entity.NewUser(name, email)
	if err != nil {
		return err
	}
	return ur.Repo.Save(ctx, usr)
}
