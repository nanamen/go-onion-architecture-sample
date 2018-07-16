package service

import (
	"context"

	"github.com/nanamen/go-echo-rest-sample/domain/repository"
)

// UserService ドメインサービスとして利用し,複数のエンティティやレポジトリを扱う処理をここで実装する.
// ※ ドメインサービスはアプリケーションサービスではないのでトランザクションの境界などは持たない.
// ※ なんでもドメインサービスで実装するとドメインモデル貧血症となるので気をつける(ドメインモデルで表現できないかよくよく検討すること).
type UserService interface {
	DoSomething(ctx context.Context, foo int) error
}

type userService struct {
	repository.UserRepository
}

// NewUserService UserServiceを取得します.
func NewUserService(r repository.UserRepository) UserService {
	return &userService{r}
}

func (u *userService) DoSomething(ctx context.Context, foo int) error {
	// some code
	return nil
}
