package system

import (
	"ferry/models/system"
	"ferry/tools"
	"ferry/tools/app"

	"github.com/gin-gonic/gin"
)

/*
  @Author : lanyulei
*/

func GetInfo(c *gin.Context) {

	var roles = make([]string, 1)
	var roleIds = make([]int, 1)
	roles[0] = tools.GetRoleName(c)

	var permissions = make([]string, 1)
	permissions[0] = "*:*:*"

	var buttons = make([]string, 1)
	buttons[0] = "*:*:*"

	RoleMenu := system.RoleMenu{}
	roleId := tools.GetRoleId(c)
	roleIds[0] = roleId
	RoleMenu.RoleId = roleId

	var mp = make(map[string]interface{})
	mp["roles"] = roles
	mp["roleIds"] = roleIds
	if tools.GetRoleName(c) == "admin" || tools.GetRoleName(c) == "系统管理员" {
		mp["permissions"] = permissions
		mp["buttons"] = buttons
	} else {
		list, _ := RoleMenu.GetPermis()
		mp["permissions"] = list
		mp["buttons"] = list
	}

	sysuser := system.SysUser{}
	sysuser.UserId = tools.GetUserId(c)
	user, err := sysuser.Get()
	if err != nil {
		app.Error(c, -1, err, "")
		return
	}

	mp["introduction"] = " am a super administrator"

	mp["avatar"] = "https://wpimg.wallstcn.com/f778738c-e4f8-4870-b634-56703b4acafe.gif"
	if user.Avatar != "" {
		mp["avatar"] = user.Avatar
	}
	mp["userName"] = user.NickName
	mp["userId"] = user.UserId
	mp["deptId"] = user.DeptId
	mp["name"] = user.NickName

	app.OK(c, mp, "")
}
