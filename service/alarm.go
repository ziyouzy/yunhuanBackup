package service

import(
	"fmt"
	"time"

	"github.com/ziyouzy/mylib/alarmcontroller"
	//"github.com/ziyouzy/mylib/model"
	"github.com/ziyouzy/mylib/nodedo"
)


//是个在主函数中需要放在生产者之前的消费者
func ActionAlarmMYSQLCreater(){
	ch :=alarmcontroller.GenerateMYSQLEntityCh()
	go func(){
	 	for entity := range ch{
	 		_ =entity
	 		fmt.Println(time.Now().Format("20060102150405"),"a,准备通过gorm录入mysql,entity:"/*,entity*/)
	 		//model.DB.Create(&entity)
	 	}
	}()
}

//是个在主函数中需要放在生产者之前的消费者
func ActionAlarmSMSSender(){
	ch := alarmcontroller.GenerateSMSbyteCh()
	go func(){
	 	for b := range ch{
			 bb :=b
			 _ =bb
	 		fmt.Println(time.Now().Format("20060102150405"),"b,准备通过serial发送,b:"/*,b*/)
	 	}
	}()
}

//NodeDoCh的子携程消费者
//Filter实现了NodeDoCh的子线程消费，以及AlarmSMSbyteCh、AlarmMYSQLEntityCh、chan bool的生产
func ActionAlarmFiler(ndch chan nodedo.NodeDo){
	ch := alarmcontroller.Filter(ndch)
	go func(){
		for bo := range ch{
			fmt.Println("c.当前的NodeDo正常吗:",bo)
		}
   }()
}

