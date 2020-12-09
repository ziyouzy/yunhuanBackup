package di

import(
	"strings"
	//"strconv"
	"fmt"
	"bytes"
	//"encoding/binary"
	"github.com/imroc/biu"
)

type DI_YOUREN_USRIO808EWR_20200924 struct{
	NodeType string
	ProtocolType string
	Handler string
	Tag string

	TimeUnixNano uint64

	Raw []byte
	DI1 []byte 
	DI2 []byte
	DI3 []byte
	DI4 []byte
	DI5 []byte
	DI6 []byte
	DI7 []byte
	DI8 []byte
}

func (p *DI_YOUREN_USRIO808EWR_20200924)FullOf(){
	if bytes.Contains(p.Raw, []byte("timeout")){
		p.DI8 =[]byte("timeout")
		p.DI7 =[]byte("timeout")
		p.DI6 =[]byte("timeout")
		p.DI5 =[]byte("timeout")
		p.DI4 =[]byte("timeout")
		p.DI3 =[]byte("timeout")
		p.DI2 =[]byte("timeout")
		p.DI1 =[]byte("timeout")
		return
	}

	var binStr string
	if bytes.Index(p.Raw,[]byte{0x49,0x4f})==0&&strings.Compare(p.Tag,"tcpsocket")==0{
		fmt.Println("p.Raw:",p.Raw)
		fmt.Println("p.Raw[7:8]:",p.Raw[7:8])
		binStr =biu.BytesToBinaryString(p.Raw[7:8])
		fmt.Println("binStr:",binStr)
	}else if strings.Compare(p.Tag,"serial")==0{
		binStr =biu.BytesToBinaryString(p.Raw[5:6])
	}

	if len(binStr) ==10{
		//binStr[0]=="[
		p.DI8 =append(p.DI8,binStr[1])
		p.DI7 =append(p.DI7,binStr[2])
		p.DI6 =append(p.DI6,binStr[3])
		p.DI5 =append(p.DI5,binStr[4])
		p.DI4 =append(p.DI4,binStr[5])
		p.DI3 =append(p.DI3,binStr[6])
		p.DI2 =append(p.DI2,binStr[7])
		p.DI1 =append(p.DI1,binStr[8])
		//binStr[9]=="]
	}else{
		//fmt.Println("len(binStr):",len(binStr),"binStr:",binStr,"binStr[0]:",binStr[0],"binStr[9]:",binStr[9],"binStr[1]:",binStr[1],"binStr[8]:",binStr[8])
		p.DI8 = []byte("undefined")
		p.DI7 = []byte("undefined")
		p.DI6 = []byte("undefined")
		p.DI5 = []byte("undefined")
		p.DI4 = []byte("undefined")
		p.DI3 = []byte("undefined")
		p.DI2 = []byte("undefined")
		p.DI1 = []byte("undefined")
	}
	//fmt.Println("p.DI8(string):",string(p.DI8),"p.DI7:",p.DI7,"p.DI6:",p.DI6,"p.DI5:",p.DI5,"p.DI4:",p.DI4,"p.DI3:",p.DI3,"p.DI2:",p.DI2,"p.DI1:",p.DI1)
	return
}


func (p *DI_YOUREN_USRIO808EWR_20200924)SelectHandlerAndTag() (string,string){
	return p.Handler, p.Tag
}



func (p *DI_YOUREN_USRIO808EWR_20200924)SelectOneValueAndTimeUnixNano(nodedohandler string, nodedotag string, nodedoname string) ([]byte,uint64){
	if strings.Compare(p.Handler,nodedohandler)!=0||strings.Compare(p.Tag, nodedotag)!=0{ return nil, 0 }

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
		return nil, 0
	}
}