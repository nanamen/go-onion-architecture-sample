package repository

import (
	"context"

	"github.com/nanamen/go-echo-rest-sample/domain/model"
)

// UserRepository interface
type UserRepository interface {
	Fetch(ctx context.Context) ([]*model.User, error)
	FetchByID(ctx context.Context, id int) (*model.User, error)
	Create(ctx context.Context, user *model.User) (*model.User, error)
	Update(ctx context.Context, user *model.User) (*model.User, error)
	Delete(ctx context.Context, id int) error
}
