package usecase

import (
	"context"

	"github.com/nanamen/go-echo-rest-sample/domain/model"
	"github.com/nanamen/go-echo-rest-sample/domain/repository"
)

// UserUseCase interfase
type UserUseCase interface {
	GetUsers(ctx context.Context) ([]*model.User, error)
	GetUser(ctx context.Context, id int) (*model.User, error)
	CreateUser(ctx context.Context, user *model.User) (*model.User, error)
	UpdateUser(ctx context.Context, id int) (*model.User, error)
	DeleteUser(ctx context.Context, id int) error
}

type userUseCase struct {
	repository.UserRepository
}

// NewUserUseCase UserUseCaseを取得します.
func NewUserUseCase(r repository.UserRepository) UserUseCase {
	return &userUseCase{r}
}

func (u *userUseCase) GetUsers(ctx context.Context) ([]*model.User, error) {
	return u.UserRepository.Fetch(ctx)
}

func (u *userUseCase) GetUser(ctx context.Context, id int) (*model.User, error) {
	return u.UserRepository.FetchByID(ctx, id)
}

func (u *userUseCase) CreateUser(ctx context.Context, user *model.User) (*model.User, error) {
	return u.UserRepository.Create(ctx, user)
}

func (u *userUseCase) UpdateUser(ctx context.Context, id int) (*model.User, error) {
	user, err := u.UserRepository.FetchByID(ctx, id)
	if err != nil {
		return nil, err
	}
	return u.UserRepository.Update(ctx, user)
}

func (u *userUseCase) DeleteUser(ctx context.Context, id int) error {
	return u.UserRepository.Delete(ctx, id)
}
