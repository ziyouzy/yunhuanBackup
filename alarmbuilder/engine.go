package alarmbuilder

import(
	"fmt"
	"github.com/ziyouzy/mylib/mysql"
	"github.com/ziyouzy/mylib/nodedo"
)
 
func NewEngine(base map[string]interface{})(engine *Engine, smstickerlimitmillisecond int, mysqltickerlimitmillisecond int){
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
		}
	}


	if smstickerlimitfloat,ok :=base["smssleepmillisecond"].(float64);!ok{
		fmt.Println("从json初始化smssleepmillisecond进行断言时失败，因此将会把smssleepmillisecond的值设置为4*60")
		smstickerlimitmillisecond =240*60*1000
	}else{
		fmt.Println("从json初始化smssleepmillisecond进行断言成功，smstimerlimitmillisecond=", int(smstickerlimitfloat))
		smstickerlimitmillisecond =int(smstickerlimitfloat)
	}

	if mysqltickerlimitfloat,ok :=base["mysqlsleepmillisecond"].(float64);!ok{
		fmt.Println("从json初始化mysqltickerlimitmillisecond进行断言时失败，因此将会把mysqltickerlimitmillisecond的值设置为4*60")
		mysqltickerlimitmillisecond =240*60*1000
	}else{
		fmt.Println("从json初始化mysqltickerlimitmillisecond进行断言成功，mysqltickerlimitmillisecond=", int(mysqltickerlimitfloat))
		mysqltickerlimitmillisecond =int(mysqltickerlimitfloat)
	}

	return
}



//和他的上层一样，都是会常驻于内存中
type Engine struct{
	//"sAT+SMSEND=861391000000,XX您好,贵公司的%s\n"
	e []string
}


func (p *Engine)JudgeOneNodeDo(nd nodedo.NodeDo) (string, bool, []string, *mysql.Alarm){
	amString := nd.PrepareSMSAlarm()
	if amString ==""{ return "", true, nil, nil }

	key := nd.GetKey()

	var sms []string
	for _, v := range p.e{
		sms =append(sms,fmt.Sprintf(v,amString))
	}
	
	alarm :=mysql.Alarm{}
	nd.PrepareMYSQLAlarm(&alarm)


	return key, false, sms, &alarm
}

