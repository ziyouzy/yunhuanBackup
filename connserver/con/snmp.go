package con

import(
	"fmt"
	"time"
	"bytes"
	"strings"
	"encoding/binary"

	"github.com/k-sone/snmpgo"
)


type SnmpConn struct{
	SNMP *snmpgo.SNMP
	Address string

	oids []string
	resoids []*snmpgo.Oid
	pdu snmpgo.Pdu
	oidsBytesLen int
	milliSecondStep int

	SendCh chan []byte
	RecvCh chan []byte



	//QuitActiveEventSender chan bool
}

func (p *SnmpConn)GenerateSendCh(){
	p.SendCh = make(chan []byte)
	go func (){
		for b := range p.SendCh{
			_ =b
		}
	}()
}

func (p *SnmpConn)GenerateRecvCh(){
	p.RecvCh = make(chan []byte)

	go func (){
		defer func(){ if p.SNMP !=nil { p.Destory() };        fmt.Println("有p.SNMP意外断开了") } ()

		recvBuffer :=bytes.NewBuffer([]byte{})

		ip := []byte(strings.Split(p.Address, ":")[0]);         tag :=[]byte("snmp")
		
		/* ---------------------------------------*/
		/** 这个for的目的只是实现一个类似阀门的东西
		  * 应对的情况是，当还未执行InitActiveEventSender()的时候
		  * 先让整体流程卡住 
		  * 从而让整体的使用方式与tcp模块一样不用考虑GenerateRecvCh()与InitActiveEventSender()
		  * 的先后执行顺序
		  */
		for {
		/* ---------------------------------------*/

			if p.pdu ==nil { fmt.Println("SNMP-pdu为nil，暂时无法实现SNMP的数据传输功能，但不会跳出循环，5秒后重试");        time.Sleep(5 * time.Millisecond);        continue }
			if p.resoids ==nil { fmt.Println("SNMP-resoids为nil，暂时无法实现SNMP的数据传输功能，但不会跳出循环，5秒后重试");         time.Sleep(5 * time.Millisecond);        continue }
			if p.oidsBytesLen ==0 { fmt.Println("SNMP-oidsbytesLen为0，暂时无法实现SNMP的数据传输功能，但不会跳出循环，5秒后重试");         time.Sleep(5 * time.Millisecond);        continue }

			for i :=0; i<=p.oidsBytesLen ; i++{	
				if i == p.oidsBytesLen { i=0 }

				tempvarbind :=p.pdu.VarBinds().MatchOid(p.resoids[i])
				tempBuf ,err:=tempvarbind.Marshal()
				if err !=nil { fmt.Println("当序列化一个snmp返回值时发生错误，存在问题的varBind对象为:",tempvarbind);        continue /* 只会回到内层for循环的初始位置 */ }
				fmt.Println("oid and request([]byte->string):",p.oids[i],string(tempBuf))

				/*核心，为原始字节数组依次添加了ip,时间,tag*/

				timeStamp :=make([]byte,8);        binary.BigEndian.PutUint64(timeStamp, uint64(time.Now().UnixNano()))			

				recvBuffer.Reset();        _,_ = recvBuffer.Write(ip);        _,_ = recvBuffer.Write([]byte(" /-/ "));        _,_ = recvBuffer.Write(timeStamp);
	 			_,_ = recvBuffer.Write([]byte(" /-/ "));       _,_ = recvBuffer.Write(tag);        _,_ = recvBuffer.Write([]byte(" /-/ "));         _,_ = recvBuffer.Write(tempBuf)  

				p.RecvCh <-recvBuffer.Bytes();

				if p.milliSecondStep ==0{
					time.Sleep(5 * time.Millisecond)
				}else{
					time.Sleep(time.Duration(p.milliSecondStep) * time.Millisecond)
				}
			}			
		}
	}()
}

func (p *SnmpConn)ActiveBaseFunctions(oidsbytes [][]byte, step int){
	if  (oidsbytes ==nil)||(len(oidsbytes) ==0)        { panic("SNMP模块在初始化过程中发生错误:oidsbytes不能为nil，且数量不可为0") }
	p.milliSecondStep =step;        p.oidsBytesLen =len(oidsbytes)

	for b := range oidsbytes{
		p.oids =append(p.oids,string(b))
	}

	var err error

	fmt.Println("SNMP-转化后的oids字符串数组为:",p.oids)

	p.resoids, err = snmpgo.NewOids(p.oids);        if err != nil {fmt.Println("SNMP-初始化resOid失败:",err);        return}
	p.pdu, err = p.SNMP.GetRequest(p.resoids);        if err != nil {fmt.Println("SNMP-初始化Pdu失败:",err);        return}

	if p.pdu.ErrorStatus() != snmpgo.NoError        { fmt.Println(p.pdu.ErrorStatus(), p.pdu.ErrorIndex()) }

	fmt.Println("SNMP-设备列表:",p.pdu.VarBinds())//获取绑定节点设备列表
}

// func (p *SnmpConn)SendBytes(b []byte) {
// 	_ =b
// 	fmt.Println("SNMP-禁止使用SnmpConn的SendBytes方法，该方法只为实现mylib.connserver.con.Con接口而存在")
// }


func (p *SnmpConn)Destory(){
	defer func(){ if p.SNMP !=nil { p.SNMP.Close()} }()
	close(p.RecvCh)
}