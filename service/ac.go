package service

import (
	"context"
	"strconv"

	"github.com/mivinci/kpt/model"
)

// AddScanRec 添加一项门禁记录
func (s *Service) AddScanRec(c context.Context, arg *model.ArgScanRec) error {
	return s.dao.InsScanRec(c, arg.UID, arg.Cid, strconv.FormatFloat(arg.Tpt, 'f', -1, 64))
}

// ScanRec 获取学号为id的门禁记录
func (s *Service) ScanRec(c context.Context, uid string) ([]*model.ScanRec, error) {
	return s.dao.SelScanRec(c, uid)
}

// ScanRecBetween 获取两个日期间的门禁记录
func (s *Service) ScanRecBetween(c context.Context, arg *model.ArgDateBetween) ([]*model.ScanRec, error) {
	return s.dao.SelScanRecsBetween(c, arg.From, arg.To)
}

// DeleteScanRec 永久删除一条门禁记录
func (s *Service) DeleteScanRec(c context.Context, id int64) error {
	return s.dao.DeleteScanRec(c, id)
}
