package http

import (
	"github.com/mivinci/abc"
	"github.com/mivinci/kpt/model"
)

func token(c abc.Context) {
	arg := &model.ArgAuth{}
	if err := c.Bind(arg); err != nil {
		return
	}
	c.JSON(svc.Token(c, arg))
}

func code(c abc.Context) {
	arg := new(struct {
		UID string `form:"uid"`
	})
	if err := c.Bind(arg); err != nil {
		return
	}
	c.JSON(nil, svc.Code(c, arg.UID))
}
