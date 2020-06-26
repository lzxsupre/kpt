package http

import (
	"strconv"

	"github.com/mivinci/abc"
)

func statTemp(c abc.Context) {
	c.JSON(svc.StatTempRecToday(c))
}

func statUserWithNoTempRecToday(c abc.Context) {
	c.Request.ParseForm()
	warn, _ := strconv.ParseBool(c.Request.Form.Get("warn"))
	c.JSON(svc.StatUserWithNoTempRecToday(c, warn))
}

func statLastRFIDRec(c abc.Context) {
	c.JSON(svc.StatLastRFIDRec(c))
}

func statUserWithNoRFIDRecInToday(c abc.Context) {
	c.Request.ParseForm()
	warn, _ := strconv.ParseBool(c.Request.Form.Get("warn"))
	c.JSON(svc.StatUserWithNoRFIDRecInToday(c, warn))
}
