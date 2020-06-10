package http

import (
	"github.com/mivinci/abc"
	"github.com/mivinci/kpt/model"
)

func punchRec(c abc.Context) {
	arg := &model.PunchRec{}
	if err := c.Bind(arg); err != nil {
		return
	}
	c.JSON(svc.PunchRec(c, arg))
}

func punchRecBetween(c abc.Context) {
	arg := &model.ArgDateBetween{}
	if err := c.Bind(arg); err != nil {
		return
	}
	c.JSON(svc.PunchRecBetween(c, arg))
}

func addPunchRec(c abc.Context) {
	arg := &model.PunchRec{}
	if err := c.Bind(arg); err != nil {
		return
	}
	c.JSON(nil, svc.AddPunchRec(c, arg))
}

func deletePunchRec(c abc.Context) {
	arg := new(struct {
		ID int64 `form:"id"`
	})
	if err := c.Bind(arg); err != nil {
		return
	}
	c.JSON(nil, svc.DeletePunchRec(c, arg.ID))
}
