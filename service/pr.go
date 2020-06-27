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
func (s *Service) AddPunchRec(c context.Context, rec *model.PunchRec) error {
	if !rec.IsTemperatureOK {
		go s.PunchRecWarn(c, rec)
	}
	if rec.DidMeetHubei {
		go s.PunchRecWarn(c, rec)
	}
	if rec.HasSymptom {
		go s.PunchRecWarn(c, rec)
	}
	if rec.IsFamilyDiagnosed {
		go s.PunchRecWarn(c, rec)
	}
	if rec.DidMeetDiagnoses {
		go s.PunchRecWarn(c, rec)
	}
	return s.dao.AddPunchRec(c, rec)
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
		log.Errorf("query users error(%v) uid(%s)", err, record.UID)
		return err
	}
	matrix := struct {
		*model.User
		Reason string
	}{
		User:   user,
		Reason: ReasonSelect(record),
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

//ReasonSelect 返回异常原因
func ReasonSelect(record *model.PunchRec) string {
	re := ""
	if record.IsTemperatureOK {
		re = re + "  体温异常"
	}
	if record.DidMeetHubei {
		re = re + "  出现疑似症状"
	}
	if record.HasSymptom {
		re = re + "  已是确诊病例"
	}
	if record.IsFamilyDiagnosed {
		re = re + "  曾有相关病史"
	}
	if record.DidMeetDiagnoses {
		re = re + "  离开学校未归"
	}
	return re
}
