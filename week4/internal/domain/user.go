//domain 这一层主要是包含 do 对象的定义，以及 usecase 和 repo 层的接口定义

package domain

import "context"

type User struct {
	Id   int
	Name string
}

type UserUsecase interface {
	GetUserInfo(ctx context.Context, id int) (*User, error)
}

type UserRepo interface {
	GetUserInfo(ctx context.Context, id int) (*User, error)
}
