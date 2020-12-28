package connserver

import(
	"fmt"
	"net"
	"strings"

	"github.com/ziyouzy/mylib/connserver/con"
)

func (p *ConnServer)TcpListenAndCollect(port string){

	var build strings.Builder
	build.WriteString(":")
	build.WriteString(port)
	port =build.String() 

	go func(){
		tcpAddr, err := net.ResolveTCPAddr("tcp", ":6668");        if err != nil { fmt.Println("tcp第一次握手错误:",err.Error());        return }
	
		listener, err := net.ListenTCP("tcp", tcpAddr);        if err != nil { fmt.Println("tcp第二次握手错误:",err.Error());        return }	
	
		defer listener.Close()

		//开始接收数据
		fmt.Println("tcp 的服务器端已开始监听")
		for {
			/*在这里就会阻塞，或者说目前这里智能监听到一种类型的连接，也就是tcp*/
			c, err := listener.Accept()
			if err != nil { fmt.Println("tcp第三次握手错误:",err.Error());        continue }
		
			fmt.Println("tcp第三次握手成功高，开始收容")
			key,client, recvch, /*sendch*/_ :=con.NewTcpCon(c, con.NEEDCRC, 15, 2)
			if recvch ==nil { fmt.Println("TCP-创建数据管道失败");        continue }
			
			p.collectClientRecvCh(recvch,key)

			p.ConnClientMap[key] =client
			fmt.Println("有新的tcp连接并入，key为:", key)
		}
	}()
}


