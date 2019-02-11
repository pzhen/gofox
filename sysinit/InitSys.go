package sysinit

import(
	"github.com/astaxie/beego"
	"gofox/controllers"
)

func init() {
	//启用Session
	beego.BConfig.WebConfig.Session.SessionOn = true
	beego.BConfig.WebConfig.Session.SessionProvider="file"
	beego.BConfig.WebConfig.Session.SessionProviderConfig = "./tmp"
	//初始化数据库
	InitDatabase()

	beego.ErrorController(&controllers.ErrorController{})
}
