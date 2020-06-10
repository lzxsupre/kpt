package dao

import (
	"context"

	"github.com/mivinci/kpt/model"
)

// QueryPunchRec selects
func (d *Dao) QueryPunchRec(c context.Context, record *model.PunchRec) (recs []*model.PunchRec, err error) {
	return recs, d.DB.Model(&model.PunchRec{}).Where(record).Order("ctime desc").Find(&recs).Error
}

// QueryPunchRecBetween selects
func (d *Dao) QueryPunchRecBetween(c context.Context, from, to string) (recs []*model.PunchRec, err error) {
	return recs, d.DB.Where("ctime BETWEEN ? AND ?", from, to).Find(&recs).Error
}

// AddPunchRec inserts
func (d *Dao) AddPunchRec(c context.Context, record *model.PunchRec) error {
	return d.DB.Create(record).Error
}

// DeletePunchRec delete
func (d *Dao) DeletePunchRec(c context.Context, id int64) error {
	return d.DB.Where("id=?", id).Delete(&model.PunchRec{}).Error
}
