//这里要用到timer而不是tikcer
//出于对节省资源的考虑，这里会先判定是否超时，如超时才会去进行下一步对NodeDo的相关计算
package alarmbuilder

import(
	"time"
	"fmt"

	"github.com/ziyouzy/mylib/mysql"	
	"github.com/ziyouzy/mylib/nodedo"	
)


var builder *AlarmBuilder
type AlarmBuilder struct{
	smsTimer *time.Timer
	smsTimerLimitSec float64
	smsTimerStop chan bool
	smsAlarmIsReady bool
	SMSAlarmCh chan []byte 


	mysqlTimer *time.Timer
	mysqlTimerLimitSec float64
	mysqlTimerStop chan bool
	mysqlAlarmIsReady bool
	MYSQLAlarmCh chan *mysql.Alarm


	//这样设计是在遵顼分层的设计思路，也就是纯粹的为了分层而去采用了组合
	//基于能组合就不继承的原则，这里无论是组合还是继承都是合理的，所以既然地位相当那还是优先组合吧
	//退一步讲，就算是继承了，唯一的原因也只是为了分层思路而继承
	e *Engine
}

func LoadSingletonPattern(sourcefromviper map[string]interface{}){builder =BuildAlarmBuilder(sourcefromviper)}
//这里模仿了time包的NewTimer的设计模式，New出来的对象生命周期很可能为主函数
func BuildAlarmBuilder(sourcefromviper map[string]interface{}) *AlarmBuilder{
	builder := AlarmBuilder{}
	e, smstimerlimitmin, mysqltimerlimitmin := NewEngine(sourcefromviper)
	fmt.Println("smstimerlimitmin:",smstimerlimitmin,"mysqltimerlimitmin:",mysqltimerlimitmin)

	//实例化内部字段（内部的Engine才会真正进行监测某个NodeDo是否超限）
	builder.e =e
	builder.smsTimerLimitSec = smstimerlimitmin * 60
	builder.mysqlTimerLimitSec  = mysqltimerlimitmin * 60
	
	builder.newSMSTimer()
	builder.newMYSQLTimer()
	return &builder
}


//func NewSMSTimer(){builder.newSMSTimer()}
func (p *AlarmBuilder)newSMSTimer(){
	if p.smsTimer !=nil{p.smsTimerStop <-true;time.Sleep(time.Second)}

	p.smsTimer =time.NewTimer(time.Duration(p.smsTimerLimitSec) * time.Second)
	p.smsAlarmIsReady =true
	p.smsTimerStop =make(chan bool)

	go func(){
		defer close(p.smsTimerStop)
		for{
  			select {
  			case <-p.smsTimer.C:
				fmt.Println("p.smsTimer已到，仅仅p.smsAlarmIsReady变为了",p.smsAlarmIsReady,time.Now().Format("20060102150405"))
				p.smsAlarmIsReady =true
				p.smsTimer.Reset(time.Duration(p.smsTimerLimitSec) *time.Second)
			case stop := <-p.smsTimerStop:
				if stop {goto CLEANUP}
			}
		}
		
		CLEANUP:
		//跳出for循环可确保再没有指针指向这个管道，下面这句确保管道内的数据排空
		//从而在Stop()后实现有效的内存回收
		if len(p.smsTimer.C)>0 {<-p.smsTimer.C}
		p.smsTimer.Stop()
	}()
}	

//func NewMYSQLTimer(){builder.newMYSQLTimer()}
func (p *AlarmBuilder)newMYSQLTimer(){
	if p.mysqlTimer !=nil    { p.mysqlTimerStop <-true;    time.Sleep(time.Second) } 

	p.mysqlAlarmIsReady =true
	p.mysqlTimerStop =make(chan bool)
	p.mysqlTimer =time.NewTimer(time.Duration(p.mysqlTimerLimitSec) * time.Second)

	go func(){
		defer close(p.mysqlTimerStop)
		for{
			select{
			case <-p.mysqlTimer.C:
				fmt.Println("p.mysqltimer已到，仅仅p.mysqlAlarmIsReady变为了",p.mysqlAlarmIsReady,time.Now().Format("20060102150405"))
				p.mysqlAlarmIsReady =true
				p.mysqlTimer.Reset(time.Duration(p.mysqlTimerLimitSec)  *time.Second)
			case  stop := <-p.mysqlTimerStop:
				if stop { goto CLEANUP}
			}
		}
		
		CLEANUP:
		//跳出for循环可确保再没有指针指向这个管道，下面这句确保管道内的数据排空
		//从而在Stop()后实现有效的内存回收
		if len(p.mysqlTimer.C)>0 {<-p.mysqlTimer.C}
		p.mysqlTimer.Stop()
	}()
}

//两个管道的子线程生产者
func StartFilter(ndch chan nodedo.NodeDo){builder.StartFilter(ndch)}
func (p *AlarmBuilder)StartFilter(ndch chan nodedo.NodeDo){
	go func(){
		for nd := range ndch{
			issafe, smsarr, alarmdbentity :=p.e.JudgeOneNodeDo(nd)

			if issafe {continue}

			if p.smsAlarmIsReady{
				go func(){
					for _,v :=range smsarr {
						p.SMSAlarmCh <-[]byte(v)
					}
				}() 
				p.smsAlarmIsReady =false
				p.smsTimer.Reset(time.Duration(p.smsTimerLimitSec) *time.Second)
			}

			if p.mysqlAlarmIsReady{
				go func(){
					p.MYSQLAlarmCh <-alarmdbentity
				}()
				p.mysqlAlarmIsReady =false
				p.mysqlTimer.Reset(time.Duration(p.mysqlTimerLimitSec)  *time.Second)
			}
		}
	}()
}


//当前未设定消费者
func GenerateSMSbyteCh(){builder.GenerateSMSbyteCh()}
func(p *AlarmBuilder)GenerateSMSbyteCh(){
	if p.SMSAlarmCh !=nil {
		close(p.SMSAlarmCh)
	}

	p.SMSAlarmCh =make(chan []byte)
}

//当前未设定消费者
func GenerateMYSQLAlarmCh(){builder.GenerateMYSQLAlarmCh()}
func(p *AlarmBuilder)GenerateMYSQLAlarmCh(){
	if p.mysqlAlarmCh !=nil{
		close(p.mysqlAlarmCh)
	}

	p.MYSQLAlarmCh =make(chan *mysql.Alarm)
}

func Quit(){builder.Quit()}
func (p *AlarmBuilder)Quit(){
	p.mysqlTimerStop<-true//对应的计时器被销毁之后会立刻defer close该管道，因此就不用在这里close了
	p.smsTimerStop<-true//对应的计时器被销毁之后会立刻defer close该管道，因此就不用在这里close了
	close(p.smsAlarmCh)//是为了返回给上层关闭事件
	close(p.mysqlAlarmCh)//是为了返回给上层关闭事件
}