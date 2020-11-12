//流水线模式的tcp socket底层
package conn

import(
	"net"
	"fmt"
	"time"
	"bytes"
	//"encoding/binary"

	"github.com/ziyouzy/mylib/utils"
)

type TcpConn struct{
	Conn net.Conn//当前连接
	NeedCRC bool
}

//无法给这里设计单例模式，因为就算把他与handler分离设计成独立的package
//引入这个包后，每当有新client连接就要执行这个方法、以及上面的NewPipelineTcpSocketConn
func (p *TcpConn)GenerateRecvCh() chan([]byte){
	ch := make(chan []byte)

	go func (){
		defer p.Conn.Close()
		defer close(ch)

		buf := make([]byte, 4096)
		for {
			now :=time.Now()
			readlen, err := p.Conn.Read(buf) /*阀门*/
			if err != nil {
				//这里应该更新为写入日志
				//fmt.Println("timeout?",err)
				return
			}			
			tempBuf :=buf[:readlen]//如494f3031f10201,crc校验时需要截取有效字段

			ip :=p.Conn.RemoteAddr().String()
			tag :="tcpsocket"
			if len(tempBuf)<=4{
				fmt.Println("接收到错误的字节数组：",tempBuf)
			}else if !p.NeedCRC{
				presentTime :=time.Now().Format(TIMEFORMAT)
				//核心，为原始字节数组依次添加了ip,时间,tag
				ch <-bytes.Join([][]byte{[]byte(ip),[]byte(presentTime),[]byte(tag),tempBuf},[]byte(" "))
			}else{
				//var presentTime []byte
				//crc校验
				if ok := utils.CRCCheck(tempBuf[4:],utils.ISLITTLEENDDIAN);ok{
					presentTime :=time.Now().Format(TIMEFORMAT)
					//核心，为原始字节数组依次添加了ip,时间,tag
					ch <-bytes.Join([][]byte{[]byte(ip),[]byte(presentTime),[]byte(tag),tempBuf},[]byte(" "))
				}else{
					fmt.Println("tcp北向通信时crc校验失败:",tempBuf[4:])
				}
			}
		}
	}()//go func()
	return ch
}

func (p *TcpConn)SendBytes(b []byte) {
	p.Conn.Write(b)
}