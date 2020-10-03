package conf

func NewConfNode(physialnode)ConfNode {
	for k,v := range confNodeMap{
		handler,tag,nodename :=strings.Split(k,"-")
		valueString :=physialnode.SeleteOneValueByProtocol(handler, tag, nodename)
		switch v["type"]{
		case "BOOL":
			var confnode BoolenConfNode
			mapstructure.Decode(v, confnode)
			return confnode
		case "INT":
			var confnode IntConfNode
			mapstructure.Decode(v, confnode)
			return confnode
		case "FLOAT32":
			var confnode Float32ConfNode
			mapstructure.Decode(v, confnode)
			return confnode
		case "COMMON", "STRING":
			var confnode CommonConfNode 
			mapstructure.Decode(v, confnode)
			return confnode
		default:
			return nil
		}
	}
}

type ConfNode interface{
	CountPhysicalNode(string)
	GetSystemAndModuleString()(string, string)
	GetJson()[]byte
	GetSystemModuleAndCountJSON(string)(string, string, []byte){
}


type BoolenConfNode struct{
	System string
	Module string
	ModuleId int
	IsOnline bool
	IsNormal bool
	Name string
	Value string

	Type string `json:"-"`
	IsOnSMS bool `json:"-"`
	Normal int `json:"-"`
	Value0 string  `json:"-"`
	Value1 string  `json:"-"`

	SMS string  `json:"-"`
}

type IntConfNode struct{
	System string
	Module string
	ModuleId int
	IsOnline bool
	IsNormal bool
	Name string
	Value string

	Type string `json:"-"`
	IsOnSMS bool `json:"-"`
	Max int `json:"-"`
	Min int `json:"-"`
	
	SMS string `json:"-"`
}


type Folat32ConfNode struct{
	System string
	Module string
	ModuleId int
	IsOnline bool
	IsNormal bool
	Name string
	Value string

	Type string `json:"-"`
	IsOnSMS bool `json:"-"`
	Max float32 `json:"-"`
	Min float32 `json:"-"`
	
	SMS string `json:"-"`
}

type CommonConfNode struct{
	System string
	Module string
	ModuleId int
	IsOnline bool
	IsNormal bool
	Name string
	Value string

	Type string `json:"-"`
	IsOnSMS bool `json:"-"`

	Min1 float32 `json:"-"`
	Max1 float32 `json:"-"`
	Judge1 bool `json:"-"`
	Value1 string `json:"-"`

	Min2 float32 `json:"-"`
	Max2 float32 `json:"-"`
	Judge2 bool `json:"-"`
	Value2 string `json:"-"`

	Min3 float32 `json:"-"`
	Max3 float32 `json:"-"`
	Judge3 bool `json:"-"`
	Value3 string `json:"-"`

	Min4 float32 `json:"-"`
	Max4 float32 `json:"-"`
	Judge4 bool `json:"-"`
	Value4 string `json:"-"`

	Min5 float32 `json:"-"`
	Max5 float32 `json:"-"`
	Judge5 bool `json:"-"`
	Value5 string `json:"-"`

	Min6 float32 `json:"-"`
	Max6 float32 `json:"-"`
	Judge6 bool `json:"-"`
	Value6 string `json:"-"`

	SMS string `json:"-"`
}

func (p *BoolConfNode)CountPhysicalNode(intstring string){
	if !p.IsOnline{
		p.IsNormal =true
		p.Value = "**"
		return
	}
		
	i, err := strconv.Atoi(intstring)
	if(strings.Compare(intstring,"timeout")==0||strings.Compare(intstring,"undefined")==0||err !=nil){
		return
	}

	if (p.Normal==i){
		p.IsNormal =true
	}
	
	if(i==0){
		p.Value=p.Value0
		return
	}else if(i==1)){
		p.Value=p.Value1
		return
	}else{
		return
	}

}

func (p *IntConfNode)CountPhysicalNode(intstring string){
	if !p.IsOnline{
		p.IsNormal =true
		p.Value = "**"
		return
	}

	i,err := strconv.Atoi(intstring)
	if(strings.Compare(intstring,"timeout")==0||strings.Compare(intstring,"undefined")==0||err !=nil){
		return
	}

	if (p.Min !=0&&p.Max !=0&&p.Min<=i<=p.Max){
		p.Value =intstring
		p.IsNormal =true
	}

	return

}

func (p *Float32ConfNode)CountPhysicalNode(float32string string){
	if !p.IsOnline{
		p.IsNormal =true
		p.Value = "**"
		return
	}

	fl,err := strconv.ParseFloat(floatstring,32)
	if(strings.Compare(floatstring,"timeout")==0||strings.Compare(floatstring,"undefined")==0||err !=nil){
		return
	}

	if (p.Min !=0&&p.Max !=0&&p.Min<=fl<=p.Max){
		p.Value =float32string
		p.IsNormal =true
	}

	return
}

func (p *CommonConfNode)CountPhysicalNode(floatstring string){
	if !p.IsOnline{
		p.IsNormal =true
		p.Value = "**"
		return
	}

	fl,err := strconv.ParseFloat(floatstring,32)
	if(strings.Compare(floatstring,"timeout")==0||strings.Compare(floatstring,"undefined")==0||err !=nil){
		return
	}

	if (p.Min1 !=0&&p.Max1 !=0&&p.Min1<=fl<=p.Max1)){
		p.IsNormal =p.Judge1//这里也需要从conf里读取配置（judge字段）
		if strings.Compare(p.Value1,"self") !=0{
			p.Value =p.Value1
			return
		}else{
			p.Value =float32string
			return
		}
	}

	if (p.Min2 !=0&&p.Max2 !=0&&p.Min2<=fl<=p.Max2){
		p.IsNormal =p.Judge2//这里也需要从conf里读取配置（judge字段）
		if strings.Compare(p.Value2,"self") !=0{
			p.Value =p.Value2
			return
		}else{
			p.Value =float32string
			return
		}
	}

	if (p.Min3 !=0&&p.Max3 !=0&&p.Min3<=fl<=p.Max3){
		p.IsNormal =p.Judge3//这里也需要从conf里读取配置（judge字段）
		if strings.Compare(p.Value3,"self") !=0{
			p.Value =p.Value3
			return
		}else{
			p.Value =float32string
			return
		}
	}

	if (p.Min4 !=0&&p.Max4 !=0){
		if(p.Min4<=fl<=p.Max4){
			p.IsNormal =p.Judge4//这里也需要从conf里读取配置（judge字段）
			if strings.Compare(p.Value4,"self") !=0{
				p.Value =p.Value4
				return
			}else{
				p.Value =float32string
				return
			}
		}
	}

	if (p.Min5 !=0&&p.Max5 !=0){
		if(p.Min5<=fl<=p.Max5){
			p.IsNormal =p.Judge5//这里也需要从conf里读取配置（judge字段）
			if strings.Compare(p.Value5,"self") !=0{
				p.Value =p.Value5
				return
			}else{
				p.Value =float32string
				return
			}
		}
	}

	if (p.Min6 !=0&&p.Max6 !=0){
		if(p.Min6<=fl<=p.Max6){
			p.IsNormal =p.Judge6//这里也需要从conf里读取配置（judge字段）
			if strings.Compare(p.Value6,"self") !=0{
				p.Value =p.Value6
				return
			}else{
				p.Value =float32string
				return
			}
		}
	}
	
	return
}

func (p *BoolenConfNode )GetSystemAndModuleString()(system string,module string){
	return p.System, p.Module
}

func (p *IntConfNode )GetSystemAndModuleString()(system string,module string){
	return p.System, p.Module
}

func (p *Folat32ConfNode)GetSystemAndModuleString()(system string,module string){
	return p.System, p.Module
}

func (p *CommonConfNode)GetSystemAndModuleString()(system string,module string){
	return p.System, p.Module
}



func (p *BoolenConfNode )GetJson()data []byte{
	data, err := json.Marshal(p);err ==nil{
		return
	}

	return
}

func (p *IntConfNode)GetJson()data []byte{
	data, err := json.Marshal(p);err ==nil{
		return
	}

	return
}

func (p *Folat32ConfNode)GetJson()data []byte{
	data, err := json.Marshal(p);err ==nil{
		return
	}

	return
}

func (p *CommonConfNode)GetJson()data []byte{
	data, err := json.Marshal(p);err ==nil{
		return
	}

	return
}


func (p *BoolenConfNode)GetSystemModuleAndCountJSON(intstring string)(system string, module string, json []byte){
	p.CountPhysicalNode(intstring)
	system, module =p.GetSystemModuleString()
	json =p.GetJson()
	return
}

func (p *IntConfNode)GetSystemModuleAndCountJSON(intstring string)(system string, module string, json []byte){
	p.CountPhysicalNode(intstring)
	system, module =p.GetSystemModuleString()
	json =p.GetJson()
	return
}

func (p *Folat32ConfNode)GetSystemModuleAndCountJSON(float32string string)(system string, module string, json []byte){
	p.CountPhysicalNode(float32string)
	system, module =p.GetSystemModuleString()
	json =p.GetJson()
	return
}

func (p *CommonConfNode)GetSystemModuleAndCountJSON(float32string string)(system string, module string, json []byte){
	p.CountPhysicalNode(floatstring)
	system, module =p.GetSystemModuleString()(
	json =p.GetJson()
	return
}
















