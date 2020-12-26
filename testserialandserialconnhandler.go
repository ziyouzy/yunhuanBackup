package main

import(
	"fmt"
	"time"
	//"encoding/json"

	//"github.com/ziyouzy/mylib/tcp"
	//"github.com/ziyouzy/mylib/model"
	"github.com/ziyouzy/mylib/service"
	"github.com/ziyouzy/mylib/alarmbuilder"
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

	service.BuildPNCh()
	service.BuildNodeDoCh()//其会从nodedobuilder包获取管道，而nodedobuilder包的初始化已通过conf.Load实现
	service.WatchingViper()//监听配置文件所发生的改变

	go func(){
		for nodedo :=range service.NodeDoCh{
			/** 模仿了time.Unix()方法的设计思路
			  * 第一个参数可传入NodeDo的管道，第二个参数可传入NodoDo个体
			  * 在使用的时候二选一，而不能两个都非nil或两个都为nil
			  */
			alarmbuilder.StartFilter(nil, nodedo)
		}
	}()


	go func(){
		for sms := range alarmbuilder.GetSMSAlarmbyteCh(){
			fmt.Println(string(sms))
		}
	}()


	go func(){
		for mysql := range alarmbuilder.GetMYSQLAlarmEntityCh(){
			fmt.Println(mysql)
		}
	}()




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