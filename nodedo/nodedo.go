//之所以把他从nodedocontroller包中提出来是因为发现alarmcontroller和nodedocontroller都需要引用nodedo这个对象
package nodedo

import(
	"github.com/mitchellh/mapstructure"
	
	"github.com/ziyouzy/mylib/model"

)

func NewNodeDo(typename string, v interface{})NodeDo{
	switch typename{
	case "bool":
		var bdo BoolenNodeDo
		mapstructure.Decode(v, &bdo)
		return &bdo
	case "int":
		var ido IntNodeDo
		mapstructure.Decode(v, &ido)
		return &ido
	case "float":
		var fdo FloatNodeDo
		mapstructure.Decode(v, &fdo)
		return &fdo
	case "common", "string":
		var cdo CommonNodeDo 
		mapstructure.Decode(v, &cdo)
		return &cdo
	default:
		return nil
	}
}

type NodeDo interface{
	UpdateOneNodeDo(string, string)
	GetJson()[]byte
	PrepareSMSAlarm()string
	PrepareMYSQLAlarm(*model.AlarmEntity)
}