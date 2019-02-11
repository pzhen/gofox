package routers

import (
	"gofox/controllers"
	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.MainController{}, "Get:Index")

	//用户
	beego.Router("/sys_user/login_out", &controllers.SysUserController{}, "Post:LogOut")
	beego.Router("/sys_user/login_form", &controllers.SysUserController{}, "Get:LoginForm")
	beego.Router("/sys_user/login_action", &controllers.SysUserController{}, "Post:LoginAction")

	beego.Router("/sys_user/list_sysuser", &controllers.SysUserController{}, "Get:GetSysUserListByPage")
	beego.Router("/sys_user/form_sysuser", &controllers.SysUserController{}, "Get:FormSysUser")
	beego.Router("/sys_user/save_sysuser", &controllers.SysUserController{}, "Post:SaveSysUser")
	beego.Router("/sys_user/delete_sysuser", &controllers.SysUserController{}, "Get:DeleteSysUser")
	beego.Router("/sys_user/modify_sysuser_status", &controllers.SysUserController{}, "Get:ModifySysUserStatus")


	//首页
	beego.Router("/sys_home/index", &controllers.SysHomeController{}, "Get:Index")
	beego.Router("/sys_home/serverInfo", &controllers.SysHomeController{}, "Get:ServerInfo")

	//菜单
	beego.Router("/sys_menu/list_sysmenu", &controllers.SysMenuController{}, "Get:GetSysMenuList")
	beego.Router("/sys_menu/form_sysmenu", &controllers.SysMenuController{}, "Get:FormSysMenu")
	beego.Router("/sys_menu/save_sysmenu", &controllers.SysMenuController{}, "Post:SaveSysMenu")
	beego.Router("/sys_menu/modify_sysmenu_status", &controllers.SysMenuController{}, "Get:ModifySysMenuStatus")
	beego.Router("/sys_menu/delete_sysmenu", &controllers.SysMenuController{}, "Get:DeleteSysMenu")

	//角色
	beego.Router("/sys_role/list_sysrole", &controllers.SysRoleController{}, "Get:GetSysRoleListByPage")
	beego.Router("/sys_role/form_sysrole", &controllers.SysRoleController{}, "Get:FormSysRole")
	beego.Router("/sys_role/save_sysrole", &controllers.SysRoleController{}, "Post:SaveSysRole")
	beego.Router("/sys_role/modify_sysrole_status", &controllers.SysRoleController{}, "Get:ModifySysRoleStatus")
	beego.Router("/sys_role/delete_sysrole", &controllers.SysRoleController{}, "Get:DeleteSysRole")

	//日志
	beego.Router("/sys_log/list_syslog", &controllers.SysLogController{}, "Get:GetSysLogListByPage")
	beego.Router("/sys_log/delete_syslog", &controllers.SysLogController{}, "Get:DeleteSysLog")
}
