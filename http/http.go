package http

import (
	"github.com/mivinci/abc"
	"github.com/mivinci/abc/middlewares/auth"
	"github.com/mivinci/abc/middlewares/cors"
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
	middleware(e)
	register(e)
	routers(e)
	e.Start(":8000")
}

func services(c *conf.Config) {
	svc = service.New(c)
	au = auth.New(c.Key.Secret)
}

func middleware(e *abc.Engine) {
	e.Use(cors.Default())
}

func routers(e *abc.Engine) {
	e.GET("/token", token)

	ac := e.NewGroup("/ac")
	ac.POST("/scan", addScanRec)
	ac.GET("/scan/between", scanRecBetween)
	ac.GET("/scan", scanRec)
	ac.DELETE("/scan", deleteScanRec)

	pr := e.NewGroup("/pr")
	pr.POST("/punch", addPunchRec)
	pr.GET("/punch/between", punchRecBetween)
	pr.GET("/punch", punchRec)
	pr.DELETE("/punch", deletePunchRec)

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
