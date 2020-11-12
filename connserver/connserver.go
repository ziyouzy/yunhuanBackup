//服务端的listen操作需要持久话到程序结束，同时所生成的总管道也是如此
package connserver

import (
	"fmt"
	"net"
	//"sync"
	"strings"
	
	"github.com/ziyouzy/mylib/tcp"
)


var cs *ConnServer
//tcphandler模块的key为ip地址(无端口号)，可与sendmsg对象无缝对接
//后期设计sendcontroller时可以在对象结构体内部将二者组合
//但是需要先完成将tcphandler重构成connhandler这一步
//或者说，ticketsender和connhander做组合，同时door和connhander也作组合
//connhander是个底层，等同于nodedo是alarmcontroller和nodedocontroller的底层一样

//他似乎也需要存在一个全局实体
type ConnServer struct{
	ConnClientMap map[string]conn.ConnClient 
	ServerRecvCh chan []byte //不需要额外设计close事件，而是与程序自身一起开启与关闭
}


func ListenAndGenerateAllRecvCh(){cs.ListenAndGenerateAllRecvCh()}
func (p *ConnServer)ListenAndGenerateAllRecvCh(){
	p.ConnClientMap =make(map[string]conn.ConnClient)
	p.generateAndCollectTcpRecvCh(":6668")
	//p.generateAndCollectUdpRecvCh(":6669")
}


func RecvCh()chan []byte {return cs.ServerRecvCh}
func (p *ConnServer)generateAndCollectTcpRecvCh(port string){
	if tcpAddr, err := net.ResolveTCPAddr("tcp", port);err != nil {
		fmt.Println(err.Error())
	}

	if listener, err := net.ListenTCP("tcp", tcpAddr);err != nil {
		fmt.Println(err.Error())
	}	

	defer listener.Close()

	go func(){
	//开始接收数据
		fmt.Println("tcp 的服务器端已开始监听")
		for {
			/*在这里就会阻塞，或者说目前这里智能监听到一种类型的连接，也就是tcp*/
			if con, err := listener.Accept();err != nil {
				fmt.Println(err.Error())
			}
		
			key,client,timeoutsec :=conn.NewConnClient(con,conn.NEEDCRC)

			go func(key string, clientrecvch chan []byte,timeout int){
				defer delete(p.ConnClientMap,key)
				defer fmt.Println("该设备", timeout, "秒无应答，连接将会从ConnClientMap中删除：",key)
				for b := range clientrecvch{
					p.ServerRecvCh<-b
				} 
			}(key, client.GenerateRecvCh(),timeoutsec)

			p.ConnClientMap[key] =client
			fmt.Println("p.ConnClientMap updated:",p.ConnClientMap)
		}
	}()
}

