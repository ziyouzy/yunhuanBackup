package alarmcontroller

import(
	"fmt"
	"github.com/ziyouzy/mylib/model"
	"github.com/ziyouzy/mylib/nodedo"
)
 
func NewEngine(base map[string]interface{})(engine *Engine, smstimerlimitmin float64, mysqltimerlimitmin float64){
	engine =new(Engine)
	smsserialize, ok :=base["smsserialize"].(string)
	if !ok{
		fmt.Println("从json初始化smsserialize进行断言时失败")
	}else{
		msgs :=base["smstel"]//断言才需要if，不断言不需要
		if smsmsgsmap, ok :=msgs.(map[string]interface{});!ok{
			fmt.Println("从json初始化smstel进行断言时失败")
		}else{
			for k, v := range smsmsgsmap{
				if name,ok :=v.(string);ok{
					/*sAT+SMSEND=86%s,%s您好,贵公司%s\n*/
					fmt.Println(fmt.Sprintf(smsserialize,name, k,"%s"))
					engine.e =append(engine.e,fmt.Sprintf(smsserialize,name, k,"%s"))
				}
			}
			engine.smsArr =make([]string,len(engine.e))
			engine.alarmDBEntity =new(model.AlarmEntity)
		}
	}


	if smstimerlimitmin,ok =base["smssleepmin"].(float64);!ok{
		fmt.Println("从json初始化smssleepmin进行断言时失败，因此将会把smssleepmin的值设置为4*60")
		smstimerlimitmin =240
	}else{
		fmt.Println("从json初始化smssleepmin进行断言成功，smstimerlimitmin=", smstimerlimitmin)
	}

	if mysqltimerlimitmin,ok =base["mysqlsleepmin"].(float64);!ok{
		fmt.Println("从json初始化mysqltimerlimitmin进行断言时失败，因此将会把mysqltimerlimitmin的值设置为4*60")
		mysqltimerlimitmin =240
	}else{
		fmt.Println("从json初始化mysqltimerlimitmin进行断言成功，mysqltimerlimitmin=", mysqltimerlimitmin)
	}

	return
}



//type Engine []string
//和他的上层一样，都是会常驻于内存中
type Engine struct{
	//"sAT+SMSEND=861391000000,孙子您好,贵公司的%s\n"
	e []string

	isSafe bool
	smsArr []string
	alarmDBEntity *model.AlarmEntity
}


func (p *Engine)JudgeOneNodeDo(nd nodedo.NodeDo) (bool,[]string,*model.AlarmEntity){
	amString := nd.PrepareSMSAlarm()
	if amString ==""{
		p.isSafe =true
		return p.isSafe,nil,nil
	}

	p.isSafe =false
	for k, v := range p.e{
		p.smsArr[k] =fmt.Sprintf(v,amString)
	}
	
	nd.PrepareMYSQLAlarm(p.alarmDBEntity)

	return p.isSafe,p.smsArr,p.alarmDBEntity
}

