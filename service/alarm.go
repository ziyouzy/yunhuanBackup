package service

import(
	"fmt"

	"github.com/ziyouzy/mylib/alarmcontroller"
	//"github.com/ziyouzy/mylib/model"
	"github.com/ziyouzy/mylib/nodedo"
)


//是个在主函数中需要放在生产者之前的消费者
func ActionAlarmMYSQLCreater(){
	ch :=alarmcontroller.GenerateAlarmMYSQLEntityCh()
	go func(){
	 	for entity := range ch{
	 		_ =entity
	 		fmt.Println("a,准备通过gorm录入mysql,entity:"/*,entity*/)
	 		//model.DB.Create(&entity)
	 	}
	}()
}

//是个在主函数中需要放在生产者之前的消费者
func ActionAlarmSMSSender(){
	ch := alarmcontroller.GenerateAlarmSMSbyteCh()
	go func(){
	 	for b := range ch{
	 		_ =b
	 		fmt.Println("b,准备通过serial发送,b:"/*,b*/)
	 	}
	}()
}

//NodeDoCh的子携程消费者
//同时这个子携程也是AlarmSMSbyteCh和AlarmMYSQLCh的生产者
func AlarmFiler(ndch chan nodedo.NodeDo){
	go func(){
		for nd := range ndch{
			alarmcontroller.Filter(nd)
		}
	}()
}

