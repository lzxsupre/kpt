package service

import (
	"context"
	"fmt"
	"math/rand"
	"strconv"
	"strings"

	"github.com/mivinci/abc/ecode"
	"github.com/mivinci/abc/log"
	"github.com/mivinci/abc/middlewares/auth"
	"github.com/mivinci/abc/time"
	"github.com/mivinci/kpt/model"
)

// Token 获取新 token
func (s *Service) Token(c context.Context, arg *model.ArgAuth) (string, error) {
	if !s.dao.CodeEqual(arg.Addr, arg.Code) {
		return "", ecode.CodeNotMatch
	}
	user, err := s.dao.QueryUser(c, &model.User{Email: arg.Addr})
	if err != nil {
		return "", ecode.UserNotFound
	}
	return auth.NewToken(s.c.Key.Secret, user.UID, auth.WithPerm(user.Status)).String()
}

// Code 发送邮件验证码
func (s *Service) Code(c context.Context, addr string) error {
	var nums [6]string
	rand.Seed(time.Now().Unix())
	perms := rand.Perm(10)
	for i := range nums {
		nums[i] = strconv.Itoa(perms[i])
	}
	code := strings.Join(nums[:], "")

	s.dao.CodeSet(addr, code)

	go func() {
		if err := s.mailer.Send(fmt.Sprintf("验证码: %s", code), fmt.Sprintf("您的验证码为 <strong>%s</strong>，10 分钟内有效，请勿转发。", code), []string{addr}); err != nil {
			log.Infof("email send eror(%v)", err)
		}
	}()
	return nil
}
