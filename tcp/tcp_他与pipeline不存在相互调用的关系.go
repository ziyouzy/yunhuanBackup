//非流水线模式
package tcp

import(
	"net"
	"fmt"

	"github.com/ziyouzy/mylib/utils"
	//"github.com/ziyouzy/mylib/yunhuanfactory/physicalnode"
)



type tcpClient struct{
	RecvCh  chan([]byte)//tcp与udp共用的数据流入管道
	SendCh chan([]byte)//当前conn的写入管道，用于实现开关门之类的功能
	Conn net.Conn//当前连接

	NeedCRC bool
}


func NewTcpClient(con net.Conn, r chan([]byte), s chan([]byte), n bool)(string,*tcpClient){
	return con.RemoteAddr().String() , &tcpClient{
		Conn :con,

		RecvCh :r,
		SendCh :s,

		NeedCRC :n,
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
			fmt.Println("crc校验成功，但是是否已经被拆了呢?",tempBuf)
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