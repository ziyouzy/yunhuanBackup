package connserver

import(
	"fmt"


	"github.com/k-sone/snmpgo"
	

	"github.com/ziyouzy/mylib/connserver/con"
	"github.com/ziyouzy/mylib/conf"
)

//snmp和tcp是不同的，tcp是客户端所以需要for循环实时监听新的连接
//snmp的本质是客户端，所以就不需要实时监听任何客户端了
//而此函数依然又一个会长期存在于内存的匿名携程函数，因为他需要守护数据管道
func (p *ConnServer)SnmpRecvCh(port string){
	snmpv1 ,err :=snmpgo.NewSNMP(snmpgo.SNMPArguments{
		Version: snmpgo.V1,
		Address: fmt.Sprintf("192.168.10.3%s",port),//目标主机地址
		Retries: 1,
		Community: "public",
	})
	
	if err != nil{fmt.Println("snmp初始化协议栈时失败",err);return}

	if err  = snmpv1.Open(); err != nil {fmt.Println("snmp开启服务时失败",err);return}

	key,client, timeout :=con.NewSNMPCon(snmpv1,"192.168.10.3")
	switch key{
	case "SNMPCONN:192.168.10.3":
		client.InitOwnActiveEventSender(conf.SnmpOids1)
	}
	p.ConnClientMap[key] =client
	clientrecvch :=client.GenerateRecvCh()
	if clientrecvch ==nil{fmt.Println("SNMP-创建数据管道失败");return}

	go func(){
		defer delete(p.ConnClientMap,key)
		defer fmt.Println("该设备", timeout, "秒无应答，连接将会从ConnClientMap中删除：",key)
		for b := range clientrecvch{
			p.ServerRecvCh<-b
		} 
	}()

	fmt.Println("p.ConnClientMap updated:",p.ConnClientMap)
}
