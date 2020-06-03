package http

import (
	"github.com/mivinci/abc"
	"github.com/mivinci/aladin"
	"github.com/mivinci/kpt/conf"
	"github.com/mivinci/kpt/service"
)

var svc *service.Service

// Init inits
func Init(c *conf.Config) {
	e := abc.New()
	register(e)
	routers(e)
	services(c)
	e.Start(":8000")
}

func services(c *conf.Config) {
	svc = service.New(c)
}

func routers(e *abc.Engine) {
	ac := e.NewGroup("/ac")
	ac.POST("/scan", addScanRec)
	ac.GET("/scan", scanRecBetween)
	ac.GET("/scan/:uid", scanRecByID)

	pr := e.NewGroup("/pr")
	pr.POST("/punch", addPunchRec)
	pr.GET("/punch", punchRecBetween)
	pr.GET("/punch/:uid", punchRecByUID)
}

func register(e *abc.Engine) {
	e.Register(func() {
		aladin.Close()
		svc.Close()
	})
}
