//serial是无状态连接所以不该设计wg.Done一类的功能
//他需要在程序运行开始的时候就持久话配置文档里所提到的各个serial端口，以及对应的各个端口管道
//当然了，管理他们的hander也是持久话到程序结束的
package connserver

// import (
// 	//"fmt"
// 	//"net"
// 	//"sync"
// 	//"strings"
	
// 	//"github.com/ziyouzy/mylib/serial"
// 	"github.com/ziyouzy/mylib/serial"
// )




// type pipelineSerialHandler struct{
// 	ConnMap map[string]*serial.PipelineSerialConn

// 	ProtocolPortsName []string//初始化时从配置文档获取
// 	ProtocolPortsBaud []int//初始化时从配置文档获取
// 	ProtocolPortsReadTimeout []int//初始化时从配置文档获取
// 	ProtocolPortsNeedCRC []bool//初始化时从配置文档获取
// }


// func (p *pipelineSerialHandler)ListenAndGenerateRecvCh()chan([]byte){	
// 	p.ConnMap =make(map[string]*serial.PipelineSerialConn)
// 	ch := make(chan []byte)//属于主函数的管道，不需要额外设计close事件，而是与程序自身一起开启与关闭

// 	//配置文档里提到了多少个端口，就监听多少个
// 	//配置文档里说某个端口叫什么名字，那就叫他什么名字
// 	//一切都在初始化时就定型了，这和tcp是最大的不同
// 	for key, portname := range p.ProtocolPortsName{
// 		portname,serialconn :=serial.NewPipelineSerialConn(portname,p.ProtocolPortsBaud[key],
// 			p.ProtocolPortsReadTimeout[key],p.ProtocolPortsNeedCRC[key])

// 		p.ConnMap[portname] =serialconn

// 		go func(oneClientCh chan []byte){
// 			for b := range oneClientCh{/*阀门*/
// 				ch <-b
// 			}
// 		}((*serialconn).GenerateRecvCh(portname))
// 	}

// 	return ch
// }