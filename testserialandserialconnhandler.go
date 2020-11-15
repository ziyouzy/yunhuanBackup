package main

import(
	"fmt"
	"time"
	//"encoding/json"

	//"github.com/ziyouzy/mylib/tcp"
	"github.com/ziyouzy/mylib/model"
	//"github.com/ziyouzy/mylib/service"
	//"github.com/ziyouzy/mylib/protocol"
	"github.com/ziyouzy/mylib/conf"
	//"github.com/ziyouzy/mylib/physicalnode"
	"github.com/ziyouzy/mylib/connserver"
)

//var tcphandler pipelineTcpHandler

func main(){
	//数据库也可以在conf.Load()里实例化，不过选在这里只是为了看着清晰一点
	model.ConnectMySQL("yunhuan_api:13131313@tcp(127.0.0.1:3306)/yh?charset=utf8")
	//同上，也可以在conf.Load()里实例化，不过选在这里只是为了看着清晰一点
	connserver.ListenAndGenerateAllRecvCh()
	//time.Sleep(5*time.Second)
	//cs :=connserver.ClientMap()
	//fmt.Println("cs:",cs)
	time.Sleep(8*time.Second)
	fmt.Println("connserver.Test() start:")
	connserver.Test()
	
	//service.TickerSendModbusToNouthBound(2)

	/*
	初始化了如下内容:
	myvipers(饿汉单例模式)
	nodedocontroller(饿汉单例模式)
	alarmcontroller(饿汉单例模式)
	*/
	conf.Load()

	// recvch :=connserver.RecvCh()
	
	// for b := range recvch{
	// 	fmt.Println("字节数组为:   ",b)
	// }
}



	// tcphandler =pipelineTcpHandler{}
	// tcphandler.ConnMap =make(map[string]*tcp.PipelineTcpSocketConn)
	// //serialhandler := pipelineSerialHandler{ProtocolPortsName:[]string{"serial1","serial2","serial3"},
	// //ProtocolPortsBaud:[]int{9600,9600,9600},ProtocolPortsReadTimeout:[]int{5,5,5},ProtocolPortsNeedCRC:[]bool{true,true,true},}
	// tcpCh :=tcphandler.ListenAndGenerateRecvCh()
	// //serialCh :=serialhandler.ListenAndGenerateRecvCh()
	// sendmap := protocol.ProtocolPrepareSendTicketMgr_YunHuan20200924() 
	// TcpSender(sendmap)


	// //其实tcpch已经可以直接作为rawCh来使用了，只不过值后可以用MergeAllCHToRawCH()进行扩展
	// rawCh :=MergeAllCHToRawCH(tcpCh/*,serialCh*/)
	// physicalNodeCh :=PhysicalConvertByProtocol(rawCh)
	// for p :=range physicalNodeCh{
	// 	fmt.Println(p)
	// }
//}

// func TcpSender(sendmap map[string]chan []byte){
// 	for k,msgch :=range sendmap{
// 		//mark 可能会是"192.168.10.2"，"192.168.10.1",或者是"serial"
// 		go func(mark string){
// 			//问题其实是因为range神坑与单独开携程使用了这个k共同导致的，或者说，毕竟range的时候k和v都只有一个副本，你开携程相当于把这一个副本作为参数传递进了这个携程，无论如何，肯定只会是一个值
// 			for msg := range msgch{
// 				if(tcphandler.ConnMap[mark] !=nil){
// 					tcphandler.ConnMap[mark].SendBytes(msg)
// 				}else{
// 					fmt.Println("mark为:",mark,"的设备暂时并不在册于tcphandler.ConnMap中，往设备客户端尽快上线")
// 				}
// 			}
// 		}(k)
// 	}
// }

// func MergeAllCHToRawCH(cs ...chan []byte)chan []byte{
// 	out :=make(chan []byte)

// 	collect := func(in chan []byte){
// 		for b :=range in{
// 			out <-b
// 		}
// 	}

// 	for _, c := range cs {
// 		go collect(c)
// 	}

// 	return out
// }

// func PhysicalConvertByProtocol(ch chan []byte) chan physicalnode.PhysicalNode{
// 	//confnodech := make(chan conf.ConfNode)
// 	physicalnodech :=make(chan physicalnode.PhysicalNode)
// 	go func(){
// 		for b := range ch{
// 			//拿到的字节数组已经不是最原始的了，最原始的是在tcp包内拿到的
// 			//tcp包加工了原始数组，添加了一些类似“头”的字节
// 			physicalNode :=protocol.ProtocolPreparePhysicalNode_YunHuan20200924(b)
// 			physicalnodech<-physicalNode
// 		}
// 	}()
// 	return physicalnodech
// }

// func Separate(confnodech chan conf.ConfNode)(chan []byte, chan []byte, chan []byte, chan []byte,  chan *model.AlarmEntity){
// 	moduleViewCh, systemViewCh, matrixViewCh, smsCh, AlarmMySQLEntityCh :=protocol.ProtocolViewNodesHandler_YunHuan20201004(confnodech)
// 	return moduleViewCh, systemViewCh, matrixViewCh, smsCh, AlarmMySQLEntityCh
// }