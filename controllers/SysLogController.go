package controllers

import (
	"gofox/models"
	"strconv"
)

type SysLogController struct {
	SysBaseController
}

func (c *SysLogController) Prepare() {
	c.SysBaseController.Prepare()
}

func (c *SysLogController) GetSysLogListByPage() {
	rowsNum := 10
	order := map[string]string{"log_id": "desc"}
	pageNum,_ := strconv.Atoi(c.Input().Get("page_num"))

	where := make(map[string]string)
	where["user_name"] 	= c.Input().Get("user_name")
	where["start_time"] = c.Input().Get("start_time")
	where["end_time"] 	= c.Input().Get("end_time")

	dataList, totalRows := models.GetSysLogListByPage(where, pageNum, rowsNum, order)
	c.PageInfo(where,totalRows, rowsNum)

	c.Data["DataList"] 	= dataList
	c.TplName = "syslog/list_sys_log.html"
}

func (c *SysLogController) DeleteSysLog() {
	id := c.Input().Get("id")
	_, err := models.DeleteSysLog(id)
	if err != nil {
		c.DisplayJson(0, "修改失败", err.Error())
	}
	c.DisplayJson(1, "删除成功", c.URLFor("SysLogController.GetSysLogListByPage"))
}
