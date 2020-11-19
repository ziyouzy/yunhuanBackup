//服务端的listen操作需要持久话到程序结束，同时所生成的总管道也是如此
package connserver

import (
	"fmt"
	"net"
	"time"
	//"sync"
	//"strings"
	
	"github.com/ziyouzy/mylib/connserver/connclient"
)


var cs *ConnServer
//tcphandler模块的key为ip地址(无端口号)，可与sendmsg对象无缝对接
//后期设计sendcontroller时可以在对象结构体内部将二者组合
//但是需要先完成将tcphandler重构成connhandler这一步
//或者说，ticketsender和connhander做组合，同时door和connhander也作组合
//connhander是个底层，等同于nodedo是alarmcontroller和nodedocontroller的底层一样

//他似乎也需要存在一个全局实体
type ConnServer struct{
	ConnClientMap map[string]connclient.ConnClient 
	ServerRecvCh chan []byte //不需要额外设计close事件，而是与程序自身一起开启与关闭
}

//返回的管道未设定消费者
func LoadSingletonPatternRecvCh()chan []byte{cs=new(ConnServer);return cs.RecvCh()}
func (p *ConnServer)RecvCh()chan []byte{
	p.ServerRecvCh =make(chan []byte)
	return p.ServerRecvCh
}

func LoadSingletonPatternListenAndCollect(){cs.ListenAndCollect()}
func (p *ConnServer)ListenAndCollect(){
	p.ConnClientMap =make(map[string]connclient.ConnClient)
	p.generateAndCollectTcpRecvCh(":6668")
	//p.generateAndCollectUdpRecvCh(":6669")
}
func (p *ConnServer)generateAndCollectTcpRecvCh(port string){
	go func(){
		tcpAddr, err := net.ResolveTCPAddr("tcp", port)
		if err != nil {fmt.Println("tcp第一次握手错误:",err.Error())}
	
		listener, err := net.ListenTCP("tcp", tcpAddr)
		if err != nil {fmt.Println("tcp第二次握手错误:",err.Error())}	
	
		defer listener.Close()


		//开始接收数据
		fmt.Println("tcp 的服务器端已开始监听")
		for {
			/*在这里就会阻塞，或者说目前这里智能监听到一种类型的连接，也就是tcp*/
			con, err := listener.Accept()
			if err != nil {fmt.Println("tcp第三次握手错误:",err.Error())}
		
			fmt.Println("tcp第三次握手成功高，开始收容")
			key,client, timeout :=connclient.NewConnClient(con,connclient.NEEDCRC)

			//clientrecvch的消费者子携程
			//同时子携程也是p.ServerRecvCh的生产者
			//p.ServerRecvCh需要在上层消费
			clientrecvch :=client.GenerateRecvCh()
			go func(){
				defer delete(p.ConnClientMap,key)
				defer fmt.Println("该设备", timeout, "秒无应答，连接将会从ConnClientMap中删除：",key)
				for b := range clientrecvch{
					//fmt.Println("bbbb:",b)
					p.ServerRecvCh<-b
				} 
			}()

			p.ConnClientMap[key] =client
			fmt.Println("p.ConnClientMap updated:",p.ConnClientMap)
		}
	}()
}

// func RecvCh()chan []byte {return cs.RecvCh()}
// func (p *ConnServer)RecvCh()chan []byte{
// 	return p.ServerRecvCh
// }

func ClientMap()map[string]connclient.ConnClient{return cs.ClientMap()}
func (p *ConnServer)ClientMap()map[string]connclient.ConnClient {
	return p.ConnClientMap
}

func Test(){cs.Test()}
func (p *ConnServer)Test(){
	ch :=make(chan []byte)
	go func(){
		for {
			fmt.Println("in test send b1")
			b :=[]byte{0xf1,0x02,0x00,0x20,0x00,0x08,0x6c,0xf6,}
			ch<-b
			//fmt.Println("exsit?",p.ConnClientMap["TCPCONN:192.168.10.2"])
			//p.ConnClientMap["TCPCONN:192.168.10.2"].SendBytes(b)
			time.Sleep(1*time.Second)
			fmt.Println("in test send b2")
			b = []byte{0xf1,0x01,0x00,0x00,0x00,0x08,0x29,0x3c,}
			ch<-b
			//p.ConnClientMap["TCPCONN:192.168.10.2"].SendBytes(b)
			time.Sleep(1*time.Second)
		}
	}()
	
	go func(){
		for {
			select{
			case b :=<-ch:
				if p.ConnClientMap["TCPCONN:192.168.10.2"] !=nil{
					p.ConnClientMap["TCPCONN:192.168.10.2"].SendBytes(b)
					fmt.Println("p.ConnClientMap len",len(p.ConnClientMap))
					fmt.Println("exsit?",p.ConnClientMap)
				}
			}
		}
	}()
}
// func SelectClients(keys []string)[]connclient.ConnClient {return cs.SelectClients(keys)}
// func (p *ConnServer)SelectClients(keys []string)[]connclient.ConnClient {
// 	clients :=make([]connclient.ConnClient,1,5)
// 	for _,v :=range keys{
// 		//fmt.Println("vvv:",v)
// 		if c :=p.ConnClientMap[v];c !=nil{
// 			clients =append(clients,c)
// 		}else{
// 			fmt.Println("ConnClientMap[v] is nil,map is ",p.ConnClientMap)
// 		}
// 	}
// 	return clients
// }