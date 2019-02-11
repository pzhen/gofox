package controllers

import (
	"gofox/models"
	"gofox/utils"
	"encoding/json"
	"strconv"
)

type SysUserController struct {
	SysBaseController
}

func (c *SysUserController) Prepare() {
	c.SysBaseController.Prepare()
}

func (c *SysUserController) LoginForm() {
	c.TplName = "sysuser/login_form.html"
}

func (c *SysUserController) LogOut() {
	c.DelSession("UserMenu")
	c.DelSession("UserSession")
	c.DisplayJson(1,"logout success",c.URLFor("SysUserController.LoginForm"))
}

func (c *SysUserController) FormSysUser() {
	userId := c.Input().Get("user_id")
	Id, _ := strconv.Atoi(userId)
	UserRow := models.GetSysUserRowById(Id)

	roleIdArr := utils.StringsSplitToSliceInt(UserRow.RoleId, ",")
	roleList := models.GetSysRoleList()

	c.Data["RoleList"] = roleList
	c.Data["RoleIdArr"] = roleIdArr
	c.Data["DataRow"] = UserRow
	c.TplName = "sysuser/form_sys_user.html"
}

func (c *SysUserController) GetSysUserListByPage() {
	rowsNum := 10
	order := map[string]string{"user_id": "desc"}
	pageNum,_ := strconv.Atoi(c.Input().Get("page_num"))

	where := make(map[string]string)
	where["user_name"] 	= c.Input().Get("user_name")
	where["start_time"] = c.Input().Get("start_time")
	where["end_time"] 	= c.Input().Get("end_time")

	dataList, totalRows := models.GetSysUserListByPage(where, pageNum, rowsNum, order)
	c.PageInfo(where,totalRows, rowsNum)

	c.Data["DataList"] = dataList
	c.TplName = "sysuser/list_sys_user.html"
}

func (c *SysUserController) SaveSysUser() {
	r := &models.SysUser{}
	if err := c.ParseForm(r); err != nil {
		c.DisplayJson(0, "数据解析失败", err)
	}

	if _, err := models.SaveSysUser(r); err != nil {
		c.DisplayJson(0, "保存失败", err)
	}

	c.DisplayJson(1, "保存成功", c.URLFor("SysUserController.GetSysUserListByPage"))
}

func (c *SysUserController) LoginAction() {

	type loginForm struct {
		UserName string `form:"username"`
		Password string `form:"password"`
	}
	u := loginForm{}
	c.ParseForm(&u)

	if len(u.UserName) == 0 || len(u.Password) == 0 {
		c.DisplayStatus(0, "账号密码不能为空!", "")
	}

	u.Password = utils.String2md5(u.Password)
	userInfo := models.GetSysUserByUserName(u.UserName)
	if userInfo.Id > 0 {
		if userInfo.UserStatus != 1 {
			c.DisplayStatus(0, "账号禁用,请联系管理员", "")
		}

		if u.Password != userInfo.Password {
			c.DisplayStatus(0, "密码错误", "")
		}

		// 用户信息
		userSession, _ := json.Marshal(userInfo)
		c.SetSession("UserSession", string(userSession))

		// 菜单权限
		v := models.GetUserMenuByRoleIdArr(*userInfo)
		m, _ := json.Marshal(v)
		c.SetSession("UserMenu", string(m))

		c.DisplayStatus(1, "登录成功,等待跳转", "/sys_home/index")
	} else {
		c.DisplayStatus(0, "账号不存在", "")
	}
}

func (c *SysUserController) ModifySysUserStatus() {
	ids := c.Input().Get("id")
	status, _ := strconv.Atoi(c.Input().Get("status"))
	_, err := models.ModifySysUserStatus(ids, status)
	if err != nil {
		c.DisplayJson(0, "修改失败", err.Error())
	}
	c.DisplayJson(1, "修改成功", c.URLFor("SysUserController.GetSysUserListByPage"))
}

func (c *SysUserController) DeleteSysUser() {
	ids := c.Input().Get("id")
	_, err := models.DeleteSysUser(ids)
	if err != nil {
		c.DisplayJson(0, "修改失败", err.Error())
	}
	c.DisplayJson(1, "删除成功", c.URLFor("SysUserController.GetSysUserListByPage"))
}

