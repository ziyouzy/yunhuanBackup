package di

import(
	"strings"
	//"strconv"
	//"fmt"
	"bytes"
	//"encoding/binary"
	"github.com/imroc/biu"
)

type DI_YOUREN_USRIO808EWR_20200924 struct{
	ProtocolNodeType string
	//ProtocolType string
	Handler string
	Tag string

	TimeUnixNano uint64

	Raw []byte

	DI1 string 
	DI2 string
	DI3 string
	DI4 string
	DI5 string
	DI6 string
	DI7 string
	DI8 string
}

func (p *DI_YOUREN_USRIO808EWR_20200924)FullOf(){
	if bytes.Contains(p.Raw, []byte("timeout")){
		p.DI8 ="timeout"
		p.DI7 ="timeout"
		p.DI6 ="timeout"
		p.DI5 ="timeout"
		p.DI4 ="timeout"
		p.DI3 ="timeout"
		p.DI2 ="timeout"
		p.DI1 ="timeout"
		return
	}

	var binStr string
	if bytes.Index(p.Raw,[]byte{0x49,0x4f})==0&&strings.Compare(p.Tag,"tcpsocket")==0{
		binStr =biu.BytesToBinaryString(p.Raw[7:8])
	}else if strings.Compare(p.Tag,"serial")==0{
		binStr =biu.BytesToBinaryString(p.Raw[5:6])
	}

	if len(binStr) ==10{
		p.DI8 =string(binStr[1])
		p.DI7 =string(binStr[2])
		p.DI6 =string(binStr[3])
		p.DI5 =string(binStr[4])
		p.DI4 =string(binStr[5])
		p.DI3 =string(binStr[6])
		p.DI2 =string(binStr[7])
		p.DI1 =string(binStr[8])
	}else{
		p.DI8 = "undefined"
		p.DI7 = "undefined"
		p.DI6 = "undefined"
		p.DI5 = "undefined"
		p.DI4 = "undefined"
		p.DI3 = "undefined"
		p.DI2 = "undefined"
		p.DI1 = "undefined"
	}
	return
}


func (p *DI_YOUREN_USRIO808EWR_20200924)SelectHandlerAndTag() (string,string){
	return p.Handler, p.Tag
}



func (p *DI_YOUREN_USRIO808EWR_20200924)SelectOneValueAndTimeUnixNano(nodedohandler string, nodedotag string, nodedoname string) (string,uint64){
	if strings.Compare(p.Handler,nodedohandler)!=0||strings.Compare(p.Tag, nodedotag)!=0{ return "", 0 }

	switch (nodedoname){
	case "di8", "DI8":
		return p.DI8,p.TimeUnixNano
	case "di7", "DI7":
		return p.DI7,p.TimeUnixNano
	case "di6", "DI6":
		return p.DI6,p.TimeUnixNano
	case "di5", "DI5":
		return p.DI5,p.TimeUnixNano
	case "di4", "DI4":
		return p.DI4,p.TimeUnixNano
	case "di3", "DI3":
		return p.DI3,p.TimeUnixNano
	case "di2", "DI2":
		return p.DI2,p.TimeUnixNano
	case "di1", "DI1":
		return p.DI1,p.TimeUnixNano
	default :
		return "", 0
	}
}