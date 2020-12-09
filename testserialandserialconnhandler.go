package main

import(
	"fmt"
	"time"
	//"encoding/json"

	//"github.com/ziyouzy/mylib/tcp"
	//"github.com/ziyouzy/mylib/model"
	"github.com/ziyouzy/mylib/service"
	//"github.com/ziyouzy/mylib/protocol"
	"github.com/ziyouzy/mylib/conf"
	"github.com/ziyouzy/mylib/nodedobuilder"
	"github.com/ziyouzy/mylib/physicalnode"
	//"github.com/ziyouzy/mylib/connserver"
)


func main(){
	conf.Load()
	time.Sleep(2 * time.Second)
	fmt.Println("service start:")


	rawCh :=service.RawCh()//rawCh的创建者a-b
	service.ConnServerListenAndCollect()//rawCh的生产者


	physicalNodeCh := physicalnode.RawChToPhysicalNodeCh(rawCh)//rawCh的消费者和physicalNodeCh的创建者和生产者a-b

	
	NodeDoCh =nodedobuilder.GenerateNodeDoCh()


	go func(){
		//只会负责实时维护内部的map[string]NodeDo类型
		//他与NodeDo的定时向下游发送，无论是在Engineing函数内部还是在GenerateNodeDoCh函数内部
		//在设计上做到了互不冲突
		for pn := range physicalNodeCh{
			nodedobuilder.Engineing(pn)
		 }
	}()

	//service.UpdateEveryExsitNodeDoTemplate(physicalNodeCh)//physiaclNodeCh的消费者
	//nodeDoCh :=service.NodeDoCh()//nodeDoCh的创建和生产者a-b-?
	

	//service.ActionAlarmFiler(conf.NodeDoCh)//nodeDoCh的消费者，以及之后那三个的生产者，chan bool的消费者
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