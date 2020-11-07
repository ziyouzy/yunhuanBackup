//之所以把他从nodedocontroller包中提出来是因为发现alarmcontroller和nodedocontroller都需要引用nodedo这个对象
package nodedo

import(
	"github.com/ziyouzy/mylib/model"
)

type NodeDo interface{
	UpdateOneNodeDo(string, string)
	GetJson()[]byte
	PrepareSMSAlarm()string
	PrepareMYSQLAlarm(*model.AlarmEntity)
}