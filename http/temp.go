package http

import (
	"github.com/mivinci/abc"
	"github.com/mivinci/kpt/model"
)

func addTempRec(c abc.Context) {
	arg := &model.TempRec{}
	if err := c.Bind(arg); err != nil {
		return
	}
	c.JSON(nil, svc.AddTempRec(c, arg))
}

func tempRec(c abc.Context) {
	arg := &model.TempRec{}
	if err := c.Bind(arg); err != nil {
		return
	}
	c.JSON(svc.TempRec(c, arg))
}
