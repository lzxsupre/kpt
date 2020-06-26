package service

import (
	"context"

	"github.com/mivinci/abc/log"
	"github.com/mivinci/kpt/model"
)

const (
	chkOut = 1 + iota
	chkIn
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

//StatUserWithNoTempRecToday get_today_user_nontemp
func (s *Service) StatUserWithNoTempRecToday(c context.Context, warn bool) (resp []*model.User, err error) {
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

	var addrs []string

	for _, user := range users.Users {
		if _, ok := recs[user.UID]; !ok {
			resp = append(resp, user)
			addrs = append(addrs, user.Email)
		}
	}

	if warn {
		go func() {
			if e := s.mailer.Send("提醒未测体温", "同学您好，您今日未测体温，请尽快去门禁处测量体温", addrs); e != nil {
				log.Errorf("send email error(%v)\n", err)
			}
		}()
	}
	return
}

// StatLastRFIDRec 获取当天刷门禁的用户的最后记录
func (s *Service) StatLastRFIDRec(c context.Context) (stat *model.StatLastRFIDRec, err error) {
	records, err := s.RFIDRecToday(c, &model.ArgRFIDRec{})
	if err != nil {
		log.Errorf("query rfid record error(%v)\n", err)
		return
	}

	stat = &model.StatLastRFIDRec{
		Total: len(records),
		Out:   make(map[string]*model.RFIDRec),
		In:    make(map[string]*model.RFIDRec),
	}

	for _, record := range records {
		switch record.Type {
		case chkIn:
			if rec, ok := stat.In[record.UID]; ok {
				if rec.Ctime < record.Ctime {
					rec.Ctime = record.Ctime
				}
				continue
			}
			stat.In[record.UID] = record
			continue
		case chkOut:
			if rec, ok := stat.Out[record.UID]; ok {
				if rec.Ctime < record.Ctime {
					rec.Ctime = record.Ctime
				}
			}
			stat.Out[record.UID] = record
			continue
		default:
		}
	}
	return
}

// StatUserWithNoRFIDRecInToday 获取当天未刷进的用户
func (s *Service) StatUserWithNoRFIDRecInToday(c context.Context, warn bool) (stat []*model.StatUserNoCheckIn, err error) {
	var (
		// 未刷入用户的 UID
		uids    []string
		addrs   []string
		users   []*model.User
		lastRec *model.StatLastRFIDRec
	)
	if lastRec, err = s.StatLastRFIDRec(c); err != nil {
		return
	}
	for uid, out := range lastRec.Out {
		if in, ok := lastRec.In[uid]; !ok {
			uids = append(uids, out.UID)

		} else if in.Ctime < out.Ctime {
			uids = append(uids, out.UID)
		}
	}
	if users, err = s.dao.QueryUsersByUIDs(c, uids); err != nil {
		return
	}
	for _, user := range users {
		stat = append(stat, &model.StatUserNoCheckIn{
			User:         user,
			LastCheckout: lastRec.Out[user.UID].Ctime,
		})
		addrs = append(addrs, user.Email)
	}
	if warn {
		go func() {
			if e := s.mailer.Send("未返校提醒", "同学您好，发现您在当天23:00点前没有返校打卡记录，请在宵禁前返校！", addrs); e != nil {
				log.Errorf("send email error(%v)\n", err)
			}
		}()
	}
	return
}
