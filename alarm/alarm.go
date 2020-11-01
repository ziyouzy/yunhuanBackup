//设计思路会与node-orm基本一致
//首先，从viper拿到的会是一个map[string]interface{}
//一个传感器一次的报警会对应一个confSMS的实体
//或者说一个physicalnode的任何一个被观测的节点如果异常都会对应一个confsms实体
//或者说每个confnode的异常都会对应一个confsms实体

//20201031有待完善该包的归属问题
package alarm

import(
	"fmt"
	//"strings"
	//"reflect"
)

type Alarm struct{
	SMS []string
	SMSSleepMin float64

	MySQLNameString string
	MySQLValueString string
	MySQLUnitString string
	MySQLContentString string
	MySQLSleepMin float64
}

func NewAlram(cn ConfNode) *confAlarm{
	var confalarm =confAlarm{}
	alarmString := cn.JudgeAlarm()
	if (alarmString !=""){
		if ticket1, ok1 :=confAlarmMap["smssleepmin"].(float64);ok1{
			if ticket2, ok2 :=confAlarmMap["mysqlsleepmin"].(float64);ok2{
				confalarm.SMSSleepMin = ticket1 
				confalarm.MySQLSleepMin= ticket2
			}else{
				fmt.Println("ok2:",ok2)
				//记录日志：关于mysqlsleepmin的json配置文件似乎些错了
			}
		}else{
			//记录日志：关于smssleepmin的json配置文件似乎些错了
			fmt.Println("ok1:",ok1)
		}

		smsserialize, ok :=confAlarmMap["smsserialize"].(string)
		if !ok{
			//记录日志：关于smsserialize的json配置文件似乎些错了
			fmt.Println("ok3:",ok)
		}
			
		msgs :=confAlarmMap["smstel"]
		smsmsgsmap, ok :=msgs.(map[string]interface{})
		if !ok{
				//记录日志：关于smstel的json配置文件似乎些错了
				fmt.Println("ok4:",ok)
		}else{
			for k, v := range smsmsgsmap{
				/*sAT+SMSEND=86%s,%s您好,贵公司%s\n*/
				confalarm.SMS =append(confalarm.SMS,fmt.Sprintf(smsserialize,v, k,alarmString))
			}
		}

		confalarm.MySQLNameString, confalarm.MySQLValueString, confalarm.MySQLUnitString, confalarm.MySQLContentString =cn.PrepareMYSQLAlarm()
		
		return &confalarm
	}else{
		return nil
	}
}

