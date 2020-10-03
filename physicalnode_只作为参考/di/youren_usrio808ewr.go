

type DI_YOUREN_USRIO808EWR_20200924 struct{
	NodeType string
	ProtocolType string

	Tag string
	InputTime string
	Value string
	Ip string

	//唯一标识，很重要，之后很多功能都需要通过他来实现
	//如494f3031f10201,代表了IO01-主控-DO，之后在生成UINode时就需要用到他了
	//在这一层只做一下简单的记录
	Handler string	

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
	
	if strings.Index(p.Value,"494f")==0{
		var tmp int
		bytesBuffer := bytes.NewBuffer([]byte(p.Value)[12:16]))
		if err := binary.Read(bytesBuffe, binary.BigEndian, &tmp);err ==nil{
			tempStr =string([]byte(strconv.FormatInt(tmp,2))[1:])

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

func (p *DI_YOUREN_USRIO808EWR_20200924)GetRaw() (string,string,string,string,string,string){
	return p.NodeType, p.ProtocolType, p.Tag, p.ImportTime, p.Value, p.Ip
}