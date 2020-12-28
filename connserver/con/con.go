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
	GenerateSendCh()
	ActiveBaseFunctions([][]byte, int)
	Destory()
}

func NewTcpCon(c net.Conn, needcrc bool, conntimeout int, activeeventstep int) (string, Con, chan []byte , chan []byte ) {
	_ = conntimeout//最后再去设计长连接专用的超时机制

	key :=fmt.Sprintf("TCPCONN:%s", strings.Split(c.RemoteAddr().String(),":")[0])
	switch key{
	case "TCPCONN:192.168.10.2":
		tcpConn := &TcpConn{ Conn :c, NeedCRC :needcrc, }

		//if conntimeout >0        { tcpConn.Conn.SetReadDeadline(time.Duration(conntimeout) * time.Second) }

		/* ---------------------------------------*/

		/** 对于TCP来说下面三件事没有先后顺序的特殊要求
		  * 首先虽然InitActiveEventSender需要先make SendCh后才能起作用，但是内部已经设计好了必要措施
		  * 那就是如果发现还未make SendCh,就会把需要处理的数据抛弃
		  * 而对于GenerateRecvCh()，以及他所负责的创建RecvCh管道来说
		  * 就更没有太多的相关性了，因为他只和从套接字读取数据这一事件有关
		  */

		tcpConn.ActiveBaseFunctions(TcpModbus1, activeeventstep)
		tcpConn.GenerateSendCh()
		tcpConn.GenerateRecvCh()

		/* ---------------------------------------*/
		
		return key,        tcpConn,        tcpConn.RecvCh,        nil

	default:
		tcpConn := &TcpConn{ Conn :c, NeedCRC :needcrc, }

		//if conntimeout >0        { tcpConn.Conn.SetReadDeadline(time.Duration(conntimeout) * time.Second) }

		tcpConn.GenerateSendCh()
		tcpConn.GenerateRecvCh()

		return key,        tcpConn,        tcpConn.RecvCh,        tcpConn.SendCh 
	}
}

func NewSNMPCon(c *snmpgo.SNMP, ip string, port string, conntimeout int, activeeventstep int)(string, Con, chan []byte , chan []byte ){
	_ = conntimeout

	address := fmt.Sprintf("%s:%s",ip, port);        key := fmt.Sprintf("SNMPCONN:%s", address)
	
	switch key{
	default:
		snmpConn := &SnmpConn{
			SNMP :        c,
			Address:        address,
		}

		/* ---------------------------------------*/

		/** SNMP与TCP不同
		  * 这里的InitActiveEventSender和GenerateRecvCh存在着先后顺序的问题
		  * InitActiveEventSender必须在前
		  * 但是为了书写规范，TCP虽然无所谓，但是也把InitActiveEventSender放在了前边
		  */

		snmpConn.ActiveBaseFunctions(SnmpOids1, activeeventstep)
		snmpConn.GenerateRecvCh()

		/* ---------------------------------------*/

		return key,        snmpConn,        snmpConn.RecvCh,        nil
	}
}