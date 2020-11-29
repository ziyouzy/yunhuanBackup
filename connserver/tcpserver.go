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
			if err != nil {fmt.Println("tcp第三次握手错误:",err.Error());return}
		
			fmt.Println("tcp第三次握手成功高，开始收容")
			key,client, timeout :=con.NewTcpCon(c,con.NEEDCRC)
			switch key{
			case "TCPCONN:192.168.10.2":
				client.InitOwnActiveEventSender(conf.TcpModbus1)
			}
			p.ConnClientMap[key] =client
			clientrecvch :=client.GenerateRecvCh()
			if clientrecvch ==nil{fmt.Println("TCP-创建数据管道失败");return}

			go func(){
				defer delete(p.ConnClientMap,key)
				defer fmt.Println("该设备", timeout, "秒无应答，连接将会从ConnClientMap中删除：",key)
				for b := range clientrecvch{
					fmt.Println("rawch_a")
					p.ServerRecvCh<-b
					fmt.Println("rawch_b")
				} 
			}()

			fmt.Println("p.ConnClientMap updated:",p.ConnClientMap)
		}
	}()
}

// func RecvCh()chan []byte {return cs.RecvCh()}
// func (p *ConnServer)RecvCh()chan []byte{
// 	return p.ServerRecvCh
// }



// func Test(){cs.Test()}
// func (p *ConnServer)Test(){
// 	ch :=make(chan []byte)
// 	go func(){
// 		for {
// 			fmt.Println("in test send b1")
// 			b :=[]byte{0xf1,0x02,0x00,0x20,0x00,0x08,0x6c,0xf6,}
// 			ch<-b
// 			//fmt.Println("exsit?",p.ConnClientMap["TCPCONN:192.168.10.2"])
// 			//p.ConnClientMap["TCPCONN:192.168.10.2"].SendBytes(b)
// 			time.Sleep(1*time.Second)
// 			fmt.Println("in test send b2")
// 			b = []byte{0xf1,0x01,0x00,0x00,0x00,0x08,0x29,0x3c,}
// 			ch<-b
// 			//p.ConnClientMap["TCPCONN:192.168.10.2"].SendBytes(b)
// 			time.Sleep(1*time.Second)
// 		}
// 	}()
	
// 	go func(){
// 		for {
// 			select{
// 			case b :=<-ch:
// 				if p.ConnClientMap["TCPCONN:192.168.10.2"] !=nil{
// 					p.ConnClientMap["TCPCONN:192.168.10.2"].SendBytes(b)
// 					fmt.Println("p.ConnClientMap len",len(p.ConnClientMap))
// 					fmt.Println("exsit?",p.ConnClientMap)
// 				}
// 			}
// 		}
// 	}()
//}