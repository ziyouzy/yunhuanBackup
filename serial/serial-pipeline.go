package serial

import (
    //"errors"
    "time"
    "fmt"
    "bytes"

    "github.com/tarm/serial"
    
    "github.com/ziyouzy/mylib/utils"
	
)

const (
	TIMEFORMAT = "200601021504.05.000000000"
	NORMALTIMEFORMAT1 = "2006-01-02 15:04:05"
	NORMALTIMEFORMAT2 ="20060102150405"
)

type PipelineSerialConn struct {
    Conn *serial.Port // 串口客户端
    NeedCRC bool
}

func NewPipelineSerialConn(portname string, portbaud int, readtimeout int, needcrc bool)(string, *PipelineSerialConn){
    c := serial.Config{Name: portname, Baud: portbaud, ReadTimeout: time.Duration(readtimeout)}
    con, err := serial.OpenPort(&c)//打开串口
    if (err !=nil){ 
        return fmt.Sprintf("创建串口连接时出现问题，串口号:%s",portname),nil
    }else{
	    return portname , &PipelineSerialConn{
		    Conn :con,
		    NeedCRC :needcrc,
        }
    }
}

//基于数据流动思想的流水线模式，管道生成的方法也是该管道阀门的所在位置
//如下的"time.Sleep(500*time.MilliSecond)"既是阀门
func (p *PipelineSerialConn)GenerateRecvCh(portname string) chan([]byte){
    ch := make(chan []byte)
    tempBuf := make([]byte, 4096)
    portName :=[]byte(portname)
    tag :=[]byte("serial")
	go func (){
        for{    
            time.Sleep(time.Duration(500)*time.Millisecond)  /*阀门*/	
            len, err := p.Conn.Read(tempBuf)
            if err !=nil{
                fmt.Println("init serial falid err:",err)
                //无论哪种失败都不能因此而关闭连接和对应的管道，只提示，这是串口与tcp设计逻辑上的最大不同点之一
                //等同于永远都不该关闭tcp的listen一样
            }
            
            oneLine :=tempBuf[:len]
            if !p.NeedCRC{
				presentTime :=time.Now().Format(TIMEFORMAT)
				ch <-bytes.Join([][]byte{portName,[]byte(presentTime),tag,oneLine},[]byte(" "))
			}else{

				if ok := utils.CRCCheck(oneLine,utils.ISLITTLEENDDIAN);ok{
					presentTime :=time.Now().Format(TIMEFORMAT)
					ch <-bytes.Join([][]byte{portName,[]byte(presentTime),tag,oneLine},[]byte(" "))
				}else{
					fmt.Println("serial北向通信时crc校验失败:",oneLine)
                }

            }//NeedCRC end
        }//for end
    }()//go  func()
    return ch
}          
        
func (p *PipelineSerialConn)SendBytes(b []byte) {
	p.Conn.Write(b)
}

