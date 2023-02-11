package rpc

import (
	"context"

	"github.com/ozline/tiktok/api-gateway/biz/model/model"
	"github.com/ozline/tiktok/pkg/errno"
	"github.com/ozline/tiktok/services/user/kitex_gen/tiktok/user"
)

// Login returns token
func UserLogin(ctx context.Context, req *user.DouyinUserLoginRequest) (string, error) {
	resp, err := userClient.Login(ctx, req)

	if err != nil {
		return "", err
	}

	if resp.StatusCode != errno.SuccessCode {
		return "", errno.NewErrNo(int64(resp.StatusCode), resp.StatusMsg)
	}

	return resp.Token, nil
}

func UserRegister(ctx context.Context, req *user.DouyinUserRegisterRequest) (bool, error) {
	resp, err := userClient.Register(ctx, req)

	if err != nil {
		return false, err
	}

	if resp.StatusCode != errno.SuccessCode {
		return false, errno.NewErrNo(int64(resp.StatusCode), resp.StatusMsg)
	}

	return true, nil
}

func UserGetInfo(ctx context.Context, req *user.DouyinUserRequest) (*model.User, error) {
	resp, err := userClient.Info(ctx, req)

	if err != nil {
		return nil, err
	}

	if resp.StatusCode != errno.SuccessCode {
		return nil, errno.NewErrNo(int64(resp.StatusCode), resp.StatusMsg)
	}

	return &model.User{
		Id:            resp.User.Id,
		Name:          resp.User.Name,
		FollowCount:   resp.User.FollowCount,
		FollowerCount: resp.User.FollowerCount,
	}, nil
}
