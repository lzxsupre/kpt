package http

import (
	"github.com/mivinci/abc"
	"github.com/mivinci/kpt/model"
)

func addScanRec(c abc.Context) {
	arg := new(model.ArgScanRec)
	if err := c.Bind(arg); err != nil {
		return
	}
	c.JSON(nil, svc.AddScanRec(c, arg))
}

func scanRecBetween(c abc.Context) {
	arg := new(model.ArgDateBetween)
	if err := c.Bind(arg); err != nil {
		return
	}
	c.JSON(svc.ScanRecBetween(c, arg))
}

func scanRec(c abc.Context) {
	arg := new(struct {
		UID string `form:"uid"`
	})
	if err := c.Bind(arg); err != nil {
		return
	}
	c.JSON(svc.ScanRec(c, arg.UID))
}

func deleteScanRec(c abc.Context) {
	arg := new(struct {
		ID int64 `form:"id"`
	})
	if err := c.Bind(arg); err != nil {
		return
	}
	c.JSON(nil, svc.DeleteScanRec(c, arg.ID))
}
