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
	SMSticker *time.Ticker
	SMStickerLimitSec float64
	SMSAlarmIsReady bool
	SMSAlarmCh chan []byte 

	MYSQLticker *time.Ticker
	MYSQLtickerLimitSec float64
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
	e, smstickerlimitmin, mysqltickerlimitmin := NewEngine(base)
	fmt.Println("smstickerlimitmin:",smstickerlimitmin,"mysqltickerlimitmin:",mysqltickerlimitmin)

	//实例化内部字段（内部的Engine才会真正进行监测某个NodeDo是否超限）
	ac.e =e
	ac.SMStickerLimitSec = smstickerlimitmin * 60
	ac.MYSQLtickerLimitSec  = mysqltickerlimitmin * 60

	ac.quit =make(chan bool)
	
	ac.initSMSTicker()
	ac.initMYSQLTicker()
	return &ac
}


func InitSMSTicker(){ac.initSMSTicker()}
func (p *AlarmController)initSMSTicker(){
	if p.SMSticker ==nil{
		//p.SMSticker =time.NewTimer(time.Duration(p.SMStickerLimitSec) * time.Second)
		p.SMSticker =time.NewTicker(3 * time.Second)
		p.SMSAlarmIsReady =true
	
		//消费者子携程
		go func(){
			for {
				select {
				case <-p.quit:
					break
				case <-p.SMSticker.C:
					p.SMSAlarmIsReady =true				
					fmt.Println("p.SMStimer已到，仅仅p.SMSAlarmIsReady变为了",p.SMSAlarmIsReady)		
				}
			}    
		}()

	}else{

		/*下面只做了两件事：停止+重置*/
		if !p.SMSticker.Stop() {
			select{
			case <-p.SMSticker.C:
			default:
			}
		}
		//p.SMStimer.Reset(time.Duration(p.SMStimerLimitSec) * time.Second)
		//p.SMSticker.Reset(time.Duration(3 * time.Second))
		p.SMSticker =time.NewTicker(3 * time.Second)
		/*----*/
	}
}

func InitMYSQLTicker(){ac.initMYSQLTicker()}
func (p *AlarmController)initMYSQLTicker(){
	if p.MYSQLticker ==nil{ 
		//p.MYSQLticker =time.NewTicker(time.Duration(p.MYSQLtickerLimitSec) * time.Second)
		p.MYSQLticker =time.NewTicker(3 * time.Second)
		p.MYSQLAlarmIsReady =true

		//消费者子携程
		go func(){
			for {
				select {
				case <-p.quit:
					break
				case <-p.MYSQLticker.C:
					p.MYSQLAlarmIsReady =true	
					fmt.Println("p.MYSQLtimer已到，仅仅p.MYSQLAlarmIsReady变为了",p.MYSQLAlarmIsReady)			
				}
			}    
		}()

	}else{

		/*下面只做了两件事：停止+重置*/
		if !p.MYSQLticker.Stop() {
			select{
			case <-p.MYSQLticker.C:
			default:
			}
		}
		//p.MYSQLticker.Reset(time.Duration(p.MYSQLtickerLimitSec) * time.Second)
		p.MYSQLticker.Reset(3 * time.Second)
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
func Filter(ndch chan nodedo.NodeDo)/*chan bool*/{/*return */ac.Filter(ndch)}
func (p *AlarmController)Filter(ndch chan nodedo.NodeDo)/*chan bool*/{
	//issafech :=make(chan bool)
	go func(){
		for nd := range ndch{
			issafe, smsarr, alarmdbentity :=p.e.JudgeOneNodeDo(nd)
			fmt.Println("issafe, smsarr, alarmdbentity-:")
			fmt.Println(issafe, smsarr, alarmdbentity)
			if issafe{
				//issafech<-true
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

				/*下面只做了两件事：停止+重置*/
				if !p.SMStimer.Stop() {
					fmt.Println("p.SMStimer.Stop()==false")
					select{
					case <-p.SMStimer.C:
					default:
					}
				}
				//p.SMStimer.Reset(time.Duration(p.SMStimerLimitSec) * time.Second)
				p.SMStimer.Reset(3 * time.Second)
				/*----*/
			}

			if p.MYSQLAlarmIsReady{
				go func(){
					p.MYSQLAlarmCh <-alarmdbentity
				}()
				p.MYSQLAlarmIsReady =false

				/*下面只做了两件事：停止+重置*/
				if !p.MYSQLtimer.Stop() {
					fmt.Println("p.MYSQLtimer.Stop()=false")
					select{
					case <-p.MYSQLtimer.C:
					default:
					}
				}
				//p.MYSQLtimer.Reset(time.Duration(p.MYSQLtimerLimitSec) * time.Second)
				p.MYSQLtimer.Reset(3 * time.Second)
				/*----*/
			}

			fmt.Println("false issafech-a")
			//issafech <-false
			fmt.Println("false issafech-b")
		}
	}()

	//return issafech
}

func Quit(){ac.Quit()}
func (p *AlarmController)Quit(){
	p.quit <- true
	close(p.SMSAlarmCh)
	close(p.MYSQLAlarmCh)
	close(p.quit)
}