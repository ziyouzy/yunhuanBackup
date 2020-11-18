package service

import(
	"fmt"

	"github.com/ziyouzy/mylib/alarmcontroller"
	//"github.com/ziyouzy/mylib/model"
	"github.com/ziyouzy/mylib/nodedo"
)


func AlarmMYSQL()/*chan *model.AlarmEntity*/{
	ch :=alarmcontroller.AlarmMYSQLEntityCh()
	go func(){
		for entity := range ch{
			fmt.Println("准备通过gorm录入mysql,entity:",entity)
			//model.DB.Create(&entity)
		}
	}()
}

func AlarmSMS()/*chan []byte*/{
	ch :=alarmcontroller.AlarmSMSbyteCh()
	go func(){
		for b := range ch{
			fmt.Println("准备通过serial发送,b:",b)
		}
	}()
}

//在realizealnodedo中可以拿到NodeDoCh，
//如上所示，Filter方法如发现超限的nodedo,就会立刻把相关数据放入上所管理的两个管道
//因此只要确保AlarmMYSQL()和AlarmSMS()在NodeDoCh的for -range前执行即可
//原因在于需要确保先创建AlarmMYSQLEntityCh()和AlarmSMSbyteCh()
func AlarmFiler(nd nodedo.NodeDo){
	alarmcontroller.Filter(nd)
}

