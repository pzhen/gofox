package controllers

import (
	"gofox/models"
	"strconv"
)

type SysRoleController struct {
	SysBaseController
}

func (c *SysRoleController) Prepare() {
	c.SysBaseController.Prepare()
}

func (c *SysRoleController) GetSysRoleListByPage() {
	rowsNum := 10
	order := map[string]string{"role_id": "desc"}
	pageNum,_ := strconv.Atoi(c.Input().Get("page_num"))

	where := make(map[string]string)
	where["role_name"] 	= c.Input().Get("role_name")
	where["start_time"] = c.Input().Get("start_time")
	where["end_time"] 	= c.Input().Get("end_time")

	dataList, totalRows := models.GetSysRoleListByPage(where, pageNum, rowsNum, order)
	c.PageInfo(where,totalRows, rowsNum)

	c.Data["RoleList"] = dataList
	c.TplName = "sysrole/list_sys_role.html"
}

func (c *SysRoleController) FormSysRole() {
	RoleId := c.Input().Get("role_id")
	Id, _ := strconv.Atoi(RoleId)
	RoleRow := models.GetSysRoleById(Id)

	//获取权限
	PowerList := models.GetSysRoleMenuActionMap(RoleId)

	// 所有菜单
	MenuList := models.GetSysMenuList()

	c.Data["RoleRow"] = RoleRow
	c.Data["MenuList"] = MenuList
	c.Data["PowerList"] = PowerList
	c.TplName = "sysrole/form_sys_role.html"
}

func (c *SysRoleController) SaveSysRole() {
	r := &models.SysRole{}
	if err := c.ParseForm(r); err != nil {
		c.DisplayJson(0, "数据解析失败", err)
	}

	if _, err := models.SaveSysRole(r); err != nil {
		c.DisplayJson(0, "保存失败", err)
	}

	c.DisplayJson(1, "保存成功", c.URLFor("SysRoleController.GetSysRoleListByPage"))
}

func (c *SysRoleController) ModifySysRoleStatus() {
	ids := c.Input().Get("id")
	roleStatus, _ := strconv.Atoi(c.Input().Get("status"))
	_, err := models.ModifySysRoleStatus(ids, roleStatus)
	if err != nil {
		c.DisplayJson(0, "修改失败", err.Error())
	}
	c.DisplayJson(1, "修改成功", c.URLFor("SysRoleController.GetSysRoleListByPage"))
}

func (c *SysRoleController) DeleteSysRole() {
	ids := c.Input().Get("id")
	_, err := models.DeleteSysRole(ids)
	if err != nil {
		c.DisplayJson(0, "修改失败", err.Error())
	}
	c.DisplayJson(1, "删除成功", c.URLFor("SysRoleController.GetSysRoleListByPage"))
}
