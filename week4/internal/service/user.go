//service 层的主要主要就是 dto 数据和 do 数据的相互转换，它实现了v1包中的相关接口

package service

import (
	"context"
	"strconv"

	v1 "github.com/Liusiyuan-git/Go-learning/week4/apis/platform/v1"
	"github.com/Liusiyuan-git/Go-learning/week4/internal/domain"
	"github.com/pkg/errors"
	"google.golang.org/grpc/metadata"
)

var MetadataError = errors.New("meta data error")
var UserIdError = errors.New("user id error")

type UserService struct {
	v1.UserServerServer
	usecase domain.UserUsecase
}

func NewUserService(usecase domain.UserUsecase) *UserService {
	return &UserService{usecase: usecase}
}

// GetUserInfo 获取用户信息
func (u *UserService) GetUserInfo(ctx context.Context, req *v1.GetUserInfoRequest) (*v1.GetUserInfoResponse, error) {
	md, err := metadata.FromIncomingContext(ctx)
	if err != nil {
		return nil, errors.Wrap(MetadataError, "get metadata err")
	}

	id, err := strconv.Atoi(md.Get("id"))
	if err != nil {
		return nil, errors.Wrapf(UserIdError, "user id not a num, data: %v", md.Get("id"))
	}

	user, err := u.usecase.GetUserInfo(ctx, id)
	if err != nil {
		return nil, err
	}

	resp := &v1.GetUserInfoResponse{
		Username: user.Name,
		Id:       user.Id,
	}

	return resp, nil
}
