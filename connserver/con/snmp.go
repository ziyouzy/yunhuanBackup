package con



import(
	"fmt"
	"time"
	"github.com/k-sone/snmpgo"
)

// type SNMP struct {
// 	conn   net.Conn
// 	args   *SNMPArguments
// 	engine *snmpEngine
// }

type SnmpConn struct{
	SNMP *snmpgo.SNMP

	oids []string
	resoids []*snmpgo.Oid
	pdu snmpgo.Pdu

	quit chan bool
}

func (p *SnmpConn)GenerateRecvCh() chan []byte{
	ch := make(chan []byte)
	go func (){
		defer p.SNMP.Close()
		defer close(ch)
		defer func(){p.quit<-true}()
		
		l :=len(p.resoids)
		if l==0{fmt.Println("SNMP-resoids为空，无法创建SNMP的数据管道");return}
		for i :=0;i<=l;i++{
			if i==l{i=0}
			tempvarbind :=p.pdu.VarBinds().MatchOid(p.resoids[i])
			tempbytes ,err:=tempvarbind.Marshal()
			if err !=nil{fmt.Println("当序列化一个snmp返回值时发生错误，存在问题的varBind对象为:",tempvarbind);continue}
			fmt.Println("oid and request([]byte->string):",p.oids[i],string(tempbytes))
			ch <-tempbytes
			time.Sleep(5*time.Millisecond)
		}
	}()
	return ch
}

func (p *SnmpConn)InitActiveEventSender(oidsbytes [][]byte){
	var err error
	for b := range oidsbytes{
		p.oids =append(p.oids,string(b))
	}

	fmt.Println("SNMP-转化后的oids字符串数组为:",p.oids)
	p.resoids, err = snmpgo.NewOids(p.oids)
	if err != nil {fmt.Println("SNMP-初始化resOid失败:",err);return}

	p.pdu, err= p.SNMP.GetRequest(p.resoids)
	if err != nil {fmt.Println("SNMP-初始化Pdu失败:",err);return}

	if p.pdu.ErrorStatus() != snmpgo.NoError{fmt.Println(p.pdu.ErrorStatus(), p.pdu.ErrorIndex())}

	fmt.Println("SNMP-设备列表:",p.pdu.VarBinds())//获取绑定节点设备列表
}

func (p *SnmpConn)SendBytes(b []byte) {
	_ =b
	fmt.Println("SNMP-禁止使用SnmpConn的SendBytes方法，该方法只为实现mylib.connserver.con.Con接口而存在")
}


