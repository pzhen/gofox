package models

import (
	"gofox/utils"
	"errors"
	"strconv"
	"strings"

	"github.com/astaxie/beego/orm"
	"math"
	"time"
)

type SysRole struct {
	Id         int    `orm:"column(role_id);auto" description:"角色ID"`
	RoleName   string `orm:"column(role_name);size(60)" description:"角色名称"`
	Intro      string `orm:"column(intro);null" description:"角色介绍"`
	RoleStatus int    `orm:"column(role_status)" description:"状态"`
	CreateTime uint   `orm:"column(create_time)"`
	UpdateTime uint   `orm:"column(update_time)"`
	MenuMap    string `orm:"-"`
}

type SysRoleMenuMap struct {
	Id       int `orm:"column(role_id);auto" description:"角色ID"`
	MenuId   int `orm:"column(menu_id)" description:"菜单id"`
	ActionId int `orm:"column(action_id);size(255);null" description:"操作权限"`
}

func init() {
	orm.RegisterModel(new(SysRole), new(SysRoleMenuMap))
}

//获取一条角色
func GetSysRoleById(id int) *SysRole {
	data := new(SysRole)
	if id > 0 {
		orm.NewOrm().QueryTable(Table_Sys_Role).Filter("role_id", id).One(data)
	}
	return data
}

//获取角色对应菜单以及菜单下方法
func GetSysRoleMenuActionMap(roleIds string) []SysRoleMenuMap {
	data  := make([]SysRoleMenuMap, 0)
	idArr := make([]int,0)
	for _, v := range utils.StringsSplitToSliceInt(roleIds, ",") {
		if row := GetSysRoleById(int(v)); row.RoleStatus == 1 {
			idArr = append(idArr, int(v))
		}
	}
	if len(idArr) == 0 {
		return data
	}
	orm.NewOrm().QueryTable(Table_Sys_Role_Menu_Map).Filter("role_id__in", idArr).All(&data)
	return data
}

//修改角色
func SaveSysRole(m *SysRole) (id int64, err error) {
	var roleInfo SysRole
	roleInfo.Id = m.Id
	roleInfo.Intro = m.Intro
	roleInfo.RoleName = m.RoleName
	roleInfo.RoleStatus = m.RoleStatus
	roleInfo.UpdateTime = uint(time.Now().Unix())

	o := orm.NewOrm()
	if m.Id > 0 {
		roleInfo.CreateTime = m.CreateTime
		id, err = o.Update(&roleInfo)
		if err == nil {
			o.Delete(&SysRoleMenuMap{Id: m.Id})
		}
	} else {
		roleInfo.CreateTime = uint(time.Now().Unix())
		id, err = o.Insert(&roleInfo)
		roleInfo.Id = int(id)
	}

	//关系入库
	MenuMapArr := strings.Split(m.MenuMap, ",")
	for _, v := range MenuMapArr {
		if v == "" {
			continue
		}
		insert := SysRoleMenuMap{}
		insert.Id = roleInfo.Id
		if strings.Contains(v, "-") == true {
			mapArr := strings.Split(v, "-")
			insert.MenuId, _ = strconv.Atoi(mapArr[0])
			insert.ActionId, _ = strconv.Atoi(mapArr[1])
		} else {
			insert.MenuId, _ = strconv.Atoi(v)
			insert.ActionId = 0
		}

		o.Insert(&insert)
	}
	return
}

//拼接where条件
func GetSysRoleListWhereSql(where map[string]string) string {
	var sql = "1=1"
	if v, ok := where["role_name"]; ok && v != "" {
		keywords := utils.TrimString(v)
		sql += " AND role_name like \"%" + keywords + "%\""
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
func GetSysRoleCount(where map[string]string) int {
	c := struct {
		Count int `orm:"column(count)"`
	}{}

	qb, _ := orm.NewQueryBuilder("mysql")
	qb.Select("count(*) as count")
	qb.From(Table_Sys_Role)
	qb.Where(GetSysRoleListWhereSql(where))
	orm.NewOrm().Raw(qb.String()).QueryRow(&c)

	return c.Count
}

//分页列表
func GetSysRoleListByPage(where map[string]string, pageNum int, rowsNum int, orderBy map[string]string) ([]*SysRole, int) {
	data := make([]*SysRole, 0)
	start := 0
	if pageNum <= 0 {
	}else {
		start = (int(math.Abs(float64(pageNum))) - 1) * rowsNum
	}

	//获取数据
	qb, _ := orm.NewQueryBuilder("mysql")
	qb.Select("*")
	qb.From(Table_Sys_Role)
	qb.Where(GetSysRoleListWhereSql(where))
	qb.OrderBy(GetSqlOrderBy(orderBy))
	qb.Limit(rowsNum).Offset(start)
	orm.NewOrm().Raw(qb.String()).QueryRows(&data)

	//获取数量
	totalRows := GetSysRoleCount(where)
	return data, totalRows
}

//获取所有角色
func GetSysRoleList() []*SysRole {
	data := make([]*SysRole, 0)
	qb, _ := orm.NewQueryBuilder("mysql")
	qb.Select("*")
	qb.From(Table_Sys_Role)
	qb.Where("role_status = 1")
	qb.OrderBy("role_id")
	qb.Desc()
	orm.NewOrm().Raw(qb.String()).QueryRows(&data)
	return data
}

//删除角色
func DeleteSysRole(ids string) (int64, error) {
	idArr := utils.StringsSplitToSliceInt(ids, ",")
	if len(idArr) == 0 {
		return 0, errors.New("参数错误")
	}
	num, err := orm.NewOrm().QueryTable(Table_Sys_Role).Filter("role_id__in", idArr).Delete()
	if err != nil {
		orm.NewOrm().QueryTable(Table_Sys_Role_Menu_Map).Filter("role_id__in", idArr).Delete()
		return num, nil
	}
	return 0, err
}

//修改状态
func ModifySysRoleStatus(ids string, roleStatus int) (int64, error) {
	idArr := utils.StringsSplitToSliceInt(ids, ",")
	if len(idArr) == 0 {
		return 0, errors.New("参数错误")
	}
	return orm.NewOrm().QueryTable(Table_Sys_Role).Filter("role_id__in", idArr).Update(orm.Params{
		"role_status": roleStatus})
}
