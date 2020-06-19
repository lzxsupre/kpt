package service

import (
	"context"

	"github.com/mivinci/kpt/model"
)

// RFIDRec 获取门禁记录
func (s *Service) RFIDRec(c context.Context, record *model.RFIDRec) ([]*model.RFIDRec, error) {
	return s.dao.QueryRFIDRec(c, record)
}

// AddRFIDRec 添加一条门禁记录
func (s *Service) AddRFIDRec(c context.Context, record *model.RFIDRec) error {
	return s.dao.AddRFIDRec(c, record)
}
