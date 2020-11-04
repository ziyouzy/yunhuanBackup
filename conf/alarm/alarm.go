package alarm

import(
	"fmt"
	//"strings"
	//"reflect"
	"github.com/ziyouzy/mylib/conf"
	"github.com/ziyouzy/mylib/model"
)
 
func NewAlramTemplate(base map[string]interface{})(at *AlarmTemplate, smstimerlimitmin float64, mysqltimerlimitmin float64){
	if smstimerlimitmin ,ok :=base["smssleepmin"].(float64);ok{
		fmt.Println("从json初始化smssleepmin进行断言时失败")
	}

	if mysqltimerlimitmin ,ok =base["mysqlsleepmin"].(float64);ok{
		fmt.Println("从json初始化mysqlsleepmin进行断言时失败")
	}

	if smsserialize, ok =confAlarmMap["smsserialize"].(string);ok{
		fmt.Println("从json初始化smsserialize进行断言时失败")
	}

	msgs :=confAlarmMap["smstel"]//断言才需要if，不断言不需要
	if smsmsgsmap, ok =msgs.(map[string]interface{});ok{
		fmt.Println("从json初始化smstel进行断言时失败")
	}else{
		for k, v := range smsmsgsmap{
			/*sAT+SMSEND=86%s,%s您好,贵公司%s\n*/
			at.SMSTemplate =append(at.SMSTemplate,fmt.Sprintf(smsserialize,v, k,"%s"))
		}
	}
}

//内部元素是两个模板，Alarm和他的上层cache一样，都是属于缓存，会常驻于内存中
type AlarmTemplate struct{
	//之所以是个切片只是因为一个告警短信需要对应多个手机号
	//同时这个切片在初始化时就完成了，结构如下：
	//:"sAT+SMSEND=861391000000,孙子您好,贵公司的%s\n"
	//最后的%s等待与某一个超限的NodeDo实现对接
	SMSTemplate []string

	//和SMSTemplate不同
	//这个字段完全不需要初始化
	//当在调用CreateAlarm返回一个新的AlarmTemplate副本时
	//才需要对这个副本的AlarmEntity填充各个字段并返回
	//或者说，这个字段的初始化就是默认值，不需要再进行额外的操作
	AlarmEntityTemplate model.AlarmEntity
}


func (p *AlarmTemplate)CreateAlarm(nd NodeDo) *AlarmTemplate{
	alarm :=AlarmEntity{}
	alarmString := cn.JudgeAlarm()
	if alarmString ==""{
		return nil
	}

	for _, v := range p.SMSTemplate{
		/*sAT+SMSEND=86XXX,XXX您好,贵公司%s\n*/
		alarm.SMS =append(alarm.SMS,fmt.Sprintf(v,alarmString))
	}
	
	alarm.MySQLNameString, alarm.MySQLValueString, alarm.MySQLUnitString, alarm.MySQLContentString =cn.PrepareMYSQLAlarm()	
	
	return &alarm
}

