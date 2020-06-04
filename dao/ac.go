package dao

import (
	"context"
	"database/sql"

	"github.com/mivinci/abc/log"
	"github.com/mivinci/kpt/model"
)

const (
	insScanRec         = "INSERT INTO scan_record(uid,cid,tpt) VALUES(?,?,?)"
	selScanRecsBetween = "SELECT uid,cid,tpt,ctime FROM scan_record WHERE ctime BETWEEN ? AND ?"
	selScanRec         = "SELECT uid,cid,tpt,ctime FROM scan_record WHERE uid=?"
	delScanRec         = "DELETE FROM scan_record WHERE id=?"
)

// InsScanRec inserts
func (d *Dao) InsScanRec(c context.Context, uid, cid, tpt string) (err error) {
	if _, err = d.db.ExecContext(c, insScanRec, uid, cid, tpt); err != nil {
		log.Errorf("InsScanRec.ExecCtx error(%v) uid(%s)", err, uid)
		return
	}
	return nil
}

// SelScanRecsBetween selects
func (d *Dao) SelScanRecsBetween(c context.Context, from, to string) ([]*model.ScanRec, error) {
	return d.selScanRecs(c, selScanRecsBetween, from, to)
}

// SelScanRec selects
func (d *Dao) SelScanRec(c context.Context, uid string) ([]*model.ScanRec, error) {
	return d.selScanRecs(c, selScanRec, uid)
}

func (d *Dao) selScanRecs(c context.Context, query string, args ...interface{}) (ScanRecs []*model.ScanRec, err error) {
	var rows *sql.Rows
	if rows, err = d.db.QueryContext(c, query, args...); err != nil {
		log.Errorf("SelScanRecs.ExecCtx error(%v)", err)
		return
	}
	for rows.Next() {
		ScanRec := new(model.ScanRec)
		if err = rows.Scan(&ScanRec.UID, &ScanRec.CID, &ScanRec.Tpt, &ScanRec.Ctime); err != nil {
			log.Errorf("SelScanRecs.Scan error(%v)", err)
			return
		}
		ScanRecs = append(ScanRecs, ScanRec)
	}
	return
}

// DeleteScanRec deletes
func (d *Dao) DeleteScanRec(c context.Context, id int64) (err error) {
	if _, err = d.db.ExecContext(c, delScanRec, id); err != nil {
		log.Errorf("DeleteScanRec.ExecCtx error(%v) uid(%d)", err, id)
		return
	}
	return
}
