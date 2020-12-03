//NewNodeDo函数，只要确保json文档的各个字段名与NodeDo四个结构体内部的字段名意义匹配，即可实现功能
//json的字段名都是小写；结构体字段名为驼峰命名法
package nodedo

import(
	"github.com/mitchellh/mapstructure"
	
	"github.com/ziyouzy/mylib/mysql"

)

func NewNodeDo(datatype string, v interface{})NodeDo{
	switch datatype{
	case "bool","boolen","BOOL","BOLLEN":
		var nd BoolenNodeDo
		mapstructure.Decode(v, &nd)
		return &nd
	case "int","INT":
		var nd IntNodeDo
		mapstructure.Decode(v, &nd)
		return &nd
	case "float","FLOAT":
		var nd FloatNodeDo
		mapstructure.Decode(v, &nd)
		return &nd
	case "common", "string","COMMON","STRING":
		var nd CommonNodeDo 
		mapstructure.Decode(v, &nd)
		return &nd
	default:
		return nil
	}
}

type NodeDo interface{
	UpdateOneNodeDo(string, string)
	GetTimeOutSec()int
	UpdateOneNodeDoAndGetTimeOutSec(string,string) int
	TimeOut()

	GetJson()[]byte
	
	PrepareSMSAlarm()string
	PrepareMYSQLAlarm(*mysql.Alarm)
}