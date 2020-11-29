package con

import(
	"net"
	//"time"
	"fmt"
	"strings"
	"github.com/k-sone/snmpgo"
)

const (
	TIMEFORMAT = "200601021504.05.000000000"
	NORMALTIMEFORMAT1 = "2006-01-02 15:04:05"
	NORMALTIMEFORMAT2 ="20060102150405"
	NEEDCRC =true
	NOCRC =false
)


type Con interface{
	GenerateRecvCh() chan([]byte)
	SendBytes([]byte)
	InitOwnActiveEventSender([][]byte)
}

func NewTcpCon(c net.Conn, needcrc bool)(string,Con,int){
	timeoutsec :=5
	//now :=time.Now()
	//c.SetReadDeadline(now.Add(time.Second * time.Duration(timeoutsec)))
	return fmt.Sprintf("TCPCONN:%s",strings.Split(c.RemoteAddr().String(),":")[0]) , &TcpConn{Conn :c,NeedCRC :needcrc,},timeoutsec
}

func NewSNMPCon(c *snmpgo.SNMP, ip string)(string,Con,int){
	timeoutsec :=5
	return fmt.Sprintf("SNMPCONN:%s",ip), &SnmpConn{SNMP :c,},timeoutsec
}