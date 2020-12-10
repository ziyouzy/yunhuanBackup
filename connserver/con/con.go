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
	sendBytes([]byte)
	InitActiveEventSender([][]byte)
}

func NewTcpCon(c net.Conn, needcrc bool, conntimeout int)(string, Con, []byte chan, []byte chan){
	
	key :=fmt.Sprintf("TCPCONN:%s",strings.Split(tcpConn.Conn.RemoteAddr().String(),":")[0]
	switch key{
	case "TCPCONN:192.168.10.2":
		tcpConn := &TcpConn{ Conn :c, NeedCRC :needcrc, }

		if conntimeout >0        { tcpConn.Conn.SetReadDeadline(time.Dustion(conntimeout) * time.Second) }

		tcpConn.GenerateRecvCh()
		tcpConn.GenerateSendCh()
		tcpConn.InitActiveEventSender()
		return key,        tcpConn,        tcpConn.RecvCh,        nil

	default:
		tcpConn := &TcpConn{ Conn :c, NeedCRC :needcrc, }

		if conntimeout >0        { tcpConn.Conn.SetReadDeadline(time.Dustion(conntimeout) * time.Second) }

		tcpConn.GenerateRecvCh()
		tcpConn.GenerateSendCh()
		return key,        tcpConn,        tcpConn.RecvCh,        tcpConn.SendCh 
	}
}

func NewSNMPCon(c *snmpgo.SNMP, ip string, conntimeout int)(string, Con, int){
	return fmt.Sprintf("SNMPCONN:%s",ip), &SnmpConn{SNMP :c,},timeoutsec
}