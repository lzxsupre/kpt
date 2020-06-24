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
	e.GET("/code", code)
	e.GET("/apptoken", appToken)

	ac := e.NewGroup("/ac", au.ServeHTTP)
	ac.POST("/scan", addScanRec)
	ac.GET("/scan/between", scanRecBetween)
	ac.GET("/scan", scanRec)
	ac.DELETE("/scan", deleteScanRec)

	ac.POST("/rfid", addRFIDRec)
	ac.POST("/temp", addTempRec)
	ac.GET("/rfid", rfidRec)
	ac.GET("/temp", tempRec)

	pr := e.NewGroup("/pr", au.ServeHTTP)
	pr.POST("/punch", addPunchRec)
	pr.GET("/punch/between", punchRecBetween)
	pr.GET("/punch", punchRec)
	pr.DELETE("/punch", deletePunchRec)

	auth := e.NewGroup("/auth", au.ServeHTTP)
	auth.GET("/user", users)
	auth.PUT("/user", updateUser)
	auth.POST("/user", addUser)
	auth.DELETE("/user", deleteUser)
	auth.GET("/app", apps)
	auth.POST("/app", addApp)
}

func register(e *abc.Engine) {
	e.Register(func() {
		svc.Close()
		aladin.Close()
	})
}
