//设计思路会与node-orm基本一致
//首先，从viper拿到的会是一个map[string]interface{}
//一个传感器一次的报警会对应一个confSMS的实体
//或者说一个physicalnode的任何一个被观测的节点如果异常都会对应一个confsms实体
//或者说每个confnode的异常都会对应一个confsms实体
package conf

import(
	"fmt"
)

type confAlarm struct{
	SMS []string
	SMSSleepMin int

	MySQLMsg string
	MySQLSleepMin int
}

func NewConfAlram(cn ConfNode) *confAlarm{
	var confalarm =confAlarm{}
	alarmString := cn.JudgeAlarm()
	if (alarmString !=""){
		if ticket1, ok1 :=confAlarmMap["smssleepmin"].(int);ok1{
			if ticket2, ok2 :=confAlarmMap["mysqlsleepmin"].(int);ok2{
				confalarm.SMSSleepMin = ticket1 
				confalarm.MySQLSleepMin= ticket2
			}else{
				//记录日志：关于mysqlsleepmin的json配置文件似乎些错了
			}
		}else{
			//记录日志：关于smssleepmin的json配置文件似乎些错了
		}

		smsserialize, ok :=confAlarmMap["smsserialize"].(string)
		if !ok{
			//记录日志：关于smsserialize的json配置文件似乎些错了
		}
			
		msgs :=confAlarmMap["smstel"]
		smsmsgsmap, ok :=msgs.(map[string]string)
		if !ok{
				//记录日志：关于smstel的json配置文件似乎些错了
		}else{
			for k, v := range smsmsgsmap{
				/*sAT+SMSEND=86%s,%s您好,贵公司%s\n*/
				confalarm.SMS =append(confalarm.SMS,fmt.Sprintf(smsserialize,v, k,alarmString))
			}
		}

		//confSMS.mysqlmsg =	
		return &confalarm
	}else{
		return nil
	}
}


// func (p *confAlarm)UpdateAlarm(s string){
// 	for k, v := range p.sms{
// 		sms[k] =fmt.Sprintf("%s%s",v , s)
// 	}
// }

// func (p *confAlarm)AllSMS() [][]byte{
// 	smsbytes := [][]byte
// 	for v :=sms{
// 		append(smsbytes,[]byte(v))
// 	}
// 	return smsbytes 
// }