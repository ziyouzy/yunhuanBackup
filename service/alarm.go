package service

import(
	"fmt"
	"time"

	"github.com/ziyouzy/mylib/alarmbuilder"
	"github.com/ziyouzy/mylib/nodedo"
)


//是个在主函数中需要放在生产者之前的消费者
func ActionAlarmMYSQLCreater(){
	ch :=alarmbuilder.GenerateMYSQLAlarmCh()
	go func(){
	 	for alarm := range ch{
	 		_ =alarm
	 		fmt.Println(time.Now().Format("20060102150405"),"a,准备通过gorm录入mysql,实体为:"/*alarm*/)
	 		//mysql.DB.Create(&alarm)
	 	}
	}()
}

//是个在主函数中需要放在生产者之前的消费者
func ActionAlarmSMSSender(){
	ch := alarmbuilder.GenerateSMSbyteCh()
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
	alarmbuilder.Filter(ndch)
}

