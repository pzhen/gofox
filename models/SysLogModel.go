package models

import (
	"github.com/astaxie/beego/orm"
	"errors"
	"gofox/utils"
	"math"
	"strconv"
)

type SysLog struct {
	Id         int    `orm:"column(log_id);auto"`
	Url        string `orm:"column(url);size(512)" description:"操作地址"`
	UrlFor     string `orm:"column(urlfor);size(255)" description:"urlFor"`
	UserId     int    `orm:"column(user_id); description:"用户ID"`
	UserName   string `orm:"column(user_name);size(64)" description:"用户名称"`
	FormData   string `orm:"column(form_data);" description:"操作数据"`
	CreateTime uint   `orm:"column(create_time); description:"操作时间"`
}

func init() {
	orm.RegisterModel(new(SysLog))
}

//添加
func AddSysLog(l *SysLog) (int64, error) {
	return orm.NewOrm().Insert(l)
}

//拼接where条件
func GetSysLogListWhereSql(where map[string]string) string {
	var sql = "1=1"
	if v, ok := where["user_name"]; ok && v != "" {
		keywords := utils.TrimString(v)
		sql += " AND user_name like \"%" + keywords + "%\""
	}

	if v, ok := where["start_time"]; ok && v != "" {
		startTime := utils.GetTimestamp(v)
		sql += " AND create_time >= " + strconv.Itoa(int(startTime))
	}

	if v, ok := where["end_time"]; ok && v != "" {
		endTime := utils.GetTimestamp(v)
		sql += " AND create_time <= " + strconv.Itoa(int(endTime))
	}
	return sql
}

//获取数量
func GetSysLogCount(where map[string]string) int {
	c := struct {
		Count int `orm:"column(count)"`
	}{}

	qb, _ := orm.NewQueryBuilder("mysql")
	qb.Select("count(*) as count")
	qb.From(Table_Sys_Log)
	qb.Where(GetSysLogListWhereSql(where))
	orm.NewOrm().Raw(qb.String()).QueryRow(&c)

	return c.Count
}

//分页列表
func GetSysLogListByPage(where map[string]string, pageNum int, rowsNum int, orderBy map[string]string) ([]*SysLog, int) {
	data := make([]*SysLog, 0)
	start := 0
	if pageNum <= 0 {
	}else {
		start = (int(math.Abs(float64(pageNum))) - 1) * rowsNum
	}

	//获取数据
	qb, _ := orm.NewQueryBuilder("mysql")
	qb.Select("*")
	qb.From(Table_Sys_Log)
	qb.Where(GetSysLogListWhereSql(where))
	qb.OrderBy(GetSqlOrderBy(orderBy))
	qb.Limit(rowsNum).Offset(start)
	orm.NewOrm().Raw(qb.String()).QueryRows(&data)

	//获取数量
	totalRows := GetSysLogCount(where)
	return data, totalRows
}

//删除
func DeleteSysLog(ids string) (int64, error) {
	idArr := utils.StringsSplitToSliceInt(ids, ",")
	if len(idArr) == 0 {
		return 0, errors.New("参数错误")
	}
	return orm.NewOrm().QueryTable(Table_Sys_Log).Filter("log_id__in", idArr).Delete()
}
