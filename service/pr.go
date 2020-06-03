package service

import (
	"context"

	"github.com/mivinci/kpt/model"
)

// PunchRecByUID 按学号获取打卡记录
func (s *Service) PunchRecByUID(c context.Context, uid string) ([]*model.PunchRec, error) {
	return s.dao.SelPunchRecByUID(c, uid)
}

// PunchRecBetween 按时间段获取打卡记录
func (s *Service) PunchRecBetween(c context.Context, arg *model.ArgDateBetween) ([]*model.PunchRec, error) {
	return s.dao.SelPunchRecBetween(c, arg.From, arg.To)
}

// AddPunchRec 添加一条打卡记录
func (s *Service) AddPunchRec(c context.Context, rec *model.PunchRec) error {
	return s.dao.InsPunchRec(c, rec)
}
