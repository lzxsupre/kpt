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
