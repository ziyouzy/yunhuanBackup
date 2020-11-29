package main

import(
	"fmt"
	//"time"
	//"encoding/json"

	//"github.com/ziyouzy/mylib/tcp"
	//"github.com/ziyouzy/mylib/model"
	"github.com/ziyouzy/mylib/service"
	//"github.com/ziyouzy/mylib/protocol"
	"github.com/ziyouzy/mylib/conf"
	//"github.com/ziyouzy/mylib/physicalnode"
	//"github.com/ziyouzy/mylib/connserver"
)


func main(){
	//先创建的先消费原则GenerateNodeDoCh()
	conf.Load()
	
	fmt.Println("service start:")
	rawCh :=service.RawCh()//rawCh的创建者a-b
	service.ConnServerListenAndCollect()//rawCh的生产者
	go func(){
		for raw := range rawCh{
			fmt.Println(raw)
		}
	}()

	//physicalNodeCh := service.RawChToPhysicalNodeCh(rawCh)//rawCh的消费者和physicalNodeCh的创建者和生产者a-b
	//go func(){
	//	for pn := range physicalNodeCh{
	//		_ =pn
	//	}
	//}()
	//service.UpdateEveryExsitNodeDoTemplate(physicalNodeCh)//physiaclNodeCh的消费者
	//nodeDoCh :=service.NodeDoCh()//nodeDoCh的创建和生产者a-b-?
	

	//service.ActionAlarmFiler(nodeDoCh)//nodeDoCh的消费者，以及之后那三个的生产者，chan bool的消费者
	//service.ActionAlarmSMSSender()//消费
	//service.ActionAlarmMYSQLCreater()//消费

	select{}
	//service.TickerSendModbusToNouthBound(2)//非流水线设计模式
	//service.TickerSendModbusToNouthBound_RangeTest()
	
	//for{}
	// for nodedo := range nodeDoCh{
	// 	fmt.Println(nodedo)
	// }
}