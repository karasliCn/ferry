package service

import (
	"encoding/json"
	"ferry/global/orm"
	"ferry/models/process"
	"ferry/models/system"
	"ferry/pkg/logger"
	"strconv"
)

/*
  @Author : lanyulei
*/

func GetVariableValueWithWorkOrderId(stateList []interface{}, creator int, workOrderId int) (err error) {
	var (
		userInfo system.SysUser
		deptInfo system.Dept
	)

	// 变量转为实际的数据
	for _, stateItem := range stateList {
		if stateItem.(map[string]interface{})["process_method"] == "variable" {
			for processorIndex, processor := range stateItem.(map[string]interface{})["processor"].([]interface{}) {
				if int(processor.(float64)) == 1 {
					// 创建者
					stateItem.(map[string]interface{})["processor"].([]interface{})[processorIndex] = creator
				} else if int(processor.(float64)) == 2 {
					// 1. 查询用户信息
					err = orm.Eloquent.Model(&userInfo).Where("user_id = ?", creator).Find(&userInfo).Error
					if err != nil {
						return
					}
					// 2. 查询部门信息
					err = orm.Eloquent.Model(&deptInfo).Where("dept_id = ?", userInfo.DeptId).Find(&deptInfo).Error
					if err != nil {
						return
					}

					// 3. 替换处理人信息
					stateItem.(map[string]interface{})["processor"].([]interface{})[processorIndex] = deptInfo.Leader
				}
			}
			stateItem.(map[string]interface{})["process_method"] = "person"
		}
		// lhz todo
		if stateItem.(map[string]interface{})["process_method"] == "template" {
			fieldId := stateItem.(map[string]interface{})["processor"].(string)
			var tplData process.TplData
			err := orm.Eloquent.Model(&process.TplData{}).Where(" work_order = ? ", workOrderId).Find(&tplData).Error
			if err != nil {
				logger.Error("tplData not found")
			}
			tplDataMap := make(map[string]interface{})
			err = json.Unmarshal(tplData.FormData, &tplDataMap)
			if err != nil {
				logger.Error("failed to unmarshal tplData")
			}
			fieldValue := tplDataMap[fieldId].([]interface{})[0].(float64)
			logger.Info("fieldValue" + strconv.FormatFloat(fieldValue, 'f', 0, 64))
			nextProcessor := make([]interface{}, 0)
			nextProcessor = append(nextProcessor, fieldValue)
			stateItem.(map[string]interface{})["processor"] = nextProcessor
			stateItem.(map[string]interface{})["process_method"] = "person"

		}
	}

	return
}

func GetVariableValueForCirculation(stateList []interface{}, h *Handle) (err error) {
	var (
		userInfo system.SysUser
		deptInfo system.Dept
	)
	creator := h.workOrderDetails.Creator

	// 变量转为实际的数据
	for _, stateItem := range stateList {
		if stateItem.(map[string]interface{})["process_method"] == "variable" {
			for processorIndex, processor := range stateItem.(map[string]interface{})["processor"].([]interface{}) {
				if int(processor.(float64)) == 1 {
					// 创建者
					stateItem.(map[string]interface{})["processor"].([]interface{})[processorIndex] = creator
				} else if int(processor.(float64)) == 2 {
					// 1. 查询用户信息
					err = orm.Eloquent.Model(&userInfo).Where("user_id = ?", creator).Find(&userInfo).Error
					if err != nil {
						return
					}
					// 2. 查询部门信息
					err = orm.Eloquent.Model(&deptInfo).Where("dept_id = ?", userInfo.DeptId).Find(&deptInfo).Error
					if err != nil {
						return
					}

					// 3. 替换处理人信息
					stateItem.(map[string]interface{})["processor"].([]interface{})[processorIndex] = deptInfo.Leader
				}
			}
			stateItem.(map[string]interface{})["process_method"] = "person"
		}
		// lhz
		if stateItem.(map[string]interface{})["process_method"] == "template" {
			fieldId := stateItem.(map[string]interface{})["processor"].([]interface{})[0].(string)
			var tplData process.TplData
			err := orm.Eloquent.Model(&process.TplData{}).Where(" work_order = ? ", h.workOrderId).Find(&tplData).Error
			if err != nil {
				logger.Error("tplData not found")
			}
			tplDataMap := make(map[string]interface{})
			err = json.Unmarshal(tplData.FormData, &tplDataMap)
			if err != nil {
				logger.Error("failed to unmarshal tplData")
			}

			fieldValue := tplDataMap[fieldId].(float64)
			logger.Info("fieldValue" + strconv.FormatFloat(fieldValue, 'f', 0, 64))
			nextProcessor := make([]interface{}, 0)
			nextProcessor = append(nextProcessor, fieldValue)
			stateItem.(map[string]interface{})["processor"] = nextProcessor
			stateItem.(map[string]interface{})["process_method"] = "person"

		}
	}

	return
}
