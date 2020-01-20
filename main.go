package main

import (
	"github.com/gogf/gf/frame/g"
	_ "gmanager/boot"
	_ "gmanager/router"
)

func main() {
	g.Server().Run()
}
