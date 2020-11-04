//针对串口的协议，不过似乎暂时用不到
package protocol

//获取串口在线表
func ProtocolPrepareSerialPorts_YunHuan20200924()([]string, []int, []int, []bool){
	return []string{"tty1","tty2","tty3",}, []int{9600,9600,2400,}, []int{500,500,500,}, []bool{true,true,true,}
}

func ProtocolPrepareSmsMgr_YunHuan20200924() []string{
	return []string{"tty1",}
}