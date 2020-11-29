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
	SMStimer *time.Timer
	SMStimerLimitSec float64
	SMSAlarmIsReady bool
	SMSAlarmCh chan []byte 

	MYSQLtimer *time.Timer
	MYSQLtimerLimitSec float64
	MYSQLAlarmIsReady bool
	MYSQLAlarmCh chan *mysql.Alarm

	//这样设计是在遵顼分层的设计思路，也就是纯粹的为了分层而去采用了组合
	//基于能组合就不继承的原则，这里无论是组合还是继承都是合理的，所以既然地位相当那还是优先组合吧
	//退一步讲，就算是继承了，唯一的原因也只是为了分层思路而继承
	e *Engine

	quit chan bool
}

func LoadSingletonPattern(sourcefromviper map[string]interface{}){builder =BuildAlarmBuilder(sourcefromviper)}
//这里模仿了time包的NewTimer的设计模式，New出来的对象生命周期很可能为主函数
func BuildAlarmBuilder(sourcefromviper map[string]interface{}) *AlarmBuilder{
	builder := AlarmBuilder{}
	e, smstimerlimitmin, mysqltimerlimitmin := NewEngine(sourcefromviper)
	fmt.Println("smstimerlimitmin:",smstimerlimitmin,"mysqltimerlimitmin:",mysqltimerlimitmin)

	//实例化内部字段（内部的Engine才会真正进行监测某个NodeDo是否超限）
	builder.e =e
	builder.SMStimerLimitSec = smstimerlimitmin * 60
	builder.MYSQLtimerLimitSec  = mysqltimerlimitmin * 60

	builder.quit =make(chan bool)
	
	builder.newSMSTimer()
	builder.newMYSQLTimer()
	return &builder
}


func NewSMSTimer(){builder.newSMSTimer()}
func (p *AlarmBuilder)newSMSTimer(){
	if p.SMStimer !=nil{
		p.quit<-true
		time.Sleep(time.Second)
	}

	//p.SMSticker =time.NewTimer(time.Duration(p.SMStickerLimitSec) * time.Second)
	p.SMStimer =time.NewTimer(3 * time.Second)
	p.SMSAlarmIsReady =true

	//消费者子携程
	go func(){
		for{
  			select {
  			case <-p.SMStimer.C:
				fmt.Println("p.SMStimer已到，仅仅p.SMSAlarmIsReady变为了",p.SMSAlarmIsReady,time.Now().Format("20060102150405"))
				p.SMSAlarmIsReady =true
				p.SMStimer.Reset(3 *time.Second)
			case <-p.quit:
				break
			}
		}
		p.SMStimer.Stop()
	}()
}	

func NewMYSQLTimer(){builder.newMYSQLTimer()}
func (p *AlarmBuilder)newMYSQLTimer(){
	if p.MYSQLtimer !=nil{
		p.quit<-true
		time.Sleep(time.Second)
	} 

	//p.MYSQLticker =time.NewTicker(time.Duration(p.MYSQLtickerLimitSec) * time.Second)
	p.MYSQLtimer =time.NewTimer(3 * time.Second)
	p.MYSQLAlarmIsReady =true

	//消费者子携程
	go func(){
		for{
			select{
			case <-p.MYSQLtimer.C:
				fmt.Println("p.MYSQLtimer已到，仅仅p.MYSQLAlarmIsReady变为了",p.MYSQLAlarmIsReady,time.Now().Format("20060102150405"))
				p.MYSQLAlarmIsReady =true
				p.MYSQLtimer.Reset(3 *time.Second)
			case <-p.quit:
				break
			}
		}
		p.MYSQLtimer.Stop()
	}()
}
// 	go func(){
// 		for range p.MYSQLticker.C{
// 			p.MYSQLAlarmIsReady =true	
// 			fmt.Println("p.MYSQLtimer已到，仅仅p.MYSQLAlarmIsReady变为了",p.MYSQLAlarmIsReady)			
// 			select {
// 			case <-p.quit:
// 				break
// 			default:
// 			}
// 		}
// 		if len(p.MYSQLticker.C)>0{
// 			fmt.Println("清空MYSQLticker.C管道中的残留内容：",<-p.MYSQLticker.C)
// 		}
// 		p.MYSQLticker.Stop()   
// 	}()
// }
// 			for {
				
// 				case <-p.MYSQLticker.C:
// 					p
// 				}
// 			}    
// 		}()

// 	}else{

// 		/*下面只做了两件事：停止+重置*/
// 		if !p.MYSQLticker.Stop() {
// 			select{
// 			case <-p.MYSQLticker.C:
// 			default:
// 			}
// 		}
// 		//p.MYSQLticker.Reset(time.Duration(p.MYSQLtickerLimitSec) * time.Second)
// 		p.MYSQLticker.Reset(3 * time.Second)
// 		/*----*/
// 	}
// }

//当前未设定消费者
func GenerateSMSbyteCh()chan []byte{return builder.GenerateSMSbyteCh()}
func(p *AlarmBuilder)GenerateSMSbyteCh()chan []byte{
	p.SMSAlarmCh =make(chan []byte)
	return p.SMSAlarmCh
}

//当前未设定消费者
func GenerateMYSQLAlarmCh()chan *mysql.Alarm{return builder.GenerateMYSQLAlarmCh()}
func(p *AlarmBuilder)GenerateMYSQLAlarmCh()chan *mysql.Alarm{
	p.MYSQLAlarmCh =make(chan *mysql.Alarm)
	return p.MYSQLAlarmCh
}

//两个管道的子线程生产者
//同时也会返回一个bool管道,也实现了这个管道的生产者
func Filter(ndch chan nodedo.NodeDo){builder.Filter(ndch)}
func (p *AlarmBuilder)Filter(ndch chan nodedo.NodeDo){
	go func(){
		for nd := range ndch{
			issafe, smsarr, alarmdbentity :=p.e.JudgeOneNodeDo(nd)
			fmt.Println("issafe, smsarr, alarmdbentity-:")
			fmt.Println(issafe, smsarr, alarmdbentity)
			if issafe{
				continue
			}

			//似乎是这里的问题，异常的nodedo会进入到这里，但是由于p.SMSAlarmIsReady==false，他不会把异常数据存入管道，而是直接执行for末尾的issafech <-false
			if p.SMSAlarmIsReady{
				go func(){
					for _,v :=range smsarr {
						p.SMSAlarmCh <-[]byte(v)
					}
				}() 
				p.SMSAlarmIsReady =false
				p.SMStimer.Reset(3 *time.Second)
			}

			if p.MYSQLAlarmIsReady{
				go func(){
					p.MYSQLAlarmCh <-alarmdbentity
				}()
				p.MYSQLAlarmIsReady =false
				p.MYSQLtimer.Reset(3 *time.Second)
			}
		}
	}()
}

func Quit(){builder.Quit()}
func (p *AlarmBuilder)Quit(){
	p.quit <- true
	close(p.SMSAlarmCh)
	close(p.MYSQLAlarmCh)
	close(p.quit)
}