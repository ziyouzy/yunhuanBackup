package model

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

/*名称；数值；正常值；单位；告警内容；时间*/
type AlarmEntity struct {
	gorm.Model

	PresentSouthID int `gorm:"comment:'从当时json配置中所读取到的设备ID'"`
	PresentSouthBound string `gorm:"comment:'从当时json配置中所读取到的北向通信目标地址'"`

	Name string 	`gorm:"comment:'异常节点的名称'"`
	Value string 	`gorm:"comment:'超限当时的数值'"`
	Unit string 	`gorm:"comment:'单位'"`
	Content string 	`gorm:"comment:'告警内容'"`
}