package service

import (
	"encoding/json"
	"ferry/global/orm"
	"ferry/models/process"
	"ferry/models/system"
	"ferry/pkg/constants"
	"ferry/pkg/pagination"
	"ferry/tools"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"strconv"
)

/*
  @Author : lanyulei
  @todo: 添加新的处理人时候，需要修改（先完善功能，后续有时间的时候优化一下这部分。）
*/

type WorkOrder struct {
	Classify int
	GinObj   *gin.Context
}

type Processor struct {
	ProcessorId   int
	ProcessorType string
}

type workOrderInfo struct {
	process.WorkOrderInfo
	Principals   string `json:"principals"`
	StateName    string `json:"state_name"`
	DataClassify int    `json:"data_classify"`
	ProcessName  string `json:"process_name"`
}

type CirculationInfo struct {
	ProcessName    string `json:"process_name"`
	Title          string `json:"title"`
	State          string `json:"state"`
	ProcessMethod  string `json:"process_method"`
	Processor      string `json:"processor"`
	CreateTime     string `json:"create_time"`
	EndTime        string `json:"update_time"`
	SuspendTime    string `json:"suspend_time"`
	ResumeTime     string `json:"resume_time"`
	ProcessorIds   []int  `json:"processor_ids"`
	ProcessorNames string `json:"processor_names"`
	Remarks        string `json:"remarks"`
}

func NewWorkOrder(classify int, c *gin.Context) *WorkOrder {
	return &WorkOrder{
		Classify: classify,
		GinObj:   c,
	}
}

func (w *WorkOrder) PureWorkOrderList() (result interface{}, err error) {
	var (
		workOrderInfoList []workOrderInfo
	)

	db, err := w.buildQuery()

	result, err = pagination.Paging(&pagination.Param{
		C:  w.GinObj,
		DB: db,
	}, &workOrderInfoList, map[string]map[string]interface{}{}, "p_process_info")
	if err != nil {
		err = fmt.Errorf("查询工单列表失败，%v", err.Error())
		return
	}
	return
}

func (w *WorkOrder) PureAllWorkOrderList() (result interface{}, err error) {
	var (
		workOrderInfoList []workOrderInfo
	)

	db, err := w.buildQuery()

	err = db.Find(&workOrderInfoList).Error
	if err != nil {
		err = fmt.Errorf("查询工单列表失败，%v", err.Error())
		return nil, err
	}
	return workOrderInfoList, nil
}

func (w *WorkOrder) buildQuery() (dbObj *gorm.DB, err error) {
	var (
		processorInfo system.SysUser
	)

	personSelectValue := "(JSON_CONTAINS(p_work_order_info.state, JSON_OBJECT('processor', %v)) and JSON_CONTAINS(p_work_order_info.state, JSON_OBJECT('process_method', 'person')))"
	roleSelectValue := "(JSON_CONTAINS(p_work_order_info.state, JSON_OBJECT('processor', %v)) and JSON_CONTAINS(p_work_order_info.state, JSON_OBJECT('process_method', 'role')))"
	departmentSelectValue := "(JSON_CONTAINS(p_work_order_info.state, JSON_OBJECT('processor', %v)) and JSON_CONTAINS(p_work_order_info.state, JSON_OBJECT('process_method', 'department')))"

	title := w.GinObj.DefaultQuery("title", "")
	startTime := w.GinObj.DefaultQuery("startTime", "")
	endTime := w.GinObj.DefaultQuery("endTime", "")
	isEnd := w.GinObj.DefaultQuery("isEnd", "")
	processor := w.GinObj.DefaultQuery("processor", "")
	priority := w.GinObj.DefaultQuery("priority", "")
	creator := w.GinObj.DefaultQuery("creator", "")
	processParam := w.GinObj.DefaultQuery("process", "")
	formData := w.GinObj.DefaultQuery("formData", "")
	db := orm.Eloquent.Model(&process.WorkOrderInfo{}).
		Where("p_work_order_info.title like ?", fmt.Sprintf("%%%v%%", title))

	if startTime != "" {
		db = db.Where("p_work_order_info.create_time >= ?", startTime)
	}
	if endTime != "" {
		db = db.Where("p_work_order_info.create_time <= ?", endTime)
	}
	if isEnd != "" {
		db = db.Where("p_work_order_info.is_end = ?", isEnd)
	}
	if creator != "" {
		db = db.Where("p_work_order_info.creator = ?", creator)
	}
	if processParam != "" {
		db = db.Where("p_work_order_info.process = ?", processParam)
	}
	if formData != "" {
		db = db.Joins("left join p_work_order_tpl_data on p_work_order_tpl_data.work_order = p_work_order_info.id").
			Where("p_work_order_tpl_data.form_data->'$.*' LIKE CONCAT('%',?,'%')", formData).
			Group("p_work_order_info.id")
	}

	// 获取当前用户信息
	switch w.Classify {
	case 1:
		// 待办工单
		// 1. 个人
		personSelect := fmt.Sprintf(personSelectValue, tools.GetUserId(w.GinObj))

		// 2. 角色
		roleSelect := fmt.Sprintf(roleSelectValue, tools.GetRoleId(w.GinObj))

		// 3. 部门
		var userInfo system.SysUser
		err := orm.Eloquent.Model(&system.SysUser{}).
			Where("user_id = ?", tools.GetUserId(w.GinObj)).
			Find(&userInfo).Error
		if err != nil {
			return nil, err
		}
		departmentSelect := fmt.Sprintf(departmentSelectValue, userInfo.DeptId)

		// 4. 变量会转成个人数据
		//db = db.Where(fmt.Sprintf("(%v or %v or %v or %v) and is_end = 0", personSelect, personGroupSelect, departmentSelect, variableSelect))
		db = db.Where(fmt.Sprintf("(%v or %v or %v) and p_work_order_info.is_end = 0", personSelect, roleSelect, departmentSelect))
	case 2:
		// 我创建的
		db = db.Where("p_work_order_info.creator = ?", tools.GetUserId(w.GinObj))
	case 3:
		// 我相关的
		db = db.Where(fmt.Sprintf("JSON_CONTAINS(p_work_order_info.related_person, '%v')", tools.GetUserId(w.GinObj)))
	case 4:
	// 所有工单
	default:
		return nil, fmt.Errorf("请确认查询的数据类型是否正确")
	}
	if processor != "" && w.Classify != 1 {
		err := orm.Eloquent.Model(&processorInfo).
			Where("user_id = ?", processor).
			Find(&processorInfo).Error
		if err != nil {
			return nil, err
		}
		db = db.Where(fmt.Sprintf("(%v or %v or %v) and p_work_order_info.is_end = 0",
			fmt.Sprintf(personSelectValue, processorInfo.UserId),
			fmt.Sprintf(roleSelectValue, processorInfo.RoleId),
			fmt.Sprintf(departmentSelectValue, processorInfo.DeptId),
		))
	}
	if priority != "" {
		db = db.Where("p_work_order_info.priority = ?", priority)
	}

	db = db.Joins("left join p_process_info on p_work_order_info.process = p_process_info.id").
		Select("p_work_order_info.*, p_process_info.name as process_name")

	return db, nil
}

func (w *WorkOrder) WorkOrderList() (result interface{}, err error) {

	var (
		principals        string
		stateList         []map[string]interface{}
		workOrderInfoList []workOrderInfo
		minusTotal        int
	)

	result, err = w.PureWorkOrderList()
	if err != nil {
		return
	}

	for i, v := range *result.(*pagination.Paginator).Data.(*[]workOrderInfo) {
		var (
			stateName string
			//structResult map[string]interface{}
			authStatus bool
		)
		err = json.Unmarshal(v.State, &stateList)
		if err != nil {
			err = fmt.Errorf("json反序列化失败，%v", err.Error())
			return
		}
		if len(stateList) != 0 {
			// 仅待办工单需要验证
			// todo：还需要找最优解决方案
			if w.Classify == 1 {
				//structResult, err = ProcessStructure(w.GinObj, v.Process, v.Id, "")
				//if err != nil {
				//	return
				//}

				authStatus, err = JudgeUserAuthorityWithStateList(w.GinObj, v.Id)
				if err != nil {
					return
				}
				if !authStatus {
					minusTotal += 1
					continue
				}
			} else {
				authStatus = true
			}

			processorListByMethod := make(map[string][]int, 0)
			processorMap := make(map[string]bool)
			if len(stateList) > 1 {
				for _, s := range stateList {
					if s["processed"] != true {
						for _, p := range s["processor"].([]interface{}) {
							pMethod := (s["process_method"]).(string)
							processorId := int(p.(float64))
							key := pMethod + strconv.Itoa(processorId)
							_, ok := processorMap[key]
							if !ok {
								processorListByMethod[pMethod] = append(processorListByMethod[pMethod], processorId)
								processorMap[key] = true
							}
						}
						if len(processorListByMethod) > 0 {
							if len(stateName) > 0 {
								stateName = stateName + ", "
							}
							stateName = stateName + s["label"].(string)
							//break
						}
					}
				}
			}
			if len(processorListByMethod) == 0 {
				processIdList := make([]int, 0)
				for _, v := range stateList[0]["processor"].([]interface{}) {
					processIdList = append(processIdList, int(v.(float64)))
				}
				processorListByMethod[stateList[0]["process_method"].(string)] = processIdList
				stateName = stateList[0]["label"].(string)
			}
			principals, err = GetPrincipal(processorListByMethod, stateList[0]["process_method"].(string))
			if err != nil {
				err = fmt.Errorf("查询处理人名称失败，%v", err.Error())
				return
			}
		}
		workOrderDetails := *result.(*pagination.Paginator).Data.(*[]workOrderInfo)
		workOrderDetails[i].Principals = principals
		workOrderDetails[i].StateName = stateName
		workOrderDetails[i].DataClassify = v.Classify
		if authStatus {
			workOrderInfoList = append(workOrderInfoList, workOrderDetails[i])
		}
	}

	result.(*pagination.Paginator).Data = &workOrderInfoList
	result.(*pagination.Paginator).TotalCount -= minusTotal

	return result, nil
}

func (w *WorkOrder) WorkOrderCirculationList() (result []CirculationInfo, err error) {

	var (
		circulationList []process.CirculationHistory
		cirRes          []CirculationInfo
	)

	allWorkOder, err := w.PureAllWorkOrderList()
	if err != nil {
		return
	}

	workOrderIdList := make([]int64, 0)
	workOrderInfoMap := make(map[int]workOrderInfo)

	for _, v := range allWorkOder.([]workOrderInfo) {
		workOrderIdList = append(workOrderIdList, int64(v.Id))
		workOrderInfoMap[v.Id] = v
	}

	err = orm.Eloquent.Model(&process.CirculationHistory{}).Where(" work_order IN (?)", workOrderIdList).Order("work_order DESC, id DESC").Find(&circulationList).Error
	if err != nil {
		return
	}

	//userIdList := make([]int, 0)
	var currentWorkOrderId int
	var lastEndTime string
	needQueryMap := make(map[string][]int)
	for _, v := range circulationList {
		woInfo, ok := workOrderInfoMap[v.WorkOrder]
		if ok {
			if currentWorkOrderId != v.WorkOrder {
				currentWorkOrderId = v.WorkOrder
				lastEndTime = ""
				state := make([]map[string]interface{}, 0)
				json.Unmarshal(woInfo.State, &state)
				for _, s := range state {
					if s["processed"] == true {
						continue
					}
					if needQueryMap[s["process_method"].(string)] == nil {
						needQueryMap[s["process_method"].(string)] = make([]int, 0)
					}

					cirInfo := CirculationInfo{
						ProcessName:   woInfo.ProcessName,
						ProcessorIds:  []int{},
						ProcessMethod: s["process_method"].(string),
						Title:         woInfo.Title,
						State:         s["label"].(string),
						CreateTime:    woInfo.UpdatedAt.Format(constants.TimeFormat),
						Remarks:       v.Remarks,
					}
					for _, processor := range s["processor"].([]interface{}) {
						cirInfo.ProcessorIds = append(cirInfo.ProcessorIds, int(processor.(float64)))
						needQueryMap[s["process_method"].(string)] = append(needQueryMap[s["process_method"].(string)], int(processor.(float64)))
					}
					//userIdList = append(userIdList, cirInfo.ProcessorIds...)
					susTime, sok := s["suspend_time"].(string)
					if sok {
						cirInfo.SuspendTime = susTime
					}
					resTime, rok := s["resume_time"].(string)
					if rok {
						cirInfo.ResumeTime = resTime
					}
					cirRes = append(cirRes, cirInfo)
				}

			}
			cirInfo := CirculationInfo{
				ProcessName:    woInfo.ProcessName,
				Title:          woInfo.Title,
				State:          v.State,
				Processor:      v.Processor,
				ProcessMethod:  "person",
				EndTime:        v.CreatedAt.Format(constants.TimeFormat),
				ProcessorNames: v.Processor,
			}
			if lastEndTime == "" {
				cirInfo.CreateTime = woInfo.CreatedAt.Format(constants.TimeFormat)
			} else {
				cirInfo.CreateTime = lastEndTime
			}
			lastEndTime = cirInfo.EndTime
			if !v.SuspendTime.IsZero() {
				cirInfo.SuspendTime = v.SuspendTime.Format(constants.TimeFormat)
			}
			if !v.ResumeTime.IsZero() {
				cirInfo.ResumeTime = v.ResumeTime.Format(constants.TimeFormat)
			}
			cirRes = append(cirRes, cirInfo)
		}
	}

	var userInfoList []system.SysUser
	var userInfoMap = make(map[int]system.SysUser)

	var roleInfoList []system.SysRole
	var roleInfoMap = make(map[int]system.SysRole)

	if len(needQueryMap) > 0 {
		userIdList, ok := needQueryMap["person"]
		if ok && len(userIdList) > 0 {
			err = orm.Eloquent.Model(&system.SysUser{}).Where(" user_id IN (?)", userIdList).Find(&userInfoList).Error
			for _, userInfo := range userInfoList {
				userInfoMap[userInfo.UserId] = userInfo
			}
		}
		roleIdList, ok := needQueryMap["role"]
		if ok && len(roleIdList) > 0 {
			err = orm.Eloquent.Model(&system.SysRole{}).Where(" role_id IN (?)", roleIdList).Find(&roleInfoList).Error
			for _, roleInfo := range roleInfoList {
				roleInfoMap[roleInfo.RoleId] = roleInfo
			}
		}
	}

	if err != nil {
		return
	}

	for idx, cirInfo := range cirRes {
		if len(cirInfo.ProcessorIds) > 0 {
			if cirInfo.ProcessMethod == "person" {
				for i, userId := range cirInfo.ProcessorIds {
					if i > 0 {
						cirRes[idx].ProcessorNames = cirInfo.ProcessorNames + ", "
					}
					cirRes[idx].ProcessorNames = cirInfo.ProcessorNames + userInfoMap[userId].NickName
				}
			}
			if cirInfo.ProcessMethod == "role" {
				for i, roleId := range cirInfo.ProcessorIds {
					if i > 0 {
						cirRes[idx].ProcessorNames = cirInfo.ProcessorNames + ", "
					}
					cirRes[idx].ProcessorNames = cirInfo.ProcessorNames + roleInfoMap[roleId].RoleName
				}
			}
		}
	}

	return cirRes, nil
}
