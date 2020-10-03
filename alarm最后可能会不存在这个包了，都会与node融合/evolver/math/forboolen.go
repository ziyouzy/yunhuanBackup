//Content英译为"内容"而非Context英译为"环境，上下文，来龙去脉"
package math

import(
	"strings"
)
func GetContentAndJudgeOnline(conf2 string, conf4 string,v string)(string,bool){
	//0:开,1:关,default:N/A,offline:**
	tempSl := strings.Split(conf2,",")
	if strings.Index(conf4,"offline")==0{
		return strings.Split(tempSl[3],":")[1], false
	}

	for _,value :=range tempSl{
		if(strings.Index(value,v)==0){
			return strings.Split(value,":")[1],true
		}
	}

	if strings.Index(tempSl[2],"default")==0{
		return strings.Split(tempSl[2],":")[1],true
	}else{
		return "conf err",true
	}
}

func JudgeSafe(conf3 string,v string) bool{
	//0,timeout,null,nil,err,error
	tempSl :=strings.Split(conf3,",")
	for _,value := range tempSl{
		if strings.Index(value,v)==0{
			return false
		}
	}
	return true
}

func JudgeNeedSms(conf5 string) bool{
	//onsms
	if strings.Index(conf5,"onsms")==0{
		return true
	}else{
		return false	
	}
}