package process

import (
	"ferry/models/base"
	"ferry/pkg/jsonTime"
	"time"
)

/*
  @Author : lanyulei
*/

// 工单流转历史
type CirculationHistory struct {
	base.Model
	Title         string            `gorm:"column:title; type: varchar(128)" json:"title" form:"title"`                    // 工单标题
	WorkOrder     int               `gorm:"column:work_order; type: int(11)" json:"work_order" form:"work_order"`          // 工单ID
	State         string            `gorm:"column:state; type: varchar(128)" json:"state" form:"state"`                    // 工单状态
	Source        string            `gorm:"column:source; type: varchar(128)" json:"source" form:"source"`                 // 源节点ID
	Target        string            `gorm:"column:target; type: varchar(128)" json:"target" form:"target"`                 // 目标节点ID
	Circulation   string            `gorm:"column:circulation; type: varchar(128)" json:"circulation" form:"circulation"`  // 流转ID
	Status        int               `gorm:"column:status; type: int(11)" json:"status" form:"status"`                      // 流转状态 1 同意， 0 拒绝， 2 其他
	Processor     string            `gorm:"column:processor; type: varchar(45)" json:"processor" form:"processor"`         // 处理人
	ProcessorId   int               `gorm:"column:processor_id; type: int(11)" json:"processor_id" form:"processor_id"`    // 处理人ID
	CostDuration  int64             `gorm:"column:cost_duration; type: int(11)" json:"cost_duration" form:"cost_duration"` // 处理时长
	Remarks       string            `gorm:"column:remarks; type: longtext" json:"remarks" form:"remarks"`                  // 备注
	SuspendTime   jsonTime.JSONTime `gorm:"column:suspend_time" json:"suspend_time" form:"suspend_time"`
	ResumeTime    jsonTime.JSONTime `gorm:"column:resume_time" json:"resume_time" form:"resume_time"`
	IsEffect      int               `gorm:"column:is_effect; type: int(11)" json:"is_effect" form:"is_effect"`
	NodeCreatedAt *time.Time        `gorm:"column:node_created_at" json:"node_created_at" type:"TIMESTAMP NULL" form:"node_created_at"`
}

func (CirculationHistory) TableName() string {
	return "p_work_order_circulation_history"
}
