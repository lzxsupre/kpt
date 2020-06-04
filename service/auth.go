package service

import (
	"context"

	"github.com/mivinci/abc/middlewares/auth"
	"github.com/mivinci/kpt/model"
)

// Token 获取新 token
func (s *Service) Token(c context.Context, arg *model.ArgAuth) (string, error) {

	return auth.NewToken(s.c.Key.Secret, arg.UID).String()
}

// AddUser 添加用户
func (s *Service) AddUser(c context.Context, user *model.User) error {
	user.Status = 1
	return s.dao.AddUser(c, user)
}

// QueryUsers 获取用户
func (s *Service) QueryUsers(c context.Context, user *model.User) (*model.UserResponse, error) {
	return s.dao.QueryUsers(c, user)
}

// UpdateUser 更新用户
func (s *Service) UpdateUser(c context.Context, user *model.User) error {
	return s.dao.UpdateUser(c, user)
}

// DeleteUser 删除用户(禁掉用户)
func (s *Service) DeleteUser(c context.Context, uid string) error {
	return s.dao.DeleteUser(c, uid)
}
