package service

import (
	"errors"
	"ferry/global/orm"
	"ferry/models/system"
	"reflect"
	"strconv"
	"strings"
)

/*
  @Author : lanyulei
  @todo: 添加新的处理人时候，需要修改（先完善功能，后续有时间的时候优化一下这部分。）
*/

func GetPrincipal(processorQueryMap map[string][]int, processorList []string) (principals string, err error) {
	/*
		person 人员
		persongroup 人员组
		department 部门
		variable 变量
	*/
	var allPrincipalList []string

	dataMap := make(map[string]map[int]string)

	for processMethod, idList := range processorQueryMap {
		switch processMethod {
		case "person":
			var userList []system.SysUser
			err = orm.Eloquent.Model(&system.SysUser{}).
				Where("user_id in (?)", idList).Find(&userList).Error
			if err != nil {
				return
			}
			for _, user := range userList {
				if dataMap["person"] == nil {
					dataMap["person"] = make(map[int]string)
				}
				dataMap["person"][user.UserId] = user.NickName
			}
		case "role":
			var roleList []system.SysRole
			err = orm.Eloquent.Model(&system.SysRole{}).
				Where("role_id in (?)", idList).
				Find(&roleList).Error
			if err != nil {
				return
			}
			for _, role := range roleList {
				if dataMap["role"] == nil {
					dataMap["role"] = make(map[int]string)
				}
				dataMap["role"][role.RoleId] = role.RoleName
			}
		case "department":
			var deptList []system.Dept
			err = orm.Eloquent.Model(&system.Dept{}).
				Where("dept_id in (?)", idList).
				Find(&deptList).Error
			if err != nil {
				return
			}
			for _, dept := range deptList {
				if dataMap["department"] == nil {
					dataMap["department"] = make(map[int]string)
				}
				dataMap["department"][dept.DeptId] = dept.DeptName
			}
		case "variable":
			if dataMap["variable"] == nil {
				dataMap["variable"] = make(map[int]string)
			}
			for _, p := range idList {
				switch p {
				case 1:
					dataMap["variable"][1] = "创建者"
				case 2:
					dataMap["variable"][2] = "创建者负责人"
				}
			}
		}
	}

	for _, methodProcessor := range processorList {
		method, processor, found := strings.Cut(methodProcessor, ":")
		if found {
			processorId, err := strconv.Atoi(processor)
			if err == nil {
				allPrincipalList = append(allPrincipalList, dataMap[method][processorId])
			}
		}
	}
	return strings.Join(allPrincipalList, ", "), nil
}

// 获取用户对应
func GetPrincipalUserInfo(stateList []interface{}, creator int) (userInfoList []system.SysUser, err error) {
	var (
		userInfo        system.SysUser
		deptInfo        system.Dept
		userInfoListTmp []system.SysUser // 临时保存查询的列表数据
		processorList   []interface{}
	)

	err = orm.Eloquent.Model(&userInfo).Where("user_id = ?", creator).Find(&userInfo).Error
	if err != nil {
		return
	}

	for _, stateItem := range stateList {

		if reflect.TypeOf(stateItem.(map[string]interface{})["processor"]) == nil {
			err = errors.New("未找到对应的处理人，请确认。")
			return
		}
		stateItemType := reflect.TypeOf(stateItem.(map[string]interface{})["processor"]).String()
		if stateItemType == "[]int" {
			for _, v := range stateItem.(map[string]interface{})["processor"].([]int) {
				processorList = append(processorList, v)
			}
		} else {
			processorList = stateItem.(map[string]interface{})["processor"].([]interface{})
		}

		switch stateItem.(map[string]interface{})["process_method"] {
		case "person":
			err = orm.Eloquent.Model(&system.SysUser{}).
				Where("user_id in (?)", processorList).
				Find(&userInfoListTmp).Error
			if err != nil {
				return
			}
			userInfoList = append(userInfoList, userInfoListTmp...)
		case "role":
			err = orm.Eloquent.Model(&system.SysUser{}).
				Where("role_id in (?)", processorList).
				Find(&userInfoListTmp).Error
			if err != nil {
				return
			}
			userInfoList = append(userInfoList, userInfoListTmp...)
		case "department":
			err = orm.Eloquent.Model(&system.SysUser{}).
				Where("dept_id in (?)", processorList).
				Find(&userInfoListTmp).Error
			if err != nil {
				return
			}
			userInfoList = append(userInfoList, userInfoListTmp...)
		case "variable": // 变量
			for _, processor := range processorList {
				if int(processor.(float64)) == 1 {
					// 创建者
					userInfoList = append(userInfoList, userInfo)
				} else if int(processor.(float64)) == 2 {
					// 1. 查询部门信息
					err = orm.Eloquent.Model(&deptInfo).Where("dept_id = ?", userInfo.DeptId).Find(&deptInfo).Error
					if err != nil {
						return
					}

					// 2. 查询Leader信息
					err = orm.Eloquent.Model(&userInfo).Where("user_id = ?", deptInfo.Leader).Find(&userInfo).Error
					if err != nil {
						return
					}
					userInfoList = append(userInfoList, userInfo)
				}
			}
		}
	}

	return
}
