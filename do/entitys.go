package do

import(
	"strconv"
	"strings"
	"fmt"
	"encoding/json"

	"github.com/ziyouzy/mylib/model"
)

type BoolenNodeDo struct{
	Id int
	ParentModuleId int 

	IsOnline bool 
	IsNormal bool 
	IsTimeOut bool
	Name string 
	Value string
	Unit string 

	Type string `json:"-"`
	IsOnSMS bool `json:"-"`
	Normal int `json:"-"`
	Value0 string  `json:"-"`
	Value1 string  `json:"-"`

	SMS string  `json:"-"`
	Date string 
}

type IntNodeDo struct{
	Id int
	ParentModuleId int

	IsOnline bool
	IsNormal bool
	IsTimeOut bool
	Name string
	Value string
	Unit string

	Type string `json:"-"`
	IsOnSMS bool `json:"-"`
	Max int `json:"-"`
	Min int `json:"-"`
	
	SMS string `json:"-"`
	Date string
}


type FloatNodeDo struct{
	Id int
	ParentModuleId int

	IsOnline bool
	IsNormal bool
	IsTimeOut bool
	Name string
	Value string
	Unit string

	Type string `json:"-"`
	IsOnSMS bool `json:"-"`
	Max float64 `json:"-"`
	Min float64 `json:"-"`
	
	SMS string `json:"-"`
	Date string
}

type CommonNodeDo struct{
	Id int
	ParentModuleId int

	IsOnline bool
	IsNormal bool
	IsTimeOut bool
	Name string
	Value string
	Unit string

	Type string `json:"-"`
	IsOnSMS bool `json:"-"`

	Min1 float64 `json:"-"`
	Max1 float64 `json:"-"`
	Judge1 bool `json:"-"`
	Value1 string `json:"-"`

	Min2 float64 `json:"-"`
	Max2 float64 `json:"-"`
	Judge2 bool `json:"-"`
	Value2 string `json:"-"`

	Min3 float64 `json:"-"`
	Max3 float64 `json:"-"`
	Judge3 bool `json:"-"`
	Value3 string `json:"-"`

	Min4 float64 `json:"-"`
	Max4 float64 `json:"-"`
	Judge4 bool `json:"-"`
	Value4 string `json:"-"`

	Min5 float64 `json:"-"`
	Max5 float64 `json:"-"`
	Judge5 bool `json:"-"`
	Value5 string `json:"-"`

	Min6 float64 `json:"-"`
	Max6 float64 `json:"-"`
	Judge6 bool `json:"-"`
	Value6 string `json:"-"`

	SMS string `json:"-"`
	Date string 
}

func (p *BoolenNodeDo)UpdateOneNodeDo(intstring string,time string){
	p.IsTimeOut =false

	if !p.IsOnline{
		p.IsNormal =true
		p.Value = "**"
		p.Date =time
		return
	}
		
	i, err := strconv.Atoi(intstring)
	if(strings.Compare(intstring,"timeout")==0||strings.Compare(intstring,"undefined")==0||err !=nil){
		return
	}

	p.Date =time

	if p.Normal==i{
		p.IsNormal =true
		p.Value=p.Value0
		return
	}else{
		p.Value=p.Value1
		return
	}
}

func (p *IntNodeDo)UpdateOneNodeDo(intstring string, time string){
	p.IsTimeOut =false

	if !p.IsOnline{
		p.IsNormal =true
		p.Value = "**"
		p.Date =time
		return
	}

	i,err := strconv.Atoi(intstring)
	if(strings.Compare(intstring,"timeout")==0||strings.Compare(intstring,"undefined")==0||err !=nil){
		return
	}

	p.Date =time

	if (p.Min !=0&&p.Max !=0&&p.Min<=i&&i<=p.Max){
		p.Value =intstring
		p.IsNormal =true
	}

	return

}

func (p *FloatNodeDo)UpdateOneNodeDo(floatstring string, time string){
	p.IsTimeOut =false

	if !p.IsOnline{
		p.IsNormal =true
		p.Date =time
		p.Value = "**"
		return
	}

	fl,err := strconv.ParseFloat(floatstring, 64)
	if(strings.Compare(floatstring,"timeout")==0||strings.Compare(floatstring,"undefined")==0||err !=nil){
		return
	}

	p.Date =time

	if (p.Min !=0&&p.Max !=0&&p.Min<=fl&&fl<=p.Max){
		p.Value =floatstring
		p.IsNormal =true
	}

	return
}

func (p *CommonNodeDo)UpdateOneNodeDo(floatstring string, time string){
	p.IsTimeOut =false
	
	if !p.IsOnline{
		p.IsNormal =true
		p.Value = "**"
		return
	}

	fl,err := strconv.ParseFloat(floatstring,64)
	if(strings.Compare(floatstring,"timeout")==0||strings.Compare(floatstring,"undefined")==0||err !=nil){
		return
	}

	p.Date =time

	if (p.Min1 !=0&&p.Max1 !=0&&p.Min1<=fl&&fl<=p.Max1){
		p.IsNormal =p.Judge1//这里也需要从conf里读取配置（judge字段）
		if strings.Compare(p.Value1,"self") !=0{
			p.Value =p.Value1
			return
		}else{
			p.Value =floatstring
			return
		}
	}

	if (p.Min2 !=0&&p.Max2 !=0&&p.Min2<=fl&&fl<=p.Max2){
		p.IsNormal =p.Judge2//这里也需要从conf里读取配置（judge字段）
		if strings.Compare(p.Value2,"self") !=0{
			p.Value =p.Value2
			return
		}else{
			p.Value =floatstring
			return
		}
	}

	if (p.Min3 !=0&&p.Max3 !=0&&p.Min3<=fl&&fl<=p.Max3){
		p.IsNormal =p.Judge3//这里也需要从conf里读取配置（judge字段）
		if strings.Compare(p.Value3,"self") !=0{
			p.Value =p.Value3
			return
		}else{
			p.Value =floatstring
			return
		}
	}

	if (p.Min4 !=0&&p.Max4 !=0&&p.Min4<=fl&&fl<=p.Max4){
		p.IsNormal =p.Judge4//这里也需要从conf里读取配置（judge字段）
		if strings.Compare(p.Value4,"self") !=0{
			p.Value =p.Value4
			return
		}else{
			p.Value =floatstring
			return
		}
	}

	if (p.Min5 !=0&&p.Max5 !=0&&p.Min5<=fl&&fl<=p.Max5){
		p.IsNormal =p.Judge5//这里也需要从conf里读取配置（judge字段）
		if strings.Compare(p.Value5,"self") !=0{
			p.Value =p.Value5
			return
		}else{
			p.Value =floatstring
			return
		}
	}

	if (p.Min6 !=0&&p.Max6 !=0&&p.Min6<=fl&&fl<=p.Max6){
		p.IsNormal =p.Judge6//这里也需要从conf里读取配置（judge字段）
		if strings.Compare(p.Value6,"self") !=0{
			p.Value =p.Value6
			return
		}else{
			p.Value =floatstring
			return
		}
	}

	return
}



func (p *BoolenNodeDo)GetJson()[]byte{
	//如果错误，则自动返回空值
	data, _ := json.Marshal(p)
	return data
}

func (p *IntNodeDo)GetJson()[]byte{
	data, _ := json.Marshal(p)
	return data
}

func (p *FloatNodeDo)GetJson()[]byte{
	data, _ := json.Marshal(p)
	return data
}

func (p *CommonNodeDo)GetJson()[]byte{
	data, _ := json.Marshal(p)
	return data
}




func (p *BoolenNodeDo)PrepareSMSAlarm()string{
	if !p.IsNormal&&p.IsOnSMS{
		return fmt.Sprintf("%s[发生异常时间为%s]", p.SMS, p.Date)
	}else{
		return ""
	}
}

func (p *IntNodeDo)PrepareSMSAlarm()string{
	if !p.IsNormal&&p.IsOnSMS{
		return fmt.Sprintf("%s[发生异常时间为%s]", p.SMS, p.Date)
	}else{
		return ""
	}
}

func (p *FloatNodeDo)PrepareSMSAlarm()string{
	if !p.IsNormal&&p.IsOnSMS{
		return fmt.Sprintf("%s[发生异常时间为%s]", p.SMS, p.Date)
	}else{
		return ""
	}
}

func (p *CommonNodeDo)PrepareSMSAlarm()string{
	if !p.IsNormal&&p.IsOnSMS{
		return fmt.Sprintf("%s[发生异常时间为%s]", p.SMS, p.Date)
	}else{
		return ""
	}
}

func (p *BoolenNodeDo)PrepareMYSQLAlarm(ae *model.AlarmEntity){
	ae.NodeID =p.Id
	ae.ModuleParentID =p.ParentModuleId

	ae.Name =p.Name
	ae.Value =p.Value
	ae.Unit =p.Unit
	ae.Content =fmt.Sprintf("%s[发生异常时间为%s]", p.SMS, p.Date)
}

func (p *IntNodeDo)PrepareMYSQLAlarm(ae *model.AlarmEntity){
	ae.NodeID =p.Id
	ae.ModuleParentID =p.ParentModuleId

	ae.Name =p.Name
	ae.Value =p.Value
	ae.Unit =p.Unit
	ae.Content =fmt.Sprintf("%s[发生异常时间为%s]", p.SMS, p.Date)
}

func (p *FloatNodeDo)PrepareMYSQLAlarm(ae *model.AlarmEntity){
	ae.NodeID =p.Id
	ae.ModuleParentID =p.ParentModuleId

	ae.Name =p.Name
	ae.Value =p.Value
	ae.Unit =p.Unit
	ae.Content =fmt.Sprintf("%s[发生异常时间为%s]", p.SMS, p.Date)
}

func (p *CommonNodeDo)PrepareMYSQLAlarm(ae *model.AlarmEntity){
	ae.NodeID =p.Id
	ae.ModuleParentID =p.ParentModuleId

	ae.Name =p.Name
	ae.Value =p.Value
	ae.Unit =p.Unit
	ae.Content =fmt.Sprintf("%s[发生异常时间为%s]", p.SMS, p.Date)
}