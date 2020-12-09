package do

import(
	"strings"
	//"strconv"
	"bytes"
	//"fmt"
	//"encoding/binary"
	"github.com/imroc/biu"
)

type DO_YOUREN_USRIO808EWR_20200924 struct{
	ProtocolNodeType string
	//ProtocolType string
	Handler string
	Tag string

	TimeUnixNano uint64

	Raw []byte

	DO1 string 
	DO2 string 
	DO3 string 
	DO4 string 
	DO5 string 
	DO6 string 
	DO7 string 
	DO8 string 
}

func (p *DO_YOUREN_USRIO808EWR_20200924)FullOf(){
	if bytes.Contains(p.Raw, []byte("timeout")){
		p.DO8 ="timeout"
		p.DO7 ="timeout"
		p.DO6 ="timeout"
		p.DO5 ="timeout"
		p.DO4 ="timeout"
		p.DO3 ="timeout"
		p.DO2 ="timeout"
		p.DO1 ="timeout"
		return
	}
	
	var binStr string
	if bytes.Index(p.Raw,[]byte{0x49,0x4f})==0&&strings.Compare(p.Tag,"tcpsocket")==0{
		binStr =biu.BytesToBinaryString(p.Raw[7:8])
	}else if strings.Compare(p.Tag,"serial")==0{
		binStr =biu.BytesToBinaryString(p.Raw[5:6])
	}

	if len(binStr) ==10{
		p.DO8 =string(binStr[1])
		p.DO7 =string(binStr[2])
		p.DO6 =string(binStr[3])
		p.DO5 =string(binStr[4])
		p.DO4 =string(binStr[5])
		p.DO3 =string(binStr[6])
		p.DO2 =string(binStr[7])
		p.DO1 =string(binStr[8])
	}else{
		p.DO8 = "undefined"
		p.DO7 = "undefined"
		p.DO6 = "undefined"
		p.DO5 = "undefined"
		p.DO4 = "undefined"
		p.DO3 = "undefined"
		p.DO2 = "undefined"
		p.DO1 = "undefined"
	}
	return
}


func (p *DO_YOUREN_USRIO808EWR_20200924)SelectHandlerAndTag() (string, string){
	return p.Handler, p.Tag
}


func (p *DO_YOUREN_USRIO808EWR_20200924)SelectOneValueAndTimeUnixNano(nodedohandler string, nodedotag string, nodedoname string) (string,uint64){
	if strings.Compare(p.Handler,nodedohandler)!=0||strings.Compare(p.Tag, nodedotag)!=0{ return "", 0 }
	switch (nodedoname){
	case "do8", "DO8":
		return p.DO8,p.TimeUnixNano
	case "do7", "DO7":
		return p.DO7,p.TimeUnixNano
	case "do6", "DO6":
		return p.DO6,p.TimeUnixNano
	case "do5", "DO5":
		return p.DO5,p.TimeUnixNano
	case "do4", "DO4":
		return p.DO4,p.TimeUnixNano
	case "do3", "DO3":
		return p.DO3,p.TimeUnixNano
	case "do2", "DO2":
		return p.DO2,p.TimeUnixNano
	case "do1", "DO1":
		return p.DO1,p.TimeUnixNano
	default :
		return "", 0
	}
}

