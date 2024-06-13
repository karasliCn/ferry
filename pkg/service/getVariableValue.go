package service

import (
	"encoding/json"
	"ferry/global/orm"
	"ferry/models/process"
	"ferry/models/system"
	"ferry/pkg/logger"
)

/*
  @Author : lanyulei
*/

func GetVariableValueWithWorkOrderId(stateList []interface{}, creator int) (err error) {
	var (
		userInfo system.SysUser
		deptInfo system.Dept
	)

	// 变量转为实际的数据
	for _, stateItem := range stateList {
		stateItemMap := stateItem.(map[string]interface{})
		if stateItemMap["process_method"] == "variable" {
			for processorIndex, processor := range stateItemMap["processor"].([]interface{}) {
				if int(processor.(float64)) == 1 {
					// 创建者
					stateItemMap["processor"].([]interface{})[processorIndex] = creator
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
					stateItemMap["processor"].([]interface{})[processorIndex] = deptInfo.Leader
				}
			}
			stateItemMap["process_method"] = "person"
		}
		if stateItemMap["processed"] != true {
			stateItemMap["processed"] = false
		}
		// no use anymore lhz add to support process_method "template"
		//if stateItem.(map[string]interface{})["process_method"] == "template" {
		//	logger.Debug("should not called here ...")
		//	var fieldId string
		//	fieldId = stateItem.(map[string]interface{})["processor"].([]interface{})[0].(string)
		//
		//	var tplDataList []process.TplData
		//	err := orm.Eloquent.Model(&process.TplData{}).Where(" work_order = ? ", workOrderId).Find(&tplDataList).Error
		//	if err != nil {
		//		logger.Error("tplData not found")
		//	}
		//	tplDataMap := make(map[string]interface{})
		//	for _, tplData := range tplDataList {
		//		tmp := make(map[string]interface{})
		//		err = json.Unmarshal(tplData.FormData, &tmp)
		//		if err != nil {
		//			logger.Error("failed to unmarshal tplData")
		//		}
		//		for key, value := range tmp {
		//			tplDataMap[key] = value
		//		}
		//	}
		//	nextProcessor := make([]interface{}, 0)
		//	nextProcessor = append(nextProcessor, tplDataMap[fieldId])
		//	stateItem.(map[string]interface{})["processor"] = nextProcessor
		//	stateItem.(map[string]interface{})["process_method"] = "person"
		//
		//}
	}

	return
}

func queryWorkOrderFormData(workOrderId int) ([]interface{}, error) {
	var tplDataList []process.TplData
	err := orm.Eloquent.Model(&process.TplData{}).Where(" work_order = ? ", workOrderId).Find(&tplDataList).Error
	if err != nil {
		logger.Error("tplData not found")
		return nil, err
	}
	tplFormList := make([]interface{}, 0)
	for _, tplData := range tplDataList {
		tmp := make(map[string]interface{})
		err = json.Unmarshal(tplData.FormData, &tmp)
		if err != nil {
			logger.Error("failed to unmarshal tplData")
			return nil, err
		}
		tplFormList = append(tplFormList, tmp)
	}
	return tplFormList, nil

}
