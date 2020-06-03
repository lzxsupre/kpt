package http

import (
	"net/http"

	"github.com/mivinci/abc"
	"github.com/mivinci/kpt/model"
)

func punchRecByUID(c abc.Context) {
	uid, ok := c.Params["uid"]
	if !ok {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}
	c.JSON(svc.PunchRecByUID(c, uid))
}

func punchRecBetween(c abc.Context) {
	arg := &model.ArgDateBetween{}
	if err := c.Bind(arg); err != nil {
		return
	}
	c.JSON(svc.PunchRecBetween(c, arg))
}

func addPunchRec(c abc.Context) {
	rec := &model.PunchRec{}
	if err := c.Bind(rec); err != nil {
		return
	}
	c.JSON(nil, svc.AddPunchRec(c, rec))
}
