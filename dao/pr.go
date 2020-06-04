package dao

import (
	"context"

	"github.com/mivinci/kpt/model"
)

// QueryPunchRecByUID selects
func (d *Dao) QueryPunchRecByUID(c context.Context, uid string) (recs []*model.PunchRec, err error) {
	return recs, d.DB.Where("uid=?", uid).Find(&recs).Error
}

// QueryPunchRecBetween selects
func (d *Dao) QueryPunchRecBetween(c context.Context, from, to string) (recs []*model.PunchRec, err error) {
	return recs, d.DB.Where("ctime BETWEEN ? AND ?", from, to).Find(&recs).Error
}

// AddPunchRec inserts
func (d *Dao) AddPunchRec(c context.Context, rec *model.PunchRec) error {
	return d.DB.Create(rec).Error
}
