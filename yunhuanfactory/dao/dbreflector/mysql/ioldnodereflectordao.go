//设计该接口的目的是让版本更新时方便进行新旧版本的更新
//这个接口内的函数对负责版本更新者而言就好比说明书
//设计这个接口和设计实现了这个接口的结构体属于同一层的工作内容
package mysql

import (
	mysqldbreflectorlv1 "github.com/ziyouzy/mylib/yunhuanfactory/dbreflector/mysql/lv1/impl"
)

type IOldNodeReflectorDao interface{
	OldNodeEntityFromMySql(string) *mysqldbreflectorlv1.OldNodeReflectorImpl
}