package http

import (
	"github.com/mivinci/abc"
	"github.com/mivinci/abc/middlewares/auth"
	"github.com/mivinci/aladin"
	"github.com/mivinci/kpt/conf"
	"github.com/mivinci/kpt/service"
)

var (
	svc *service.Service
	au  *auth.Auth
)

// Init inits
func Init(c *conf.Config) {
	services(c)
	e := abc.New()
	register(e)
	routers(e)
	e.Start(":8000")
}

func services(c *conf.Config) {
	svc = service.New(c)
	au = auth.New(c.Key.Secret)
}

func routers(e *abc.Engine) {
	e.GET("/token", token)

	ac := e.NewGroup("/ac")
	ac.POST("/scan", addScanRec)
	ac.GET("/scan", scanRecBetween)
	ac.GET("/scan/:uid", scanRecByID)

	pr := e.NewGroup("/pr", au.ServeHTTP)
	pr.POST("/punch", addPunchRec)
	pr.GET("/punch", punchRecBetween)
	pr.GET("/punch/:uid", punchRecByUID)

	user := e.NewGroup("/auth")
	user.GET("/user", users)
	user.PUT("/user", updateUser)
	user.POST("/user", addUser)
	user.DELETE("/user", deleteUser)
}

func register(e *abc.Engine) {
	e.Register(func() {
		svc.Close()
		aladin.Close()
	})
}
