package service

import (
	"context"

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
	return s.dao.AddPunchRec(c, rec)
}

// DeletePunchRec 永久删除一条打卡记录
func (s *Service) DeletePunchRec(c context.Context, id int64) error {
	return s.dao.DeletePunchRec(c, id)
}
