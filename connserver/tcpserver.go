package connserver

import(
	"fmt"
	"net"

	"github.com/ziyouzy/mylib/connserver/con"
	"github.com/ziyouzy/mylib/conf"
)

func (p *ConnServer)TcpRecvCh(port string){
	go func(){
		tcpAddr, err := net.ResolveTCPAddr("tcp", port)
		if err != nil {fmt.Println("tcp第一次握手错误:",err.Error());return}
	
		listener, err := net.ListenTCP("tcp", tcpAddr)
		if err != nil {fmt.Println("tcp第二次握手错误:",err.Error());return}	
	
		defer listener.Close()


		//开始接收数据
		fmt.Println("tcp 的服务器端已开始监听")
		for {
			/*在这里就会阻塞，或者说目前这里智能监听到一种类型的连接，也就是tcp*/
			c, err := listener.Accept()
			if err != nil { fmt.Println("tcp第三次握手错误:",err.Error()); continue }
		
			fmt.Println("tcp第三次握手成功高，开始收容")
			key,client, timeout :=con.NewTcpCon(c,con.NEEDCRC)
			switch key{
			case "TCPCONN:192.168.10.2":
				client.InitActiveEventSender(conf.TcpModbus1)
			}
			p.ConnClientMap[key] =client
			clientrecvch :=client.GenerateRecvCh()
			if clientrecvch ==nil { fmt.Println("TCP-创建数据管道失败"); continue }
			p.fanInRecvCh(clientrecvch,key,timeout)

			fmt.Println("p.ConnClientMap updated:",p.ConnClientMap)
		}
	}()
}

func (p *ConnServer)fanInRecvCh(clientrecvch chan []byte, key string, timeout int){
	go func(){
		defer close(clientrecvch)
		defer delete(p.ConnClientMap,key)
		defer fmt.Println("该设备", timeout, "秒无应答，连接将会从ConnClientMap中删除：",key)

		for b := range clientrecvch{
			p.ServerRecvCh<-b
		}
	}()
}


