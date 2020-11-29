package nodedo

import(
	"github.com/mitchellh/mapstructure"
	
	"github.com/ziyouzy/mylib/mysql"

)

func NewNodeDo(datatype string, v interface{})NodeDo{
	switch datatype{
	case "bool","boolen":
		var nd BoolenNodeDo
		mapstructure.Decode(v, &nd)
		return &nd
	case "int":
		var nd IntNodeDo
		mapstructure.Decode(v, &nd)
		return &nd
	case "float":
		var nd FloatNodeDo
		mapstructure.Decode(v, &nd)
		return &nd
	case "common", "string":
		var nd CommonNodeDo 
		mapstructure.Decode(v, &nd)
		return &nd
	default:
		return nil
	}
}

type NodeDo interface{
	UpdateOneNodeDo(string, string)
	GetJson()[]byte
	PrepareSMSAlarm()string
	PrepareMYSQLAlarm(*mysql.Alarm)
}