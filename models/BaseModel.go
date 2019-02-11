//建议不同模块分开定义
package models

import "strings"

const (
	//SYS 模块
	Table_Sys_User          = "sys_user"
	Table_Sys_Log           = "sys_log"
	Table_Sys_Role          = "sys_role"
	Table_Sys_Menu          = "sys_menu"
	Table_Sys_Role_Menu_Map = "sys_role_menu_map"
)

const (
	//other 模块
)

func GetSqlOrderBy(orderBy map[string]string) string {
	var sort = ""
	for k,v := range orderBy {
		sort += k + " " + v + ","
	}
	sort = strings.TrimRight(sort, ",")
	return sort
}