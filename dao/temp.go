package dao

import (
	"context"

	"github.com/mivinci/kpt/model"
)

// QueryTempRec selects
func (d *Dao) QueryTempRec(c context.Context, record *model.TempRec) (recs []*model.TempRec, err error) {
	return recs, d.DB.Model(&model.TempRec{}).Where(record).Order("ctime desc").Find(&recs).Error
}

// AddTempRec inserts
func (d *Dao) AddTempRec(c context.Context, record *model.TempRec) error {
	return d.DB.Create(record).Error
}
