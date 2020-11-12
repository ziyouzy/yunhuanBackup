/*
会存在一个service函数里既会包含serverconn模块的接收操作，也会包含发送操作
当然也会有较为简单的service
如非service层包含了service的相关工作，如将为nodedo设计的service方法，在nodedo包，函数名会标有"Service"，方法会在真正的service包实现业务整合，但是在service包内，如非必要，函数名与方法名不需要包含"Service"
service很可能会分为北向与南向因此service包内的函数命名规则有可能会是"SouthServiceTickerSendNodeDo()"或"NouthServiceTickerSendModbus()"
*/

//拿NouthServiceTickerSendModbus()和SouthServiceTickerSendNodeDo()举例，目前的核心问题在于NodeDoCh和SendModbusCh这两个管道在哪里合成，是在主函数中比较合理，还是在Service里比较合理
package service

func TickerSendModbusToNouthBound(connserver){
	//map[string]connCh,connch分别为tcp与udp
	conchmap := protocol.ProtocolPrepareSendTicketMgr_YunHuan20200924()
	for k,ch :=range conchmap{
		//mark 可能会是"192.168.10.2"，"192.168.10.1",或者是"serial"
		go func(key string){
			//问题其实是因为range神坑与单独开携程使用了这个k共同导致的，或者说，毕竟range的时候k和v都只有一个副本，你开携程相当于把这一个副本作为参数传递进了这个携程，无论如何，肯定只会是一个值
			for b := range ch{
				if(connserver.ConnMap[key] !=nil){
					connserver.ConnMap[key].SendBytes(b)
				}else{
					fmt.Println("mark为:",key,"的设备暂时并不在册于tcphandler.ConnMap中，往设备客户端尽快上线")
				}
			}
		}(k)
	}
}
