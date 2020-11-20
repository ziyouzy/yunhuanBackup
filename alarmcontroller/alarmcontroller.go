//这里要用到timer而不是tikcer
//出于对节省资源的考虑，这里会先判定是否超时，如超时才会去进行下一步对NodeDo的相关计算
package alarmcontroller

import(
	"time"
	"fmt"

	"github.com/ziyouzy/mylib/model"	
	"github.com/ziyouzy/mylib/nodedo"	
)


var ac *AlarmController
type AlarmController struct{
	SMStimer *time.Timer
	SMStimerLimitSec int
	SMSAlarmIsReady bool
	SMSAlarmCh chan []byte 

	MYSQLtimer *time.Timer
	MYSQLtimerLimitSec int
	MYSQLAlarmIsReady bool
	MYSQLAlarmCh chan *model.AlarmEntity

	//这样设计是在遵顼分层的设计思路，也就是纯粹的为了分层而去采用了组合
	//基于能组合就不继承的原则，这里无论是组合还是继承都是合理的，所以既然地位相当那还是优先组合吧
	//退一步讲，就算是继承了，唯一的原因也只是为了分层思路而继承
	e *Engine

	quit chan bool
}

func LoadSingletonPattern(base map[string]interface{}){ac =BuildAlarmController(base)}
//这里模仿了time包的NewTimer的设计模式，New出来的对象生命周期很可能为主函数
func BuildAlarmController(base map[string]interface{}) *AlarmController{
	ac := AlarmController{}
	e, smstimerlimitmin, mysqltimerlimitmin := NewEngine(base)
	fmt.Println("smstimerlimitmin:",smstimerlimitmin,"mysqltimerlimitmin:",mysqltimerlimitmin)

	//实例化内部字段（内部的Engine才会真正进行监测某个NodeDo是否超限）
	ac.e =e
	ac.SMStimerLimitSec = int(smstimerlimitmin) * 60
	ac.MYSQLtimerLimitSec  = int(mysqltimerlimitmin) * 60

	ac.quit =make(chan bool)
	
	ac.initSMSTimer()
	ac.initMYSQLTimer()
	fmt.Println("怎么又build了一次AlarmController,ac.SMStimerLimitSec :",ac.SMStimerLimitSec,"ac.MYSQLtimerLimitSec:",ac.MYSQLtimerLimitSec)
	return &ac
}


func InitSMSTimer(){ac.initSMSTimer()}
func (p *AlarmController)initSMSTimer(){
	if p.SMStimer ==nil{
		fmt.Println("int(p.SMStimerLimitSec)):",int(p.SMStimerLimitSec),"p.SMStimerLimitSec:",p.SMStimerLimitSec)
		p.SMStimer =time.NewTimer(time.Duration(int(p.SMStimerLimitSec)) * time.Second)
		p.SMSAlarmIsReady =true
	
		//消费者子携程
		go func(){
			for {
				select {
				case <-p.quit:
					break
				case <-p.SMStimer.C:
					fmt.Println("p.SMStimer.C怎么这么块就更新了")
					p.SMSAlarmIsReady =true				
				}
			}    
		}()

	}else{

		/*下面只做了两件事：停止+重置*/
		if !p.SMStimer.Stop() {
			select{
			case <-p.SMStimer.C:
			default:
			}
		}
		p.SMStimer.Reset(time.Duration(p.SMStimerLimitSec) * time.Second)
		/*----*/
	}
}

func InitMYSQLTimer(){ac.initMYSQLTimer()}
func (p *AlarmController)initMYSQLTimer(){
	if p.MYSQLtimer ==nil{ 
		p.MYSQLtimer =time.NewTimer(time.Duration(p.MYSQLtimerLimitSec) * time.Second)
		p.MYSQLAlarmIsReady =true

		//消费者子携程
		go func(){
			for {
				select {
				case <-p.quit:
					break
				case <-p.MYSQLtimer.C:
					fmt.Println("p.MYSQLtimer.C怎么这么块就更新了")
					p.MYSQLAlarmIsReady =true				
				}
			}    
		}()

	}else{
		fmt.Println("MYSQLtimer :=nil")

		/*下面只做了两件事：停止+重置*/
		if !p.MYSQLtimer.Stop() {
			select{
			case <-p.MYSQLtimer.C:
			default:
			}
		}
		p.MYSQLtimer.Reset(time.Duration(p.MYSQLtimerLimitSec) * time.Second)
		/*----*/
	}
}

//当前未设定消费者
func GenerateSMSbyteCh()chan []byte{return ac.GenerateSMSbyteCh()}
func(p *AlarmController)GenerateSMSbyteCh()chan []byte{
	p.SMSAlarmCh =make(chan []byte)
	return p.SMSAlarmCh
}

//当前未设定消费者
func GenerateMYSQLEntityCh()chan *model.AlarmEntity{return ac.GenerateMYSQLEntityCh()}
func(p *AlarmController)GenerateMYSQLEntityCh()chan *model.AlarmEntity{
	p.MYSQLAlarmCh =make(chan *model.AlarmEntity)
	return p.MYSQLAlarmCh
}

//两个管道的子线程生产者
//同时也会返回一个bool管道,也实现了这个管道的生产者
func Filter(ndch chan nodedo.NodeDo)chan bool{return ac.Filter(ndch)}
func (p *AlarmController)Filter(ndch chan nodedo.NodeDo)chan bool{
	issafech :=make(chan bool)
	go func(){
		for nd := range ndch{		
			issafe, smsarr, alarmdbentity :=p.e.JudgeOneNodeDo(nd)
			if issafe{
				issafech<-true
				continue
			}

			if p.SMSAlarmIsReady{
				for _,v :=range smsarr {
					p.SMSAlarmCh <-[]byte(v)
				} 
				p.SMSAlarmIsReady =false

				/*下面只做了两件事：停止+重置*/
				if !p.SMStimer.Stop() {
					select{
					case <-p.SMStimer.C:
					default:
					}
				}
				p.SMStimer.Reset(time.Duration(p.SMStimerLimitSec) * time.Second)
				/*----*/
			}

			if p.MYSQLAlarmIsReady{
				p.MYSQLAlarmCh <-alarmdbentity
				p.MYSQLAlarmIsReady =false

				/*下面只做了两件事：停止+重置*/
				if !p.MYSQLtimer.Stop() {
					select{
					case <-p.MYSQLtimer.C:
					default:
					}
				}
				p.MYSQLtimer.Reset(time.Duration(p.MYSQLtimerLimitSec) * time.Second)
				/*----*/
			}

			issafech <-false
		}
	}()

	return issafech
}

func Quit(){ac.Quit()}
func (p *AlarmController)Quit(){
	p.quit <- true
	close(p.SMSAlarmCh)
	close(p.MYSQLAlarmCh)
	close(p.quit)
}