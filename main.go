package main

import (
	"github.com/astaxie/beego"

	_ "gofox/routers"
	_ "gofox/sysinit"
	"time"
)

func convertM(in int) (out int64) {
	out = int64(in) / 1000000
	return
}

func convertT(in uint) (out string) {
	tm := time.Unix(int64(in), 0)
	out = tm.Format("2006-01-02 15:04:05")
	return
}

func main() {
	beego.AddFuncMap("convertm", convertM)
	beego.AddFuncMap("convertt", convertT)

	beego.Run()
}

