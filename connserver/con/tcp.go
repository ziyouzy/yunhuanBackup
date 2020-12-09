package con

import(
	"net"
	//"net/url"
	"fmt"
	"time"
	"bytes"
	"io"
	"encoding/binary"

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
		buffer :=bytes.NewBuffer([]byte{});
		for {
			readlen, err := p.Conn.Read(buf) /*阀门*/
			tempBuf :=buf[:readlen]//如494f3031f10201,crc校验时需要截取有效字段
			//这里应该更新为写入日志
			if err != nil &&err !=io.EOF{fmt.Println("Conn.Read err:",err); break}		
			if err ==io.EOF{continue}

			if readlen==4{
				switch {
				case tempBuf[0] ==0x49&&tempBuf[1] ==0x4f&&tempBuf[2] ==0x30&&tempBuf[3] ==0x31:
					fmt.Println("接收到了IO1握手成功时发来的内容：",tempBuf);        continue
				default :
					fmt.Println("接收到错误的字节数组：",tempBuf);        continue
				}
			}
			
			if readlen<4{ fmt.Println("接收到错误的字节数组：",tempBuf);        continue }

			//核心，为原始字节数组依次添加了ip,时间,tag
			ip :=[]byte(p.Conn.RemoteAddr().String());        tag :=[]byte("tcpsocket");        /*l :=len(tempBuf)*/
			timeStamp :=make([]byte,8);        binary.BigEndian.PutUint64(timeStamp, uint64(time.Now().UnixNano()))			

			buffer.Reset();        _,_ = buffer.Write(ip);        _,_ = buffer.Write([]byte(" /-/ "));        _,_ = buffer.Write(timeStamp);
			 _,_ = buffer.Write([]byte(" /-/ "));       _,_ = buffer.Write(tag);        _,_ = buffer.Write([]byte(" /-/ "));         _,_ = buffer.Write(tempBuf)  

			//if !p.NeedCRC{ ch <-bytes.Join([][]byte{ip,timeStamp,tag,tempBuf},[]byte(" "));        continue }
			if !p.NeedCRC{ ch <-buffer.Bytes();        continue }
			
			if ok := utils.CRCCheck(tempBuf[4:],utils.ISLITTLEENDDIAN);ok{
					ch <-buffer.Bytes()
			}else{
				fmt.Println("tcp北向通信时crc校验失败:",tempBuf[4:])
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
			time.Sleep(100*time.Millisecond)
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
