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

	//"github.com/ziyouzy/mylib/connserver"
	//"github.com/ziyouzy/mylib/nodedobuilder"
	//"github.com/ziyouzy/mylib/physicalnode"
	//"github.com/ziyouzy/mylib/connserver"
)


func main(){
	conf.Load()
	time.Sleep(2 * time.Second)
	fmt.Println("service start:")
	

	//rawCh :=connserver.RawCh()
	//physicalNodeCh := physicalnode.RawChToPhysicalNodeCh(rawCh)//rawCh的消费者和physicalNodeCh的创建者和生产者
	//nodedobuilder.Engineing(physicalNodeCh)

	service.BuildPNCh()
	service.BuildNodeDoCh()//其会从nodedobuilder包获取管道，而nodedobuilder包的初始化已通过conf.Load实现
	service.WatchingViper()//监听配置文件所发生的改变

	// for nodedo := range service.NodeDoCh{
	// 	fmt.Println(nodedo)
	// }
	// go func(){
	// 	//只会负责实时维护内部的map[string]NodeDo类型
	// 	//他与NodeDo的定时向下游发送，无论是在Engineing函数内部还是在GenerateNodeDoCh函数内部
	// 	//在设计上做到了互不冲突
	// 	for pn := range physicalNodeCh{
	// 		nodedobuilder.Engineing(pn)
	// 	 }
	// }()

	// go func(){
	// 	for nodedo := range conf.NodeDoCh{
	// 		fmt.Println(nodedo)
	// 	}
	// }

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