/*
会存在一个service函数里既会包含serverconn模块的接收操作，也会包含发送操作
当然也会有较为简单的service
如非service层包含了service的相关工作，如将为nodedo设计的service方法，在nodedo包，函数名会标有"Service"，方法会在真正的service包实现业务整合，但是在service包内，如非必要，函数名与方法名不需要包含"Service"
service很可能会分为北向与南向因此service包内的函数命名规则有可能会是"SouthServiceTickerSendNodeDo()"或"NouthServiceTickerSendModbus()"
*/

//拿NouthServiceTickerSendModbus()和SouthServiceTickerSendNodeDo()举例，目前的核心问题在于NodeDoCh和SendModbusCh这两个管道在哪里合成，是在主函数中比较合理，还是在Service里比较合理
package service

import(
	"time"
	//"fmt"

	"github.com/ziyouzy/mylib/connserver"
)
func TickerSendModbusToNouthBound_RangeTest(){
	modbusMatrix0 := [][]byte{
		{0xf1,0x01,0x00,0x00,0x00,0x08,0x29,0x3c,},
		{0xf1,0x02,0x00,0x20,0x00,0x08,0x6c,0xf6,},
	}
	time.Sleep(5*time.Second)
	
	client :=connserver.ClientMap()["TCPCONN:192.168.10.2"]

	for{
		for _,modbus := range modbusMatrix0{
			client.SendBytes(modbus)
			time.Sleep(5*time.Millisecond)
		}
	}
}

func TickerSendModbusToNouthBound(step int){	
	modbusMatrix0 := [][]byte{
		{0xf1,0x01,0x00,0x00,0x00,0x08,0x29,0x3c,},
		{0xf1,0x02,0x00,0x20,0x00,0x08,0x6c,0xf6,},
	}
	time.Sleep(5*time.Second)

	client :=connserver.ClientMap()["TCPCONN:192.168.10.2"]
	for i :=0;i<=2;i++{
		if i==2{i=0}
		client.SendBytes(modbusMatrix0[i])
		time.Sleep(5*time.Millisecond)
	}
}



