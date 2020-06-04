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
