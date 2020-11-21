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
	"fmt"

	"github.com/ziyouzy/mylib/connserver"
)

func TickerSendModbusToNouthBound(step int){
	modbusMatrix0 := [][]byte{
		{0xf1,0x01,0x00,0x00,0x00,0x08,0x29,0x3c,},
		{0xf1,0x02,0x00,0x20,0x00,0x08,0x6c,0xf6,},
	}
	time.Sleep(5*time.Second)
	client :=connserver.ClientMap()["TCPCONN:192.168.10.2"]

	//connNames0 :=[]string{"TCPCONN:192.168.10.2",}

	
	go func(){
		for{
			for _,modbus := range modbusMatrix0{
				fmt.Println("client.SendBytes(modbus):",modbus)
				client.SendBytes(modbus)
				time.Sleep(1*time.Second)
			}
		}
		// for{
		// 	time.Sleep(1*time.Second)
		// 	clients :=connserver.ClientMap()
		// 	if clients ==nil{
		// 		fmt.Println("当前connserver.ClientMap()为空")
		// 		continue
		// 	}else{
		// 		for _, name := range connNames0{
		// 			if clients[name] ==nil{
		// 				fmt.Println("当前connserver.ClientMap[",name,"]并不存在")
		// 				continue
		// 			}else{					
		// 				for _, modbus := range modbusMatrix0{
		// 					clients[name].SendBytes(modbus)
		// 					//fmt.Println("sended:",modbus)
		// 					time.Sleep(1*time.Second)
		// 				}
		// 			}
		// 		}
		// 	}
		// }
	}()
}


