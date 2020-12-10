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

	RecvCh chan byte
	SendCh chan byte
}

//无法给这里设计单例模式，因为就算把他与handler分离设计成独立的package
//引入这个包后，每当有新client连接就要执行这个方法、以及上面的NewPipelineTcpSocketConn
func (p *TcpConn)GenerateRecvCh(){
	p.RecvCh := make(chan []byte)
	go func (){

		defer func(){if p.Conn !=nil { p.Quit();        fmt.Println("有p.Conn意外断开了") }}() //主要是方式意外退出，quit函数会让p.Conn ==nil

		byteSpoon := make([]byte, 4096);        recvBuffer :=bytes.NewBuffer([]byte{})

		for {
			readlen, err := p.Conn.Read(byteSpoon) /*阀门*/
			tempBuf :=byteSpoon[:readlen]//如494f3031f10201,crc校验时需要截取有效字段
			//这里应该更新为写入日志
			if err != nil &&err !=io.EOF{ fmt.Println("Conn.Read err:",err); break }		
			if err ==io.EOF { continue }

			if readlen == 4{
				switch {
				case tempBuf[0] ==0x49&&tempBuf[1] ==0x4f&&tempBuf[2] ==0x30&&tempBuf[3] ==0x31:
					fmt.Println("接收到了IO1握手成功时发来的内容：",tempBuf);        continue
				default :
					fmt.Println("接收到错误的字节数组：",tempBuf);        continue
				}
			}
			
			if readlen<4{ fmt.Println("接收到错误的字节数组：",tempBuf);        continue }


			/*核心，为原始字节数组依次添加了ip,时间,tag*/
			ip :=[]byte(p.Conn.RemoteAddr().String());        tag :=[]byte("tcpsocket")
			timeStamp :=make([]byte,8);        binary.BigEndian.PutUint64(timeStamp, uint64(time.Now().UnixNano()))			

			recvBuffer.Reset();        _,_ = buffer.Write(ip);        _,_ = buffer.Write([]byte(" /-/ "));        _,_ = buffer.Write(timeStamp);
			 _,_ = recvBuffer.Write([]byte(" /-/ "));       _,_ = recvBuffer.Write(tag);        _,_ = recvBuffer.Write([]byte(" /-/ "));         _,_ = recvBuffer.Write(tempBuf)  

			if !p.NeedCRC{ p.RecvCh <-recvBuffer.Bytes();        continue }

			if ok := utils.CRCCheck(tempBuf[4:],utils.ISLITTLEENDDIAN);ok{
				p.RecvCh <-recvBuffer.Bytes()
			}else{
				fmt.Println("tcp北向通信时crc校验失败:",tempBuf)
			}
			/*---------------------------------------------------------*/


		}//for over
	}()//go func() over
}


func (p *TcpConn)GenerateSendCh(){
	p.SendCh := make(chan []byte)
	go func (){
		for b := p.SendCh{
			sendBytes(b)
		}
	}()
}

func (p *TcpConn)sendBytes(b []byte) {
	_, err :=p.Conn.Write(b)
	if err !=nil{ fmt.Println("write err:",err) }
}

func (p *TcpConn)InitActiveEventSender(modbus [][]byte, step int){
	l :=len(modbus)
	fmt.Println("in InitOwnActiveEventSender,len:",l)
	go func(){
		for i :=0; i<=l; i++{
			select{
			case <-p.QuitActiveEventSender:
				goto CLEANUP
			default:
			}
			if i==l{ i=0 }
			p.SendCh(modbus[i])
			time.Sleep(time.Duration(step)*time.Millisecond)
		}

		CLEANUP:
		// 不存在别的操作，因为for循环的不是一个管道，所以只要实现跳出了select就好
	}()
}

func (p *TcpConn)quitActiveEventSender(){
	go func(){
		p.QuitActiveEventSender<-true
		time.Sleep(time.Second)
		close(p.QuitActiveEventSender)
	}()
}

func (p *TcpConn)Quit(){
	defer func(){ if p.Conn !=nil { p.Conn.Close() }()

	p.quitActiveEventSender()
	close(p.SendCh)
	close(p.RecvCh)
}
