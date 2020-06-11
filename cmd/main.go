package main

import (
	"fmt"
	"os"

	"github.com/mivinci/abc/log"
	"github.com/mivinci/kpt/conf"
	"github.com/mivinci/kpt/http"
)

var (
	env string
)

func main() {
	if env = os.Getenv("KPTENV"); env == "production" {
		log.Init(log.WithFile("server.log"))
		fmt.Printf("env: %s\n", env)
	} else {
		log.Init()
	}
	http.Init(conf.New())
}
