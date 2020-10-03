//流水线模式的tcp socket底层
package tcp

import(
	"net"
	"fmt"
	"time"
	"bytes"
	//"encoding/binary"

	"github.com/ziyouzy/mylib/utils"
)

const (
	TIMEFORMAT = "200601021504.05.000000000"
	NORMALTIMEFORMAT1 = "2006-01-02 15:04:05"
	NORMALTIMEFORMAT2 ="20060102150405"
)

type PipelineTcpSocketConn struct{
	Conn net.Conn//当前连接
	NeedCRC bool
}


func NewPipelineTcpSocketConn(c net.Conn, needcrc bool)(string, *PipelineTcpSocketConn){
	return c.RemoteAddr().String() , &PipelineTcpSocketConn{
		Conn :c,
		NeedCRC :needcrc,
	}
}

//无法给这里设计单例模式，因为就算把他与handler分离设计成独立的package
//引入这个包后，每当有新client连接就要执行这个方法、以及上面的NewPipelineTcpSocketConn
func (p *PipelineTcpSocketConn)GenerateRecvCh() chan([]byte){
	ch := make(chan []byte)
	go func (){
		defer p.Conn.Close()
		defer close(ch)

		ip :=[]byte(p.Conn.RemoteAddr().String())
		tag :=[]byte("tcpsocket")

		buf := make([]byte, 4096)
		for {
			now :=time.Now()
			p.Conn.SetReadDeadline(now.Add(time.Second * 5))
			readlen, err := p.Conn.Read(buf) /*阀门*/
			if err != nil {
				//这里应该更新为写入日志
				//fmt.Println("timeout?",err)
				return
			}			
			tempBuf :=buf[:readlen]//如494f3031f10201,crc校验时需要截取有效字段
			//fmt.Println(fmt.Sprintf("Recvbytes from %v : %v",p.Conn.RemoteAddr(),tempBuf))
			if len(tempBuf)<=4{
				fmt.Println("接收到错误的字节数组：",tempBuf)
			}else if !p.NeedCRC{
				presentTime :=time.Now().Format(TIMEFORMAT)
				ch <-bytes.Join([][]byte{ip,[]byte(presentTime),tag,tempBuf},[]byte(" "))
			}else{
				//var presentTime []byte
				//crc校验
				if ok := utils.CRCCheck(tempBuf[4:],utils.ISLITTLEENDDIAN);ok{
					presentTime :=time.Now().Format(TIMEFORMAT)
					ch <-bytes.Join([][]byte{ip,[]byte(presentTime),tag,tempBuf},[]byte(" "))
				}else{
					fmt.Println("tcp北向通信时crc校验失败:",tempBuf[4:])
				}
			}
		}
	}()//go func()
	return ch
}

func (p *PipelineTcpSocketConn)SendBytes(b []byte) {
	//fmt.Println(fmt.Sprintf("SendBytes to %v: %v",p.Conn.RemoteAddr(),b))
	p.Conn.Write(b)
}