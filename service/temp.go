package service

import (
	"bytes"
	"context"
	"fmt"

	"github.com/mivinci/abc/log"
	"github.com/mivinci/kpt/model"
)

// TempRec 获取体温数据
func (s *Service) TempRec(c context.Context, record *model.TempRec) ([]*model.TempRec, error) {
	return s.dao.QueryTempRec(c, record)
}

// AddTempRec 添加一条门禁记录
func (s *Service) AddTempRec(c context.Context, record *model.TempRec) error {
	if record.Temp > s.c.Bus.Threhold.MaxTemp || record.Temp < s.c.Bus.Threhold.MinTemp {
		go s.TempRecWarn(c, record)
	}
	return s.dao.AddTempRec(c, record)
}

// TempRecWarn 体温超过阈值，发送提醒邮件给管理员
func (s *Service) TempRecWarn(c context.Context, record *model.TempRec) error {
	admins, err := s.dao.QueryUsers(c, &model.User{Status: 2})
	if err != nil {
		log.Errorf("query admins error(%v)\n", err)
		return err
	}
	user, err := s.dao.QueryUser(c, &model.User{UID: record.UID})
	if err != nil {
		log.Errorf("query user error(%v)\n", err)
		return err
	}
	addrs := make([]string, 0, admins.Total)
	for _, admin := range admins.Users {
		addrs = append(addrs, admin.Email)
	}
	addrs = append(addrs, user.Email)

	// 异常原因
	reason := fmt.Sprintf("体温：%.1f°C，", record.Temp)

	matrix := model.WarnEmail{
		User:   user,
		Reason: reason,
	}

	var buf bytes.Buffer
	if err := s.tmpl.Execute(&buf, &matrix); err != nil {
		log.Errorf("execute html template error(%v)\n", err)
		return err
	}

	if err = s.mailer.Send(fmt.Sprintf("%s体温异常", user.Name), buf.String(), addrs); err != nil {
		log.Errorf("send email error(%v)\n", err)
		return err
	}
	return nil
}
