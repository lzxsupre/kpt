package service

import (
	"context"
	"strconv"

	"github.com/mivinci/kpt/model"
)

// AddScanRec 添加一项门禁记录
func (s *Service) AddScanRec(c context.Context, arg *model.ArgScanRec) error {
	return s.dao.InsScanRec(c, arg.UID, arg.CID, strconv.FormatFloat(arg.Tpt, 'f', -1, 64))
}

// ScanRecByID 获取学号为id的打卡记录
func (s *Service) ScanRecByID(c context.Context, uid string) ([]*model.ScanRec, error) {
	return s.dao.SelScanRecByID(c, uid)
}

// ScanRecBetween 获取两个日期间的打卡记录
func (s *Service) ScanRecBetween(c context.Context, arg *model.ArgDateBetween) ([]*model.ScanRec, error) {
	return s.dao.SelScanRecsBetween(c, arg.From, arg.To)
}
