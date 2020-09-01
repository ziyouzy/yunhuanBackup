package impl

import(
	"strings"
	"fmt"
	"os"
)

type AllNodesEvolverImpl struct{
	//全局变量句柄(指针)
	//不需要，因为"github.com/joho/godotenv"在main初始化时就已经把所有环境变量持久化到内存了
}

func (this *AllNodesEvolverImpl)Evolver(strType string, value string)string{
	if os.Getenv(strType) ==""{
		return fmt.Sprintf("os.Getenv(%s) is nil",strType)
	}
	confSl := strings.Split(os.Getenv(strType),"_")
	switch (confSl[0]){
		case "Boolen":
			return this.evolveBoolenStr(value, confSl...)
		// case "Float":
		// 	return this.evolveFloatStr(value, confSl...)
		// case "Int":
		// 	return this.evolveIntStr(value, confSl...)
		default:
			return ""
	}  
}

func getContentAndJudgeOnline(conf2 string, conf4 string,v string)(string,bool){
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

func judgeSafe(conf3 string,v string) bool{
	//0,timeout,null,nil,err,error
	tempSl :=strings.Split(conf3,",")
	for _,value := range tempSl{
		if strings.Index(value,v)==0{
			return false
		}
	}
	return true
}

func judgeNeedSms(conf5 string) bool{
	//onsms
	if strings.Index(conf5,"onsms")==0{
		return true
	}else{
		return false	
	}
}

func (this *AllNodesEvolverImpl)evolveBoolenStr(v string,conf ...string)string{
	/*export Di1="Boolen_前门_0:开,1:关,default:N/A,offline:**_0,timeout,null,nil,err,error_online_onsms_[%T]前门被打开"*/
	title :=conf[1]

	content, isOnline :=getContentAndJudgeOnline(conf[2],conf[4],v)
	isSafe :=judgeSafe(conf[3],v)
	isNeedSms :=judgeNeedSms(conf[5])

	var sms string
	if len(conf) ==7&&strings.Index(conf[6],"[%T]")==0{
		//[%T]前门被打开
		sms =conf[6]
	}
	
	return fmt.Sprintf("%s_%s_%t_%t_%t_%s",title,content,isSafe,isOnline,isNeedSms,sms)
}

// func (this *AllNodesEvolverImpl)evolveFloatStr(strType string, strPtr*string){

// }

// func (this *AllNodesEvolverImpl)evolveIntStr(strType string, strPtr*string){

// }