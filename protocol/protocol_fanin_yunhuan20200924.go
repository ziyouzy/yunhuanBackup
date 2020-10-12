//协议细节
//1.494f3031f10201是“开关量输入”（DI）,只有两根线，也是目前最常用的，0101是“开关量输出”(DO)，3根线
package protocol

import(
	"bytes"
	"time"
	//"fmt"


	"github.com/ziyouzy/mylib/physicalnode"
)

//生成核心物理节点
func ProtocolPreparePhysicalNode_YunHuan20200924(b []byte)physicalnode.PhysicalNode{
	//从这里开始就是协议的描述了
	//这里要先用tag区分tcp、udp、serial、snmp
	//以及ui端传来的开关门等请求(localqt,localweb,remoteqt,remoteweb)

	//ip,[]byte(presentTime),tag,tempBuf}
	bufarr :=bytes.Fields(b)//按照空白分割
	//mark :=string(bufarr[0])//ip或者portname
	//presenttime :=string(bufarr[1])
	/*mark和presenttime字段会在把buf传入工厂函数后，函数内部会再次拆分并合成*/
	
	tag :=string(bufarr[2])
	buf :=bufarr[3]
	//开始实现协议
	switch (tag){
	case "tcpsocket":
		switch {//s1
		case buf[0]==0x49&&buf[1]==0x4f&&buf[2]==0x30&&buf[3]==0x31:
			switch {//s2
			case buf[4]==0xf1:
				switch {//s3
				case buf[5]==0x01&&buf[6]==0x01:
					return physicalnode.NewPhysicalNodeFromBytes(b, tag, "YUNHUAN20200924","DI20200924")
				case buf[5]==0x02&&buf[6]==0x01:
					return physicalnode.NewPhysicalNodeFromBytes(b, tag, "YUNHUAN20200924","DO20200924")
				default:
					return nil
				}//s3
			default:
				return nil	
			}//s2
		default:
			return nil		
		}//s1

	case "serial":
		switch{
		case buf[0]==0xf1:
			switch {//s1
			case buf[1]==0x01&&buf[2]==0x01:
				//fmt.Println("f10101!")
				return physicalnode.NewPhysicalNodeFromBytes(b, tag, "YUNHUAN20200924","DO20200924")
			case buf[1]==0x02&&buf[2]==0x01:
				//fmt.Println("f10201!")
				return physicalnode.NewPhysicalNodeFromBytes(b, tag, "YUNHUAN20200924","DI20200924")
			default:
				return nil
			}//s1
		default:
			return nil
		}//s2
	//case "localqt":
		//反序列化buf为某种功能结构体(如physicalnode.NewDoorMgr)
		//return 这个结构体
	default:
		return nil	
	}//tag
}

//--------
//--------
//--------

//获取发送数据管理器
func ProtocolPrepareSendTicketMgr_YunHuan20200924() map[string]chan []byte{
	chs :=make(map[string]chan []byte)
	ticker := time.NewTicker(time.Duration(1)*time.Millisecond)
 	modbus := [][]byte{
		{0xf1,0x01,0x00,0x00,0x00,0x08,0x29,0x3c,},
		{0xf1,0x02,0x00,0x20,0x00,0x08,0x6c,0xf6,},
	}

	tcpch :=make(chan []byte)
	go func(){
		defer close(tcpch)
		for {
			for _,b :=range modbus{
				tcpch<-b
			}
        	/*presentTime*/ _ /*:*/= <-ticker.C/*阀门*/
        	//fmt.Println(/*presentTime.String(),":*/"发送协议已经完成一个周期")
		}
	}()

	udpch :=make(chan []byte)
	go func(){
		defer close(udpch)
		for {
			for _,b :=range modbus{
				udpch<-b
			}
        	/*presentTime*/ _ /*:*/= <-ticker.C/*阀门*/
        	//fmt.Println(/*presentTime.String(),":*/"发送协议已经完成一个周期")
		}
	}()

	serialch :=make(chan []byte)
	go func(){
		defer close(serialch)
		for {
			for _,b :=range modbus{
				serialch<-b
			}
        	/*presentTime*/ _ /*:*/= <-ticker.C/*阀门*/
        	//fmt.Println(/*presentTime.String(),":*/"发送协议已经完成一个周期")
		}
	}()

	chs["192.168.10.2"] =tcpch
	//chs["tty0"] =tcpch
	chs["192.168.11.2"]=udpch

	return chs
}

//获取串口在线表
func ProtocolPrepareSerialPorts_YunHuan20200924()([]string, []int, []int, []bool){
	return []string{"tty1","tty2","tty3",}, []int{9600,9600,2400,}, []int{500,500,500,}, []bool{true,true,true,}
}

func ProtocolPrepareSmsMgr_YunHuan20200924() []string{
	return []string{"tty1",}
}

func ProtocolPrepareDoorMgr_YunHuan20200924() map[string]map[string][][]byte{
	maps :=make(map[string]map[string][][]byte)	

	map1 :=make(map[string][][]byte)	
	map1["前门"]=[][]byte{[]byte{0xf1,0x02,0x04,0x04,0x01},[]byte{0xf1,0x02,0x04,0x04,0x02}}
	map1["后门"]=[][]byte{[]byte{0xf1,0x02,0x04,0x03,0x01},[]byte{0xf1,0x02,0x04,0x04,0x02}}

	map2 :=make(map[string][][]byte)	
	map2["前门1"]=[][]byte{[]byte{0xf1,0x02,0x04,0x04},[]byte{0xf1,0x02,0x04,0x04,0x02}}
	map2["后门1"]=[][]byte{[]byte{0xf1,0x02,0x04,0x03},[]byte{0xf1,0x02,0x04,0x04,0x02}}
	map2["前门2"]=[][]byte{[]byte{0xf1,0x02,0x04,0x04},[]byte{0xf1,0x02,0x04,0x04,0x02}}
	map2["后门2"]=[][]byte{[]byte{0xf1,0x02,0x04,0x03},[]byte{0xf1,0x02,0x04,0x04,0x02}}

	maps["192.168.10.254"]=map1
	maps["192.168.10.253"]=map2

	return maps
}

// func ProtocolPrepareCameraMgr_YunHuan20200924(){

// }
