package controllers

import (
	"github.com/astaxie/beego"
	"os"
	"runtime"
)

type SysHomeController struct {
	SysBaseController
}

func (c *SysHomeController) Prepare() {
	c.SysBaseController.Prepare()
}

func (c *SysHomeController) Index() {
	c.TplName = "syshome/index.html"
}

func (c *SysHomeController) ServerInfo() {
	// 系统配置信息
	c.Data["OS"] 			= beego.AppConfig.String("os")
	c.Data["Author"] 		= beego.AppConfig.String("author")
	c.Data["GOPATH"] 		= os.Getenv("GOPATH")
	c.Data["GOVersion"] 	= runtime.Version()
	c.Data["UploadLimit"] 	= beego.AppConfig.String("uploadlimit")
	c.Data["MySqlVersion"] 	= beego.AppConfig.String("mysqlversion")

	c.TplName = "syshome/server_info.html"
}
