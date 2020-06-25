package http

import "github.com/mivinci/abc"

func statTemp(c abc.Context) {
	c.JSON(svc.StatTempRecToday(c))
}

func statUserWithNoTempRecToday(c abc.Context) {
	c.JSON(svc.StatUserWithNoTempRecToday(c))
}
