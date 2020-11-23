package alarmcontroller

import(
	"fmt"
	"github.com/ziyouzy/mylib/model"
	"github.com/ziyouzy/mylib/nodedo"
)
 
func NewEngine(base map[string]interface{})(engine *Engine, smstickerlimitmin float64, mysqltickerlimitmin float64){
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
			//engine.smsArr =make([]string,len(engine.e))
			//engine.alarmDBEntity =new(model.AlarmEntity)
		}
	}


	if smstickerlimitmin,ok =base["smssleepmin"].(float64);!ok{
		fmt.Println("从json初始化smssleepmin进行断言时失败，因此将会把smssleepmin的值设置为4*60")
		smstickerlimitmin =240
	}else{
		fmt.Println("从json初始化smssleepmin进行断言成功，smstimerlimitmin=", smstickerlimitmin)
	}

	if mysqltickerlimitmin,ok =base["mysqlsleepmin"].(float64);!ok{
		fmt.Println("从json初始化mysqltickerlimitmin进行断言时失败，因此将会把mysqltickerlimitmin的值设置为4*60")
		mysqltickerlimitmin =240
	}else{
		fmt.Println("从json初始化mysqltickerlimitmin进行断言成功，mysqltickerlimitmin=", mysqltickerlimitmin)
	}

	return
}



//type Engine []string
//和他的上层一样，都是会常驻于内存中
type Engine struct{
	//"sAT+SMSEND=861391000000,孙子您好,贵公司的%s\n"
	e []string

	//isSafe bool
	//smsArr []string
	//alarmDBEntity *model.AlarmEntity
}


func (p *Engine)JudgeOneNodeDo(nd nodedo.NodeDo) (bool,[]string,*model.AlarmEntity){
	fmt.Println("JudgeOneNodeDo0")
	amString := nd.PrepareSMSAlarm()
	fmt.Println("JudgeOneNodeDo1.amString:",amString)
	if amString ==""{
		return true,nil,nil
	}
	fmt.Println("JudgeOneNodeDo2")
	var sms []string
	for _, v := range p.e{
		sms =append(sms,fmt.Sprintf(v,amString))
	}
	fmt.Println("JudgeOneNodeDo3")
	ae :=model.AlarmEntity{}
	//ae := *model.AlarmEntity
	//ae :=new(model.AlarmEntity)
	// fmt.Println("JudgeOneNodeDo4")
	nd.PrepareMYSQLAlarm(&ae)
	fmt.Println("JudgeOneNodeDo5")
	return false,sms,&ae
}

