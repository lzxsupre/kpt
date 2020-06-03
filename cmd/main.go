package main

import (
	"github.com/mivinci/abc/log"
	"github.com/mivinci/kpt/conf"
	"github.com/mivinci/kpt/http"
)

func main() {
	log.Init()
	http.Init(conf.New())
}
