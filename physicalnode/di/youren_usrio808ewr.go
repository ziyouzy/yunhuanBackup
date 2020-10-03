package di

import(
	"strings"
	"strconv"
	//"fmt"
	//"bytes"
	//"encoding/binary"
)

type DI_YOUREN_USRIO808EWR_20200924 struct{
	NodeType string
	ProtocolType string

	InputTime string
	Value string
	Mark string

	//唯一标识，很重要，之后很多功能都需要通过他来实现
	//如494f3031f10201,代表了IO01-主控-DO，之后在生成UINode时就需要用到他了
	//在这一层只做一下简单的记录
	Handler string
	Tag string
	//nodename为临时变量	

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
	if strings.Contains(p.Value, "timeout"){
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

	if strings.Compare(p.Tag,"tcpsocket")==0&&strings.Index(p.Value,"494f")==0{
		hex,err := strconv.ParseInt(p.Value[12:16],16,0)
		if err ==nil{
			tempStr :=string([]byte(strconv.FormatInt(hex,2)[1:]))
			p.DI8 =string([]byte(tempStr)[0:1])
			p.DI7 =string([]byte(tempStr)[1:2])
			p.DI6 =string([]byte(tempStr)[2:3])
			p.DI5 =string([]byte(tempStr)[3:4])
			p.DI4 =string([]byte(tempStr)[4:5])
			p.DI3 =string([]byte(tempStr)[5:6])
			p.DI2 =string([]byte(tempStr)[6:7])
			p.DI1 =string([]byte(tempStr)[7:8])
			return
		}
	}

	if strings.Compare(p.Tag,"serial")==0{
		hex,err := strconv.ParseInt(p.Value[8:12],16,0)
		if err ==nil{
			tempStr :=string([]byte(strconv.FormatInt(hex,2)[1:]))
			p.DI8 =string([]byte(tempStr)[0:1])
			p.DI7 =string([]byte(tempStr)[1:2])
			p.DI6 =string([]byte(tempStr)[2:3])
			p.DI5 =string([]byte(tempStr)[3:4])
			p.DI4 =string([]byte(tempStr)[4:5])
			p.DI3 =string([]byte(tempStr)[5:6])
			p.DI2 =string([]byte(tempStr)[6:7])
			p.DI1 =string([]byte(tempStr)[7:8])
			return
		}
	}

	p.DI8 = "undefined"
	p.DI7 = "undefined"
	p.DI6 = "undefined"
	p.DI5 = "undefined"
	p.DI4 = "undefined"
	p.DI3 = "undefined"
	p.DI2 = "undefined"
	p.DI1 = "undefined"
	return
}

func (p *DI_YOUREN_USRIO808EWR_20200924)GetNodeType() string{
	return p.NodeType
}

func (p *DI_YOUREN_USRIO808EWR_20200924)GetHandler() string{
	return p.Handler
}

func (p *DI_YOUREN_USRIO808EWR_20200924)GetRaw() (string,string,string,string,string,string,string){
	return p.NodeType, p.ProtocolType, p.Tag, p.InputTime, p.Value, p.Mark, p.Handler
}

func (p *DI_YOUREN_USRIO808EWR_20200924)SeleteOneValueByProtocol(protocolhandler string, protocoltag string, protocolnodename,string) string{
	if strings.Compare(p.Handler,protocolhandler)!=0||strings.Compare(p.Tag, protocoltag)!=0{
		return nil
	}

	switch (protocolnodename){
	case "DI8":
		return p.DI8
	case "DI7":
		return p.DI7
	case "DI6":
		return p.DI6
	case "DI5":
		return p.DI5
	case "DI4":
		return p.DI4
	case "DI3":
		return p.DI3
	case "DI2":
		return p.DI2
	case "DI1":
		return p.DI1
	default :
		return nil
	}
}