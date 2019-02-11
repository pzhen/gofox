package models

import (
	"github.com/astaxie/beego/orm"
	"encoding/json"
	"gofox/utils"
	"strconv"
	"math"
	"time"
	"errors"
)

type SysUser struct {
	Id         int    `orm:"column(user_id);auto" description:"用户id"`
	UserName   string `orm:"column(user_name);size(64)" description:"登录名"`
	NickName   string `orm:"column(nick_name);size(64)" description:"昵称"`
	RoleId     string `orm:"column(role_id);size(64)" description:"角色id"`
	Photo      string `orm:"column(photo);size(128)" description:"头像"`
	Password   string `orm:"column(password);size(32)" description:"密码"`
	Salt       string `orm:"column(salt);size(6)" description:"密码盐值"`
	Email      string `orm:"column(email);size(64)"`
	Mobile     string `orm:"column(mobile);size(32)"`
	CreateTime uint   `orm:"column(create_time)"`
	UpdateTime uint   `orm:"column(update_time);null"`
	LastTime   uint   `orm:"column(last_time)"`
	LastIp     string `orm:"column(last_ip);size(15)"`
	LoginCount uint   `orm:"column(login_count)"`
	UserType   int    `orm:"column(user_type)"`
	UserStatus int    `orm:"column(user_status)"`
}

func init() {
	orm.RegisterModel(new(SysUser))
}

//获取一条角色
func GetSysUserRowById(id int) *SysUser {
	data := new(SysUser)
	if id > 0 {
		orm.NewOrm().QueryTable(Table_Sys_User).Filter("user_id", id).One(data)
	}
	return data
}

func GetUserInfoBySession(s interface{}) *SysUser {
	u := new(SysUser)
	value, ok := s.(string)
	if !ok {
		return u
	}
	json.Unmarshal([]byte(value), &u)
	return u
}

func GetUserMenuBySession(s interface{}) map[int]*UserMenuIterm {
	um := make(map[int]*UserMenuIterm)
	value, ok := s.(string)
	if !ok {
		return um
	}
	json.Unmarshal([]byte(value), &um)
	return um
}

func GetSysUserByUserName(userName string) *SysUser {
	u := new(SysUser)
	userName = utils.TrimString(userName)
	if userName == "" {
		return u
	}
	orm.NewOrm().QueryTable(Table_Sys_User).Filter("user_name__contains", userName).One(u)
	return u
}


//拼接where条件
func GetSysUserListWhereSql(where map[string]string) string {
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
func GetSysUserCount(where map[string]string) int {
	c := struct {
		Count int `orm:"column(count)"`
	}{}

	qb, _ := orm.NewQueryBuilder("mysql")
	qb.Select("count(*) as count")
	qb.From(Table_Sys_User)
	qb.Where(GetSysRoleListWhereSql(where))
	orm.NewOrm().Raw(qb.String()).QueryRow(&c)

	return c.Count
}

//分页列表
func GetSysUserListByPage(where map[string]string, pageNum int, rowsNum int, orderBy map[string]string) ([]*SysUser, int) {
	data := make([]*SysUser, 0)
	start := 0
	if pageNum <= 0 {
	}else {
		start = (int(math.Abs(float64(pageNum))) - 1) * rowsNum
	}

	//获取数据
	qb, _ := orm.NewQueryBuilder("mysql")
	qb.Select("*")
	qb.From(Table_Sys_User)
	qb.Where(GetSysUserListWhereSql(where))
	qb.OrderBy(GetSqlOrderBy(orderBy))
	qb.Limit(rowsNum).Offset(start)
	orm.NewOrm().Raw(qb.String()).QueryRows(&data)

	//获取数量
	totalRows := GetSysUserCount(where)
	return data, totalRows
}


func SaveSysUser(u *SysUser) (int64, error) {
	o := orm.NewOrm()
	u.UpdateTime = uint(time.Now().Unix())
	u.Password = utils.String2md5(u.Password)
	if u.Id > 0 {
		if err := o.Read(&SysUser{Id: u.Id}); err == nil {
			if num, err := o.Update(u) ;err != nil {
				return 0, err
			} else {
				return num, nil
			}
		} else {
			return 0, err
		}
	} else {
		u.CreateTime = uint(time.Now().Unix())
		if id, err := o.Insert(u); err == nil {
			return id, nil
		} else {
			return 0, err
		}
	}
}

//删除角色
func DeleteSysUser(ids string) (int64, error) {
	idArr := utils.StringsSplitToSliceInt(ids, ",")
	if len(idArr) == 0 {
		return 0, errors.New("参数错误")
	}
	num, err := orm.NewOrm().QueryTable(Table_Sys_User).Filter("user_id__in", idArr).Delete()
	return num, err
}

//修改状态
func ModifySysUserStatus(ids string, status int) (int64, error) {
	idArr := utils.StringsSplitToSliceInt(ids, ",")
	if len(idArr) == 0 {
		return 0, errors.New("参数错误")
	}
	return orm.NewOrm().QueryTable(Table_Sys_User).Filter("user_id__in", idArr).Update(orm.Params{
		"user_status": status})
}


