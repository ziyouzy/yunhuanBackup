package main

import(
	"fmt"
	//"encoding/json"

	"github.com/ziyouzy/mylib/protocol"
	"github.com/ziyouzy/mylib/conf"
)

func main(){
	conf.InitConfMap()
	tcphandler :=pipelineTcpHandler{}
	// serialhandler := pipelineSerialHandler{ProtocolPortsName:[]string{"serial1","serial2","serial3"},
	// 	ProtocolPortsBaud:[]int{9600,9600,9600},ProtocolPortsReadTimeout:[]int{5,5,5},ProtocolPortsNeedCRC:[]bool{true,true,true},}
//--
	tcpCh :=tcphandler.ListenAndGenerateRecvCh()
	//serialCh :=serialhandler.ListenAndGenerateRecvCh()
//--
	//sendmap目前会返回两个字段，一个是tcp的管道，一个是udp的管道
	sendmap := protocol.ProtocolPrepareSendTicketMgr_YunHuan20200924() 
	for k,msgch :=range sendmap{
		o :=k//完美的解决了问题，其实就是之前经常在百度上看到的关于那个range的坑的各类文章，区别在于，不只是v存在这个问题，k也同样如此
		go func(){
			//问题其实是因为range神坑与单独开携程使用了这个k共同导致的，或者说，毕竟range的时候k和v都只有一个副本，你开携程相当于把这一个副本作为参数传递进了这个携程，无论如何，肯定只会是一个值
			for msg := range msgch{
				if(tcphandler.ConnMap[o] !=nil){
					tcphandler.ConnMap[o].SendBytes(msg)
				}else{
					//fmt.Println("k is nil:",k)
				}
			}
		}()
	}
//--
//--
	rawCh :=MergeConCS(tcpCh/*,serialCh*/)
	confNodeCh :=Convert(rawCh)
	moduleViewCh, systemViewCh, matrixViewCh, smsCh := Separate(confNodeCh)
//--	
for{
	select {
	case temp :=<-moduleViewCh:
		fmt.Println("moduleView:",string(temp),"\n")
	case temp :=<-systemViewCh:
		fmt.Println("systemView:",string(temp),"\n")
	case temp :=<-matrixViewCh:
		fmt.Println("matrixView:",string(temp),"\n")
	case temp :=<-smsCh:
		fmt.Println("sms:",string(temp))
	}
}
	

	//separate&merge：分离&融合
	// moduleViewCh, systemViewCh, matrixViewCh, smsCh := Separate(confNodeCh)
	
	// for{
	// 	select {
	// 	case temp :=<-moduleViewCh:
	// 		fmt.Println("moduleView:",string(temp))
	// 	case temp :=<-systemViewCh:
	// 		fmt.Println("systemView:",string(temp))
	// 	case temp :=<-matrixViewCh:
	// 		fmt.Println("matrixView:",string(temp))
	// 	case temp :=<-smsCh:
	// 		fmt.Println("sms:",string(temp))
	// 	}

	// default:
	// 	fmt.Println("default")
	//}
}

func MergeConCS(cs ...chan []byte)chan []byte{
	out :=make(chan []byte)

	collect := func(in chan []byte){
		for b :=range in{
			out <-b
		}
	}

	for _, c := range cs {
		go collect(c)
	}

	return out
}

func Convert(ch chan []byte) chan conf.ConfNode{
	confnodech := make(chan conf.ConfNode)
	go func(){
		for b := range ch{
			PhysicalNode :=protocol.ProtocolPreparePhysicalNode_YunHuan20200924(b)
			//fmt.Println("PhysicalNode:",PhysicalNode)
			ConfNodeArr :=conf.NewConfNodeArr(PhysicalNode)
			for _, confnode := range ConfNodeArr{
				confnodech<-confnode
			}
		}
	}()
	return confnodech
}

func Separate(confnodech chan conf.ConfNode)(chan []byte, chan []byte, chan []byte, chan []byte){
	moduleViewCh, systemViewCh, matrixViewCh, smsCh :=protocol.ProtocolViewNodesHandler_YunHuan20201004(confnodech)
	return moduleViewCh, systemViewCh, matrixViewCh, smsCh
}