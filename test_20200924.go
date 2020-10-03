package main

import(
	"github.com/ziyouzy/mylib/protocol"
	"github.com/ziyouzy/mylib/physicalnode"
	//"github.com/ziyouzy/mylib/tcp"--->tcp的handler暂时设计成属于当前main包
	"fmt"
)

func main(){
	//初始化底层连接 start
	tcpHandler :=pipelineTcpHandler{}
	tcpBytesCh := tcpHandler.ListenAndGenerateRecvCh()

	portnameArr, baudArr, readtimeoutArr,needcrcArr := GetSerialPortsByProtocol_YunHuan20200924()
	serialHandler :=pipelineSerialHandler{
		ProtocolPortsName : portnameArr,
		ProtocolPortsBaud : baudArr,
		ProtocolPortsReadTimeout : readtimeoutArr,
		ProtocolPortsNeedCRC : needcrcArr,
	}
	serialBytesCh := serialHandler.ListenAndGenerateRecvCh()
	//初始化底层连接 end
	


	//protocol part start
	physicalNodeCh :=physicalNodeChByProtocol(tcpBytesCh)
	go sendByProtocol(&tcpHandler)
	//protocol part end

	//扇出短信机管道和uinode管道 start
	//短信机需要去结合配置文件，配置文件是否属于协议的一部分？不是！！！
	//协议和配置是彼此独立平行，但是用法一致的东西！！！
	//这个思路不错，也就是说，会存在独立的setting/conf包，设计套路和protocol包一样
	for phy := range physicalNodeCh{
		fmt.Println(phy.GetHandler())
		//hander的格式是:494f3031f10101(tcpsokcet)
		//可以为physicalNode接口设计一个多返回值方法(如forViewMatrix，或forViewNodeAssertor)，从而判断出其可以生成哪种uinode
	}
	


	//之后就要从这个testCh里取出数据了
	for{

	}
}

func physicalNodeChByProtocol(in chan []byte)chan physicalnode.PhysicalNode{
	out :=make(chan physicalnode.PhysicalNode)
	go func(){
		defer close(out)
		for b := range in{
			physicalNodeEntity :=protocol.RecvProtocol_YunHuan20200924(b)
			physicalNodeEntity.FullOf()
			out<-physicalNodeEntity/*阀门*/
			//fmt.Println("Handler:",physicalNodeEntity)
		}
	}()
	return out
}

//南向，river的上游向，存在固定的发送频率
func sendBySorthProtocol(tcpHandler *pipelineTcpHandler) {	
	pMap :=protocol.SendHandlerBySorthProtocol_YunHuan20200924()
	//pMap的key为ip地址，而ip地址是协议的一部分，程序运行起来后，不可发生增减
	//ticket也是在protocol.SendHandlerBySorthProtocol_YunHuan20200924()里实现的
	for mark,ch := range pMap{
		//range的ch也是报文的一部分，包含了不同ip地址，每个ip地址都对应了不同的报文组
		for b :=range ch {
			tempbytes :=b
			switch{
			case tcpHandler.ConnMap[mark]:
				tcpHandler.ConnMap[ip].SendBytes(tempbytes)
			// case serialHandler.ConnMap[mark]:
			// 	serialHandler.ConnMap[ip].SendBytes(tempbytes)
			// case udpHandler.ConnMap[mark]:
			// 	udpHandler.ConnMap[ip].SendBytes(tempbytes)
			}
		}
	}
}