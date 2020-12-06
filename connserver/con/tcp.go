package con

import(
	"net"
	//"net/url"
	"fmt"
	"time"
	"bytes"
	"io"

	"github.com/ziyouzy/mylib/utils"
)

type TcpConn struct{
	Conn net.Conn//当前连接
	NeedCRC bool
	QuitActiveEventSender chan bool
}

//无法给这里设计单例模式，因为就算把他与handler分离设计成独立的package
//引入这个包后，每当有新client连接就要执行这个方法、以及上面的NewPipelineTcpSocketConn
func (p *TcpConn)GenerateRecvCh() chan []byte{
	ch := make(chan []byte)
	go func (){
		defer close(ch)
		defer func(){if p.Conn !=nil { p.Conn.Close() }}()

		buf := make([]byte, 4096)
		for {
			readlen, err := p.Conn.Read(buf) /*阀门*/
			//这里应该更新为写入日志
			if err != nil &&err !=io.EOF{fmt.Println("Conn.Read err:",err); break}		
			if err ==io.EOF{continue}
					
			tempBuf :=buf[:readlen]//如494f3031f10201,crc校验时需要截取有效字段
			//fmt.Println(url.PathUnescape(string(tempBuf)))
			ip :=p.Conn.RemoteAddr().String();        tag :="tcpsocket";        l :=len(tempBuf)
			if l==4{
				switch {
				case tempBuf[0] ==0x49&&tempBuf[1] ==0x4f&&tempBuf[2] ==0x30&&tempBuf[3] ==0x31:
					fmt.Println("接收到了IO1握手成功时发来的内容：",tempBuf)
				}
			}else if l<4{ 
				fmt.Println("接收到错误的字节数组：",tempBuf);    continue
			}else if !p.NeedCRC{
				//fmt.Println("??????!:",tempBuf)
				presentTime :=time.Now().UnixNano()
				//核心，为原始字节数组依次添加了ip,时间,tag
				//ch <-bytes.Join([][]byte{[]byte(ip),[]byte(presentTime),[]byte(tag),tempBuf},[]byte(" "))
				ch <-bytes.Join([][]byte{[]byte(ip),,[]byte(tag),tempBuf},[]byte(" "))
			}else{
				//crc校验
				//fmt.Println("??????:",tempBuf)
				if ok := utils.CRCCheck(tempBuf[4:],utils.ISLITTLEENDDIAN);ok{
					//presentTime :=time.Now().Format(TIMEFORMAT)
					timeBuf := make([]byte, 8)
					binary.BigEndian.PutUint64(timeBuf, time.Now().Format(TIMEFORMAT))
					//核心，为原始字节数组依次添加了ip,时间,tag
					ch <-bytes.Join([][]byte{[]byte(ip),timeBuf,[]byte(tag),tempBuf},[]byte(" "))
				}else{
					fmt.Println("tcp北向通信时crc校验失败:",tempBuf[4:])
				}
			}
		}
	}()//go func()
	return ch
}

func (p *TcpConn)InitActiveEventSender(modbus [][]byte){
	l :=len(modbus)
	fmt.Println("in InitOwnActiveEventSender,len:",l)
	go func(){
		for i :=0;i<=l;i++{
			select{
			case <-p.QuitActiveEventSender:
				break
			default:
			}
			if i==l{i=0}
			p.SendBytes(modbus[i])
			time.Sleep(1000*time.Millisecond)
		}
	}()
}

func (p *TcpConn)SendBytes(b []byte) {
	_,err :=p.Conn.Write(b)
	if err !=nil{ fmt.Println("write err:",err) }
}

func (p *TcpConn)Quit(){
	p.Conn.Close()
	p.QuitActiveEventSender<-true
}
