package dao

import (
	"context"

	"github.com/mivinci/kpt/model"
)

// SelPunchRecByUID selects
func (d *Dao) SelPunchRecByUID(c context.Context, uid string) (recs []*model.PunchRec, err error) {
	return recs, d.orm.Where("uid=?", uid).Find(&recs).Error
}

// SelPunchRecBetween selects
func (d *Dao) SelPunchRecBetween(c context.Context, from, to string) (recs []*model.PunchRec, err error) {
	return recs, d.orm.Where("ctime BETWEEN ? AND ?", from, to).Find(&recs).Error
}

// InsPunchRec inserts
func (d *Dao) InsPunchRec(c context.Context, rec *model.PunchRec) error {
	return d.orm.Create(rec).Error
}
