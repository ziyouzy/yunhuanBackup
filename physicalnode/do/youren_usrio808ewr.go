package do

import(
	"strings"
	//"strconv"
	"bytes"
	"fmt"
	//"encoding/binary"
	"github.com/imroc/biu"
)

type DO_YOUREN_USRIO808EWR_20200924 struct{
	NodeType string
	ProtocolType string
	Handler string
	Tag string

	TimeUnixNano uint64

	Raw []byte
	DO1 []byte 
	DO2 []byte 
	DO3 []byte 
	DO4 []byte 
	DO5 []byte 
	DO6 []byte 
	DO7 []byte 
	DO8 []byte 
}

func (p *DO_YOUREN_USRIO808EWR_20200924)FullOf(){
	if bytes.Contains(p.Raw, []byte("timeout")){
		p.DO8 =[]byte("timeout")
		p.DO7 =[]byte("timeout")
		p.DO6 =[]byte("timeout")
		p.DO5 =[]byte("timeout")
		p.DO4 =[]byte("timeout")
		p.DO3 =[]byte("timeout")
		p.DO2 =[]byte("timeout")
		p.DO1 =[]byte("timeout")
		return
	}
	
	var binStr string
	if bytes.Index(p.Raw,[]byte{0x49,0x4f})==0&&strings.Compare(p.Tag,"tcpsocket")==0{
		binStr =biu.BytesToBinaryString(p.Raw[7:8])
	}else if strings.Compare(p.Tag,"serial")==0{
		binStr =biu.BytesToBinaryString(p.Raw[5:6])
	}

	if len(binStr) ==10{
		//binStr[0]=="[
		p.DO8 =append(p.DO8,binStr[1])
		p.DO7 =append(p.DO7,binStr[2])
		p.DO6 =append(p.DO6,binStr[3])
		p.DO5 =append(p.DO5,binStr[4])
		p.DO4 =append(p.DO4,binStr[5])
		p.DO3 =append(p.DO3,binStr[6])
		p.DO2 =append(p.DO2,binStr[7])
		p.DO1 =append(p.DO1,binStr[8])
		//binStr[9]=="]
	}else{
		p.DO8 = []byte("undefined")
		p.DO7 = []byte("undefined")
		p.DO6 = []byte("undefined")
		p.DO5 = []byte("undefined")
		p.DO4 = []byte("undefined")
		p.DO3 = []byte("undefined")
		p.DO2 = []byte("undefined")
		p.DO1 = []byte("undefined")
	}

	fmt.Println("p.DO8:",p.DO8,"p.DO7:",p.DO7,"p.DO6:",p.DO6,"p.DO5:",p.DO5,"p.DO4:",p.DO4,"p.DO3:",p.DO3,"p.DO2:",p.DO2,"p.DO1:",p.DO1)
	return
}


func (p *DO_YOUREN_USRIO808EWR_20200924)SelectHandlerAndTag() (string, string){
	return p.Handler, p.Tag
}


func (p *DO_YOUREN_USRIO808EWR_20200924)SelectOneValueAndTimeUnixNano(nodedohandler string, nodedotag string, nodedoname string) ([]byte,uint64){
	if strings.Compare(p.Handler,nodedohandler)!=0||strings.Compare(p.Tag, nodedotag)!=0{ return nil, 0 }
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
		return nil, 0
	}
}

