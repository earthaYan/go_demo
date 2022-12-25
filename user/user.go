package user

import (
	"context"

	"github.com/marmotedu/iam/internal/apiserver/store"
)

type User struct {
}

func (c *User) GetUserInfo(ctx context.Context, r *GetUserInfoRequest) (*GetUserInfoResponse, error) {
	// 通过判断 r.Username 是否为 nil，来判断客户端是否传入了 Username 参数。
	if r.Username != nil {
		return store.Client().Users().GetUserByName(r.Class, r.Username)

	}
	return store.Client().Users().GetUserByID(r.Class, r.UserId)
}
