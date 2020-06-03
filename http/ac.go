package http

import (
	"net/http"

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

func scanRecByID(c abc.Context) {
	uid, ok := c.Params["uid"]
	if !ok {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}
	c.JSON(svc.ScanRecByID(c, uid))
}

func scanRecBetween(c abc.Context) {
	arg := new(model.ArgDateBetween)
	if err := c.Bind(arg); err != nil {
		return
	}
	c.JSON(svc.ScanRecBetween(c, arg))
}
