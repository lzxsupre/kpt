package dao

import (
	"context"

	"github.com/mivinci/kpt/model"
)

// AddRFIDRec inserts
func (d *Dao) AddRFIDRec(c context.Context, record *model.RFIDRec) error {
	return d.DB.Create(record).Error
}

// QueryRFIDRec selects
func (d *Dao) QueryRFIDRec(c context.Context, record *model.RFIDRec) (recs []*model.RFIDRec, err error) {
	return recs, d.DB.Model(&model.RFIDRec{}).Where(record).Order("ctime desc").Find(&recs).Error
}
