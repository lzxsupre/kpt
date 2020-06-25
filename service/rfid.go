package service

import (
	"context"
	"fmt"
	"time"

	"github.com/mivinci/abc/log"
	"github.com/mivinci/kpt/model"
)

const (
	rfidIn = 1
)

// RFIDRec 获取门禁记录
func (s *Service) RFIDRec(c context.Context, record *model.RFIDRec) ([]*model.RFIDRec, error) {
	return s.dao.QueryRFIDRec(c, record)
}

// AddRFIDRec 添加一条门禁记录
func (s *Service) AddRFIDRec(c context.Context, record *model.RFIDRec) error {
	if record.Type == rfidIn {
		// 若为刷出则发送邮件
		go s.RFIDOutNotaionSend(c, record.UID)
	}
	return s.dao.AddRFIDRec(c, record)
}

// RFIDOutNotaionSend 发送门禁刷出提醒
func (s *Service) RFIDOutNotaionSend(c context.Context, uid string) error {
	user, err := s.dao.QueryUser(c, &model.User{UID: uid})
	if err != nil {
		log.Errorf("Query user error(%v) uid(%s)\n", err, uid)
		return err
	}

	body := fmt.Sprintf("检测到您在 %s 刷卡外出。在外注意防护，戴好口罩 :)", time.Now().Format("2006-01-02 15:04:05"))

	if err = s.mailer.Send("您已成功刷卡外出", body, []string{user.Email}); err != nil {
		log.Errorf("send email error(%v)\n", err)
		return err
	}
	return nil
}
