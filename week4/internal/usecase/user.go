//这一层主要是业务逻辑，业务逻辑相关代码都应该在这一层写，当然有时候我们的代码可能就只是保存一下数据没啥业务逻辑，可能是直接调用一下 repo 的方式

package usecase

import (
	"context"

	"github.com/Liusiyuan-git/Go-learning/week4/internal/domain"
)

type user struct {
	repo domain.UserRepo
}

func NewUserUsecase(repo domain.IUserRepo) domain.UserUsecase {
	return &user{repo: repo}
}

func (u *user) GetUserInfo(ctx context.Context, id int) (*domain.User, error) {
	return u.repo.GetUserInfo(ctx, id)
}
