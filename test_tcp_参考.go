//package main

import(
	"github.com/ziyouzy/mylib/utils"
	"github.com/ziyouzy/mylib/yunhuanfactory/physicalnode"
	"net"
	"fmt"
)

func main(){
	testBytesCh :=make(chan []byte)
	testTcpHandler :=tcpHandler{}
	testTcpHandler.BindChanToListenAndRead(testCh)

	//之后就要从这个testCh里取出数据了
}


type tcpHandler struct{
	//tcp与udp共用的数据流入管道，需要在整体程序初始化前绑定
	//也就是说，初始化的函数要做两件事，并且先后顺序不能错:
	//1.绑定管道
	//2.listen客户端，并进行之后的handle操作
	RecvCh chan([]byte)

	//ip地址和结构体指针的哈系表
	ConnMap map[string]*tcpClient
}

func (p *tcpHandler)bindRecvCh(ch chan([]byte)){
	p.RecvCh =ch
}

/*先绑定管道后监听*/
func (p *tcpHandler)BindChanToListenAndRead(ch chan([]byte)){
	//初始化map
	p.ConnMap =make(map[string]*tcpClient)
	//绑定
	p.bindRecvCh(ch)

	//监听
	tcpAddr, err := net.ResolveTCPAddr("tcp", ":6668")
	if err != nil {
		fmt.Println(err.Error())
	}
	listener, err := net.ListenTCP("tcp", tcpAddr)
	if err != nil {
		fmt.Println(err.Error())
	}	
	defer listener.Close()
	fmt.Println("tcpHandler now is listening")

	//开始接收数据
	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println(err.Error())
		}
		ip,tcpconn :=newTcpClient(conn, p.RecvCh)
		p.ConnMap[ip] =tcpconn
		p.ConnMap[ip].Start("RECEIVE&SEND")
		fmt.Println("收到客户端的连接，客户端ip为：",ip)
		fmt.Println("用打印ip地址的方式测试ConnMap是否正常(不带*):",p.ConnMap[ip].Conn.RemoteAddr())
		fmt.Println("用打印ip地址的方式测试ConnMap是否正常(带*):",(*p.ConnMap[ip]).Conn.RemoteAddr())
	}
}


type tcpClient struct{
	RecvCh  chan([]byte)//tcp与udp共用的数据流入管道
	SendCh chan([]byte)//当前conn的写入管道，用于实现开关门之类的功能
	Conn net.Conn//当前连接
}

func newTcpClient(con net.Conn,ch chan([]byte))(string,*tcpClient){
	return con.RemoteAddr().String() , &tcpClient{
		Conn :con,
		RecvCh :ch,
	}
}

func (p *tcpClient)Start(mode string){
	switch (mode){
	case "RECEIVEONLY":
		go p.readThread()
	case "SENDONLY":
		go p.sendThread()
	case "RECEIVE&SEND":
		go p.readThread()
		go p.sendThread()
	default:
		fmt.Println("mode undefined")
	}
}

func (p *tcpClient)readThread() {
	//除非循环跳出否则函数不会回收
	defer p.Conn.Close()

	buf := make([]byte, 4096)
	for {
		readlen, err := p.Conn.Read(buf)
		tempBuf :=buf[:readlen]
		fmt.Println("readlen:",readlen,"tempbuf:",tempBuf)

		if err != nil {
			//这里应该更新为写入日志
			return
		}

		//crc校验
		if ok := utils.CRCCheck(tempBuf,true);ok{
			fmt.Println("是否已经被拆了呢?",tempBuf)
			p.RecvCh <-tempBuf
		}else{
			fmt.Println("crc校验失败:",tempBuf)
		}
	}
}

func (p *tcpClient)sendThread() {
	//除非管道关闭否则该函数不会回收
	defer p.Conn.Close()
	for tempbytes := range p.SendCh{
		p.Conn.Write(tempbytes)
	}
}