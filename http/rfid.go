package http

import (
	"github.com/mivinci/abc"
	"github.com/mivinci/kpt/model"
)

func addRFIDRec(c abc.Context) {
	arg := &model.RFIDRec{}
	if err := c.Bind(arg); err != nil {
		return
	}
	c.JSON(nil, svc.AddRFIDRec(c, arg))
}

func rfidRec(c abc.Context) {
	arg := &model.RFIDRec{}
	if err := c.Bind(arg); err != nil {
		return
	}
	c.JSON(svc.RFIDRec(c, arg))
}

func rfidToday(c abc.Context) {
	arg := &model.ArgRFIDRec{}
	if err := c.Bind(arg); err != nil {
		return
	}
	c.JSON(svc.RFIDRecToday(c, arg))
}
