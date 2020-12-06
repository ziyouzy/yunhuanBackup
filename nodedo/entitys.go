package nodedo

import(
	"strconv"
	"strings"
	"fmt"
	"encoding/json"

	"github.com/ziyouzy/mylib/mysql"
)

type BoolenNodeDo struct{
	SouthId int //从json配置文档获取
	SouthBound string //从json配置文档获取
	Name string //从json配置文档获取
	Type string `json:"-"`//从json配置文档获取
	//单位
	Unit string //从json配置文档获取
	//无论从哪个类型的physicalnode拿到的RawStr一定会是string
	RawStr string//由UpdateOneNodeDo方法实时更新

	IsOnline bool //从json配置文档获取
	IsOnSMS bool `json:"-"`//从json配置文档获取

	IsTimeOut bool
	TimeOutSec int //从json配置文档获取
	TimeUnixNano uint64 //由UpdateOneNodeDo方法实时更新

	IsNormal bool //由UpdateOneNodeDo方法实时更新
	//Condition:条件-该字段描述正常所需满足的条件值 
	ConditionValue bool `json:"-"`//从json配置文档获取
	//UpdateOneNodeDo方法内会基于后端正异常逻辑所判断出的结果，生成对应的值，前端拿到后也可以忽略这个字段，而是结合其他字段去设计逻辑，自定义显示的值
	FrontEndStr string//由UpdateOneNodeDo方法实时更新

	SMS string  `json:"-"`//从json配置文档获取
}


type IntNodeDo struct{
	SouthId int //从json配置文档获取
	SouthBound string //从json配置文档获取
	Name string //从json配置文档获取
	Type string `json:"-"` //从json配置文档获取
	//单位
	Unit string //从json配置文档获取
	//无论从哪个类型的physicalnode拿到的RawStr一定会是string
	RawStr string //由UpdateOneNodeDo方法实时更新

	IsOnline bool //从json配置文档获取
	IsOnSMS bool `json:"-"`//从json配置文档获取

	IsTimeOut bool
	TimeOutSec int //从json配置文档获取
	TimeUnixNano uint64 //由UpdateOneNodeDo方法实时更新

	IsNormal bool //由UpdateOneNodeDo方法实时更新
	Max int `json:"-"` //从json配置文档获取
	Min int `json:"-"` //从json配置文档获取
	//UpdateOneNodeDo方法内会基于后端正异常逻辑所判断出的结果，生成对应的值，前端拿到后也可以忽略这个字段，而是结合其他字段去设计逻辑，自定义显示的值
	FrontEndStr string //由UpdateOneNodeDo方法实时更新

	SMS string `json:"-"`//从json配置文档获取
}


type FloatNodeDo struct{
	SouthId int //从json配置文档获取
	SouthBound string //从json配置文档获取
	Name string //从json配置文档获取
	Type string `json:"-"`//从json配置文档获取
	//单位
	Unit string //从json配置文档获取
	//无论从哪个类型的physicalnode拿到的RawStr一定会是string
	RawStr string //由UpdateOneNodeDo方法实时更新

	IsOnline bool //从json配置文档获取
	IsOnSMS bool `json:"-"`//从json配置文档获取

	IsTimeOut bool //由UpdateOneNodeDo或TimeOut方法实时更新
	TimeOutSec int //从json配置文档获取
	TimeUnixNano uint64 //由UpdateOneNodeDo方法实时更新

	IsNormal bool //由UpdateOneNodeDo方法实时更新
	Max float64 `json:"-"`//从json配置文档获取
	Min float64 `json:"-"` //从json配置文档获取
	//UpdateOneNodeDo方法内会基于后端正异常逻辑所判断出的结果，生成对应的值，前端拿到后也可以忽略这个字段，而是结合其他字段去设计逻辑，自定义显示的值
	FrontEndStr string //由UpdateOneNodeDo方法实时更新
	
	SMS string `json:"-"`//从json配置文档获取
}


type CommonNodeDo struct{
	SouthId int //从json配置文档获取
	SouthBound string //从json配置文档获取
	Name string //从json配置文档获取
	Type string `json:"-"`//从json配置文档获取
	//单位
	Unit string //从json配置文档获取
	//无论从哪个类型的physicalnode拿到的RawStr一定会是string
	RawStr string //由UpdateOneNodeDo方法实时更新

	IsOnline bool //从json配置文档获取
	IsOnSMS bool `json:"-"`//从json配置文档获取

	IsTimeOut bool
	TimeOutSec int //从json配置文档获取
	TimeUnixNano uint64  //由UpdateOneNodeDo方法实时更新

	IsNormal1 bool `json:"-"` //由UpdateOneNodeDo方法实时更新
	Min1 float64 `json:"-"` //从json配置文档获取
	Max1 float64 `json:"-"` //从json配置文档获取
	IsNormal2 bool `json:"-"` //由UpdateOneNodeDo方法实时更新
	Min2 float64 `json:"-"` //从json配置文档获取
	Max2 float64 `json:"-"`//从json配置文档获取
	IsNormal3 bool `json:"-"` //由UpdateOneNodeDo方法实时更新
	Min3 float64 `json:"-"` //从json配置文档获取
	Max3 float64 `json:"-"` //从json配置文档获取
	IsNormal4 bool `json:"-"` //由UpdateOneNodeDo方法实时更新
	Min4 float64 `json:"-"` //从json配置文档获取
	Max4 float64 `json:"-"` //从json配置文档获取
	IsNormal5 bool `json:"-"` //由UpdateOneNodeDo方法实时更新
	Min5 float64 `json:"-"` //从json配置文档获取
	Max5 float64 `json:"-"` //从json配置文档获取
	IsNormal6 bool `json:"-"` //由UpdateOneNodeDo方法实时更新
	Min6 float64 `json:"-"` //从json配置文档获取
	Max6 float64 `json:"-"`//从json配置文档获取
	//UpdateOneNodeDo方法内会基于后端正异常逻辑所判断出的结果，生成对应的值，前端拿到后也可以忽略这个字段，而是结合其他字段去设计逻辑，自定义显示的值
	FrontEndStr string //由UpdateOneNodeDo方法实时更新89

	SMS string `json:"-"`//从json配置文档获取
}

func (p *BoolenNodeDo)UpdateOneNodeDo(value string,time uint64){
	p.IsTimeOut =false;        if (time - p.TimeUnixNano>p.TimeOutSec){p.IsTimeOut =true}
	p.TimeUnixNano =time//在这里设计判定是否过期的逻辑，也要用到TimeOutSec字段，同时需要将字段内容转化为时间戳
	
	
	i, err := strconv.Atoi(value)
	//p.IsTimeOut和从physicalNode所返回的"timeout"字符串是有区别的，后者是更为底层的问题，是用来兼容就系统而特殊设计的
	if(err !=nil || strings.Compare(value,"timeout")==0||strings.Compare(value,"undefined")==0||!p.IsOnline){
		p.IsNormal =true; p.RawStr = "**"; p.FrontEndStr  ="**"
		return
	}

	p.IsNormal =false; p.RawStr =value;p.FrontEndStr ="异常"
	if (i ==0&&p.ConditionValue==false){p.FrontEndStr ="正常"; p.IsNormal =true}
	if (i ==1&&p.ConditionValue==true){p.FrontEndStr ="正常"; p.IsNormal =true}
	return
}

func (p *IntNodeDo)UpdateOneNodeDo(value string, time string){
	p.IsTimeOut =false; p.DateTime =time
	i,err := strconv.Atoi(value)
	//p.IsTimeOut和从physicalNode所返回的"timeout"字符串是有区别的，后者是更为底层的问题，是用来兼容就系统而特殊设计的
	if(err !=nil || strings.Compare(value,"timeout")==0||strings.Compare(value,"undefined")==0||!p.IsOnline){
		p.IsNormal =true; p.RawStr = "**"; p.FrontEndStr ="**"
		return
	}

	p.IsNormal =false; p.RawStr =value;p.FrontEndStr =value
	if (p.Min<=i&&i<=p.Max){ p.IsNormal =true}
	return
}

func (p *FloatNodeDo)UpdateOneNodeDo(value string, time string){
	p.IsTimeOut =false; p.DateTime =time
	fl,err := strconv.ParseFloat(value, 64)
	//p.IsTimeOut和从physicalNode所返回的"timeout"字符串是有区别的，后者是更为底层的问题，是用来兼容就系统而特殊设计的
	if (err !=nil || strings.Compare(value,"timeout")==0||strings.Compare(value,"undefined")==0||!p.IsOnline){
		p.IsNormal =true; p.RawStr = "**"; p.FrontEndStr = "**"
		return
	}

	p.IsNormal =false; p.RawStr =value;p.FrontEndStr =value
	if (p.Min<=fl&&fl<=p.Max){p.IsNormal =true}
	return
}

func (p *CommonNodeDo)UpdateOneNodeDo(value string, time string){
	//value一定会是个可以转化为float64的字符串，同时也一定不会是个汉字或者字母的字符串
	//设计CommonNodeDo的目的是用来处理如下情况:
	/*
		某设备：
		1--正常关闭
		2--异常关闭
		3--供电正常
		4--供电异常
		1和3为IsNormail
	*/
	//于是就需要跳跃性的比较某个float64类型的大小值了
	p.IsTimeOut =false; p.DateTime =time
	fl,err := strconv.ParseFloat(value,64)
	//p.IsTimeOut和从physicalNode所返回的"timeout"字符串是有区别的，后者是更为底层的问题，是用来兼容就系统而特殊设计的
	if(err !=nil || strings.Compare(value,"timeout")==0||strings.Compare(value,"undefined")==0||!p.IsOnline){
		p.IsNormal1 =true; p.IsNormal2 =true; p.IsNormal3 =true; p.IsNormal4 =true; p.IsNormal5 =true; p.IsNormal6 =true; p.RawStr = "**"; p.FrontEndStr = "**"
		return
	}

	p.IsNormal1 =false; p.IsNormal2 =false; p.IsNormal3 =false; p.IsNormal4 =false; p.IsNormal5 =false; p.IsNormal6 =false; p.RawStr =value;p.FrontEndStr ="异常"
	if (p.Min1<=fl&&fl<=p.Max1){p.IsNormal1=true}
	if (p.Min2<=fl&&fl<=p.Max2){p.IsNormal2=true}
	if (p.Min3<=fl&&fl<=p.Max3){p.IsNormal3=true}
	if (p.Min4<=fl&&fl<=p.Max4){p.IsNormal4=true}
	if (p.Min5<=fl&&fl<=p.Max5){p.IsNormal5=true}
	if (p.Min6<=fl&&fl<=p.Max6){p.IsNormal6=true}
	if p.IsNormal1&&p.IsNormal2&&p.IsNormal3&&p.IsNormal4&&p.IsNormal5&&p.IsNormal6{
		p.FrontEndStr ="正常"
	}
	return
}

//无法通过接口直接拿到内部字段，只能设计独立的方法
func (p *BoolenNodeDo)GetTimeOutSec()int {return p.TimeOutSec}
func (p *IntNodeDo)GetTimeOutSec()int {return p.TimeOutSec}
func (p *FloatNodeDo)GetTimeOutSec()int {return p.TimeOutSec}
func (p *CommonNodeDo)GetTimeOutSec()int {return p.TimeOutSec}

func (p *BoolenNodeDo)UpdateOneNodeDoAndGetTimeOutSec(value string, time string) int{p.UpdateOneNodeDo(value,time); return p.GetTimeOutSec()}
func (p *IntNodeDo)UpdateOneNodeDoAndGetTimeOutSec(value string, time string) int{p.UpdateOneNodeDo(value,time);return p.GetTimeOutSec()}
func (p *FloatNodeDo)UpdateOneNodeDoAndGetTimeOutSec(value string, time string) int{p.UpdateOneNodeDo(value,time);return p.GetTimeOutSec()}
func (p *CommonNodeDo)UpdateOneNodeDoAndGetTimeOutSec(value string, time string) int{p.UpdateOneNodeDo(value,time);return p.GetTimeOutSec()}

//触发超时的信号会在nodedobuilder发送
func (p *BoolenNodeDo)TimeOut(){p.IsTimeOut =true}
func (p *IntNodeDo)TimeOut(){p.IsTimeOut =true}
func (p *FloatNodeDo)TimeOut(){p.IsTimeOut =true}
func (p *CommonNodeDo)TimeOut(){p.IsTimeOut =true}



//如果错误，则自动返回空值
func (p *BoolenNodeDo)GetJson()[]byte{;data, _ := json.Marshal(p); return data}
func (p *IntNodeDo)GetJson()[]byte{data, _ := json.Marshal(p); return data}
func (p *FloatNodeDo)GetJson()[]byte{data, _ := json.Marshal(p); return data}
func (p *CommonNodeDo)GetJson()[]byte{data, _ := json.Marshal(p); return data}


func (p *BoolenNodeDo)PrepareSMSAlarm()string{
	if !p.IsNormal&&p.IsOnSMS{
		return fmt.Sprintf("%s[发生异常时间为%s]", p.SMS, p.DateTime)
	}else{
		return ""
	}
}

func (p *IntNodeDo)PrepareSMSAlarm()string{
	if !p.IsNormal&&p.IsOnSMS{
		return fmt.Sprintf("%s[发生异常时间为%s]", p.SMS, p.DateTime)
	}else{
		return ""
	}
}

func (p *FloatNodeDo)PrepareSMSAlarm()string{
	if !p.IsNormal&&p.IsOnSMS{
		return fmt.Sprintf("%s[发生异常时间为%s]", p.SMS, p.DateTime)
	}else{
		return ""
	}
}

func (p *CommonNodeDo)PrepareSMSAlarm()string{
	if !p.IsNormal1||!p.IsNormal2||!p.IsNormal3||!p.IsNormal4||!p.IsNormal5||!p.IsNormal6{
		if p.IsOnSMS{
			return fmt.Sprintf("%s[发生异常时间为%s]", p.SMS, p.DateTime)
		}else{
			return ""
		}
	}else{
		return ""
	}
}

func (p *BoolenNodeDo)PrepareMYSQLAlarm(ae *mysql.Alarm){
	ae.PresentSouthID =p.SouthId;        ae.PresentSouthBound =p.SouthBound
	ae.Name =p.Name;        ae.RawStr =p.RawStr;         ae.FrontEndStr =p.FrontEndStr;        ae.Unit =p.Unit
	ae.Content =fmt.Sprintf("%s[发生异常时间为%s]", p.SMS, p.DateTime)
}

func (p *IntNodeDo)PrepareMYSQLAlarm(ae *mysql.Alarm){
	ae.PresentSouthID =p.SouthId;        ae.PresentSouthBound =p.SouthBound
	ae.Name =p.Name;        ae.RawStr =p.RawStr;         ae.FrontEndStr =p.FrontEndStr;        ae.Unit =p.Unit
	ae.Content =fmt.Sprintf("%s[发生异常时间为%s]", p.SMS, p.DateTime)
}

func (p *FloatNodeDo)PrepareMYSQLAlarm(ae *mysql.Alarm){
	ae.PresentSouthID =p.SouthId;        ae.PresentSouthBound =p.SouthBound
	ae.Name =p.Name;        ae.RawStr =p.RawStr;         ae.FrontEndStr =p.FrontEndStr;        ae.Unit =p.Unit
	ae.Content =fmt.Sprintf("%s[发生异常时间为%s]", p.SMS, p.DateTime)
}

func (p *CommonNodeDo)PrepareMYSQLAlarm(ae *mysql.Alarm){
	ae.PresentSouthID =p.SouthId;        ae.PresentSouthBound =p.SouthBound
	ae.Name =p.Name;        ae.RawStr =p.RawStr;         ae.FrontEndStr =p.FrontEndStr;        ae.Unit =p.Unit
	ae.Content =fmt.Sprintf("%s[发生异常时间为%s]", p.SMS, p.DateTime)
}

