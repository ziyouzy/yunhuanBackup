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

var (

	TcpModbus1 = [][]byte{
		{0xf1,0x01,0x00,0x00,0x00,0x08,0x29,0x3c,},
		{0xf1,0x02,0x00,0x20,0x00,0x08,0x6c,0xf6,},
	}

	SnmpOids1 = [][]byte{
		[]byte("test1"),
		[]byte("test2"),
	}

) 


type Con interface{
	GenerateRecvCh()
	InitActiveEventSender([][]byte, int)
	Destory()
}

func NewTcpCon(c net.Conn, needcrc bool, conntimeout int, activeeventstep int) (string, Con, chan []byte , chan []byte ) {
	_ = conntimeout//最后再去设计长连接专用的超时机制

	key :=fmt.Sprintf("TCPCONN:%s",strings.Split(c.RemoteAddr().String(),":")[0])
	switch key{
	case "TCPCONN:192.168.10.2":
		tcpConn := &TcpConn{ Conn :c, NeedCRC :needcrc, }

		//if conntimeout >0        { tcpConn.Conn.SetReadDeadline(time.Duration(conntimeout) * time.Second) }

		tcpConn.GenerateRecvCh()
		tcpConn.GenerateSendCh()
		tcpConn.InitActiveEventSender(TcpModbus1, activeeventstep)
		return key,        tcpConn,        tcpConn.RecvCh,        nil

	default:
		tcpConn := &TcpConn{ Conn :c, NeedCRC :needcrc, }

		//if conntimeout >0        { tcpConn.Conn.SetReadDeadline(time.Duration(conntimeout) * time.Second) }

		tcpConn.GenerateRecvCh()
		tcpConn.GenerateSendCh()
		return key,        tcpConn,        tcpConn.RecvCh,        tcpConn.SendCh 
	}
}

func NewSNMPCon(c *snmpgo.SNMP, ip string, conntimeout int, activeeventstep int)(string, Con, chan []byte , chan []byte ){
	_ = conntimeout

	key := fmt.Sprintf("SNMPCONN:%s",ip)
	switch key{
	default:
		snmpConn := &SnmpConn{SNMP :c,}
		snmpConn.InitActiveEventSender(SnmpOids1, activeeventstep)
		snmpConn.GenerateRecvCh()
		return key,        snmpConn,        snmpConn.RecvCh,        nil
	}
}