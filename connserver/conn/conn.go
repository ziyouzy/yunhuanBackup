
package connclient

const (
	TIMEFORMAT = "200601021504.05.000000000"
	NORMALTIMEFORMAT1 = "2006-01-02 15:04:05"
	NORMALTIMEFORMAT2 ="20060102150405"
	NEEDCRC =true
	NOCRC =false
)

func NewConnClient(c net.Conn, needcrc bool)(string,ConnClient,int){
	timeoutsec :=5
	c.SetReadDeadline(now.Add(time.Second * timeoutsec))
	return fmt.Sprintf("TCPCONN:%s",strings.Split(c.RemoteAddr().String(),":")[0]) , &TcpConn{
			Conn :c,
			NeedCRC :needcrc,
		},timeoutsec
	}	
}

type ConnClient interface{
	GenerateRecvCh() chan([]byte)
	SendBytes([]byte)
}
