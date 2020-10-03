//不需要依赖注入全局变量句柄(指针)之类的东西
//因为"github.com/joho/godotenv"在main初始化时就已经把所有环境变量持久化到内存了
//他的使用方式是在程序最外层实例化并状态保持
//虽然evolver包属于pysicalnode的子包，但是使用方式上如同health.go所示
//与创建dbentity、创建physicalnode、平级别调用
//再次说明，他的持久化状态与环境变量，gorm.DB一致
//相反的dbentity和pysicalnode都会在实现需求后销毁，他则不同

package evolver

import(
	"strings"
	"errors"
	"fmt"
	"os"

	"github.com/ziyouzy/mylib/yunhuanfactory/physicalnode/evolver/math"
)

type PhysicalNode struct{

}

//nodeType:诸如"DI1"、"DI2"、"DI4"、"DI7"、"DO2"、"KTWD"、"UPSU"、"UPSI"、"ZNDBU"(暂定这些)
//value从physicalnode各个结构体内的字段值直接获取
//其实使用方式是提取后再次赋值给同一字段
func (p *PhysicalNode)Evolve(nodeType string, pValue *string) error{
	if os.Getenv(nodeType) ==""{
		return errors.New(fmt.Sprintf("os.Getenv(%s) is nil",nodeType))
	}

	//环境变量在程序初始化时已经引入了对应的配置文件
	//confSl的本质就是配置文件
	confSl := strings.Split(os.Getenv(nodeType),"_")
	switch (confSl[0]){
		case "Boolen":
		*pValue = p.evolveBoolenStr(*pValue, confSl...)
		return nil
		//fmt.Println("*pValue:",*pValue)
		// case "Float":
		// 	return this.evolveFloatStr(value, confSl...)
		// case "Int":
		// 	return this.evolveIntStr(value, confSl...)
		default:
			return errors.New(fmt.Sprintf("从配置文档所读取的[%s]是未知的配置类型",confSl[0]))
	}
	
	//return errors.New(fmt.Sprintf("Evolve过程中从配置文档读取数据发生错误"))
}

//这是个私有方法，对具体类型的断言（boolen、float或str）存在于其所在的上层方法
func (p *PhysicalNode)evolveBoolenStr(v string,conf ...string)string{
	//配置文件举例：
	/*export Di1="Boolen_前门_0:开,1:关,default:N/A,offline:**_0,timeout,null,nil,err,error_online_onsms_[%T]前门被打开"*/
	title :=conf[1]

	content, isOnline :=math.GetContentAndJudgeOnline(conf[2],conf[4],v)
	isSafe :=math.JudgeSafe(conf[3],v)
	isNeedSms :=math.JudgeNeedSms(conf[5])

	var sms string
	if len(conf) ==7&&strings.Index(conf[6],"[%T]")==0{
		//[%T]前门被打开
		sms =conf[6]
	}
	
	return fmt.Sprintf("%s_%s_%t_%t_%t_%s",title,content,isSafe,isOnline,isNeedSms,sms)
}

// func (p *PhysicalNode)evolveFloatStr(strType string, strPtr*string){

// }

// func (p *PhysicalNode)evolveIntStr(strType string, strPtr*string){

// }