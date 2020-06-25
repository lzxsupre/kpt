package service

import (
	"context"

	"github.com/mivinci/abc/log"
	"github.com/mivinci/kpt/model"
)

//StatTempRecToday get_today_temp_data
func (s *Service) StatTempRecToday(c context.Context) (*model.StatTempRec, error) {
	records, err := s.TempRecToday(c, &model.ArgTempRec{})
	if err != nil {
		log.Errorf("query temp record error(%v)\n", err)
		return nil, err
	}
	return s.StatTempRec(records)
}

// StatTempRec 获取体温状态
func (s *Service) StatTempRec(records []*model.TempRec) (stat *model.StatTempRec, err error) {
	var (
		abnormal []*model.UIDTemp
		normal   []*model.UIDTemp
	)
	for _, record := range records {
		uidTemp := &model.UIDTemp{
			UID:  record.UID,
			Temp: record.Temp,
		}
		if s.isTempOK(record.Temp) {
			normal = append(normal, uidTemp)
			continue
		}
		abnormal = append(abnormal, uidTemp)
	}
	stat = &model.StatTempRec{
		Total: len(records),
		Abnormal: &model.UIDTemps{
			Total: len(abnormal),
			Temps: abnormal,
		},
		Normal: &model.UIDTemps{
			Total: len(normal),
			Temps: normal,
		},
	}
	return
}

func (s *Service) isTempOK(t float32) bool {
	if t > s.c.Bus.Threhold.MaxTemp || t < s.c.Bus.Threhold.MinTemp {
		return false
	}
	return true
}

//StatUserWithnoTempRecToday get_today_user_nontemp
func (s *Service) StatUserWithNoTempRecToday(c context.Context) (resp []*model.User, err error) {
	users, err := s.dao.QueryUsers(c, &model.User{})
	if err != nil {
		log.Errorf("query all users error(%v)\n", err)
		return
	}
	records, err := s.TempRecToday(c, &model.ArgTempRec{})
	if err != nil {
		log.Errorf("query temp record of today(%v)\n", err)
		return
	}
	recs := make(map[string]struct{}, len(records))
	for _, record := range records {
		recs[record.UID] = struct{}{}
	}

	for _, user := range users.Users {
		if _, ok := recs[user.UID]; !ok {
			resp = append(resp, user)
		}
	}
	return
}
