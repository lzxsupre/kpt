package http

import (
	"github.com/mivinci/abc"
	"github.com/mivinci/kpt/model"
)

func app(c abc.Context) {
	arg := new(struct {
		ID string `form:"appid" validate:"required"`
	})
	if err := c.Bind(arg); err != nil {
		return
	}
	c.JSON(svc.App(c, arg.ID))
}

func apps(c abc.Context) {
	arg := &model.App{}
	if err := c.Bind(arg); err != nil {
		return

	}
	c.JSON(svc.Apps(c, arg))
}

func addApp(c abc.Context) {
	arg := &model.App{}
	if err := c.Bind(arg); err != nil {
		return
	}
	c.JSON(svc.AddApp(c, arg))
}

func appToken(c abc.Context) {
	arg := &model.App{}
	if err := c.Bind(arg); err != nil {
		return
	}
	c.JSON(svc.AppToken(c, arg))
}
