//协议细节
//1.494f3031f10201是“开关量输入”（DI）,只有两根线，也是目前最常用的，0101是“开关量输出”(DO)，3根线
package protocol

import(
	"bytes"
	"time"
	//"fmt"


	"github.com/ziyouzy/mylib/physicalnode"
)

//获取发送数据管理器
func ProtocolPrepareSendTicketMgr_YunHuan20200924() map[string]chan []byte{
	chs :=make(map[string]chan []byte)
	ticker := time.NewTicker(time.Duration(1900)*time.Millisecond)

 	tcpmodbus := [][]byte{
		{0xf1,0x01,0x00,0x00,0x00,0x08,0x29,0x3c,},
		{0xf1,0x02,0x00,0x20,0x00,0x08,0x6c,0xf6,},
	}

	//20200924号协议是三者相同的
	type udpmodbus tcpmodbus
	//type serialmodbus tcpmodbus

	tcpch :=make(chan []byte)
	go func(){
		defer close(tcpch)
		for {
			for _,b :=range tcpmodbus{
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
			for _,b :=range udpmodbus{
				udpch<-b
			}
        	/*presentTime*/ _ /*:*/= <-ticker.C/*阀门*/
        	//fmt.Println(/*presentTime.String(),":*/"发送协议已经完成一个周期")
		}
	}()

	//为了让程序顺利跑通，20200924号协议暂时抛弃了serial的配置
	// serialch :=make(chan []byte)
	// go func(){
	// 	defer close(serialch)
	// 	for {
	// 		for _,b :=range serialmodbus{
	// 			serialch<-b
	// 		}
    //     	/*presentTime*/ _ /*:*/= <-ticker.C/*阀门*/
    //     	//fmt.Println(/*presentTime.String(),":*/"发送协议已经完成一个周期")
	// 	}
	// }()
	//chs["tty0"] =tcpch

	//tcp和udp都需要建立连接，以此发来消息的节点必须是具备ip地址的
	//值后如果有另一个tcp则也可以基于旧协议进行响应的函数重构
	//也就是管道的添加操作，添加的管道同样会append到一个map中并返回给下一层
	chs["192.168.10.2"] =tcpch
	chs["192.168.11.2"]=udpch

	return chs
}

