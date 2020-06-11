package service

import (
	"bytes"
	"context"
	"fmt"

	"github.com/mivinci/abc/log"
	"github.com/mivinci/kpt/model"
)

// PunchRec 获取打卡记录
func (s *Service) PunchRec(c context.Context, record *model.PunchRec) ([]*model.PunchRec, error) {
	return s.dao.QueryPunchRec(c, record)
}

// PunchRecBetween 按时间段获取打卡记录
func (s *Service) PunchRecBetween(c context.Context, arg *model.ArgDateBetween) ([]*model.PunchRec, error) {
	return s.dao.QueryPunchRecBetween(c, arg.From, arg.To)
}

// AddPunchRec 添加一条打卡记录
func (s *Service) AddPunchRec(c context.Context, record *model.PunchRec) error {
	if !record.IsTemperatureOK {
		go s.PunchRecWarn(c, record)
	}
	return s.dao.AddPunchRec(c, record)
}

// DeletePunchRec 永久删除一条打卡记录
func (s *Service) DeletePunchRec(c context.Context, id int64) error {
	return s.dao.DeletePunchRec(c, id)
}

// PunchRecWarn 给管理员发提示邮件
func (s *Service) PunchRecWarn(c context.Context, record *model.PunchRec) error {
	admins, err := s.dao.QueryUsers(c, &model.User{Status: 2})
	if err != nil {
		log.Errorf("query users error(%v)", err)
		return err
	}
	user, err := s.dao.QueryUser(c, &model.User{UID: record.UID})
	if err != nil {
		log.Errorf("query user error(%v) uid(%s)", err, record.UID)
		return err
	}
	matrix := struct {
		*model.User
		Rec *model.PunchRec
	}{
		User: user,
		Rec:  record,
	}

	var buf bytes.Buffer
	if err := s.tmpl.Execute(&buf, matrix); err != nil {
		log.Errorf("execute template error(%v)", err)
		return err
	}

	addrs := make([]string, 0, len(admins.Users))
	for _, admin := range admins.Users {
		addrs = append(addrs, admin.Email)
	}
	if err = s.mailer.Send(fmt.Sprintf("%s打卡异常", record.Name), buf.String(), addrs); err != nil {
		log.Errorf("send warning email error(%v) uid(%s)", err, record.UID)
	}
	return err
}

func yesOrNo(b bool) string {
	if b {
		return "是"
	}
	return "否"
}
