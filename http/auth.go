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

func addUser(c abc.Context) {
	arg := &model.User{}
	if err := c.Bind(arg); err != nil {
		return
	}
	c.JSON(nil, svc.AddUser(c, arg))
}

func users(c abc.Context) {
	arg := &model.User{}
	if err := c.Bind(arg); err != nil {
		return
	}
	c.JSON(svc.QueryUsers(c, arg))
}

func updateUser(c abc.Context) {
	arg := &model.User{}
	if err := c.Bind(arg); err != nil {
		return
	}
	c.JSON(nil, svc.UpdateUser(c, arg))
}

func deleteUser(c abc.Context) {
	arg := new(struct {
		UID string `form:"uid"`
	})
	if err := c.Bind(arg); err != nil {
		return
	}
	c.JSON(nil, svc.DeleteUser(c, arg.UID))
}
