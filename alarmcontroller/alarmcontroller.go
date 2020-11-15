//这里要用到timer而不是tikcer
//出于对节省资源的考虑，这里会先判定是否超时，如超时才会去进行下一步对NodeDo的相关计算
package alarmcontroller

import(
	"time"

	"github.com/ziyouzy/mylib/model"	
	"github.com/ziyouzy/mylib/nodedo"	
)


var ac *AlarmController
type AlarmController struct{
	SMStimer *time.Timer
	SMStimerLimitSec float64
	SMSAlarmIsReady bool
	SMSAlarmCh chan []byte 

	MYSQLtimer *time.Timer
	MYSQLtimerLimitSec float64
	MYSQLAlarmIsReady bool
	MYSQLAlarmCh chan *model.AlarmEntity

	//这样设计是在遵顼分层的设计思路，也就是纯粹的为了分层而去采用了组合
	//基于能组合就不继承的原则，这里无论是组合还是继承都是合理的，所以既然地位相当那还是优先组合吧
	//退一步讲，就算是继承了，唯一的原因也只是为了分层思路而继承
	e *Engine

	quit chan bool
}

func AssembleEngine(base map[string]interface{}){ac =BuildAlarmController(base)}
//这里模仿了time包的NewTimer的设计模式，New出来的对象生命周期很可能为主函数
func BuildAlarmController(base map[string]interface{}) *AlarmController{
	ac := AlarmController{}
	e, smstimerlimitmin, mysqltimerlimitmin := NewEngine(base)

	//实例化内部字段（内部的Engine才会真正进行监测某个NodeDo是否超限）
	ac.e =e
	ac.SMStimerLimitSec = smstimerlimitmin * 60
	ac.MYSQLtimerLimitSec  = mysqltimerlimitmin * 60

	ac.SMSAlarmIsReady =true
	ac.MYSQLAlarmIsReady =true

	ac.SMSAlarmCh =make(chan []byte)
	ac.MYSQLAlarmCh =make(chan *model.AlarmEntity)
	ac.quit =make(chan bool)
	
	ac.initSMSTimer()
	ac.initMYSQLTimer()
	return &ac
}


func InitSMSTimer(){ac.initSMSTimer()}
func (p *AlarmController)initSMSTimer(){
	p.SMStimer =time.NewTimer(time.Duration(p.SMStimerLimitSec) * time.Second)
	go func(){
		for {
			select {
			case <-p.quit:
				break
			case <-p.SMStimer.C:
				p.SMSAlarmIsReady =true				
			}
		}    
	}()
}

func InitMYSQLTimer(){ac.initMYSQLTimer()}
func (p *AlarmController)initMYSQLTimer(){
	p.MYSQLtimer =time.NewTimer(time.Duration(p.MYSQLtimerLimitSec) * time.Second)
	go func(){
		for {
			select {
			case <-p.quit:
				break
			case <-p.MYSQLtimer.C:
				p.MYSQLAlarmIsReady =true				
			}
		}    
	}()
}

func Filter(nd nodedo.NodeDo)bool{return ac.Filter(nd)}
func (p *AlarmController)Filter(nd nodedo.NodeDo)bool{
	issafe, smsArr, alarmDbEntity :=p.e.JudgeOneNodeDo(nd)
	if issafe{
		return issafe
	}

	if p.SMSAlarmIsReady{
		for _,v :=range smsArr {
			p.SMSAlarmCh <-[]byte(v)
		} 
		p.SMSAlarmIsReady =false
		if !p.SMStimer.Stop() {
			<-p.SMStimer.C
		}
		p.SMStimer.Reset(time.Duration(p.SMStimerLimitSec) * time.Second)
	}

	if p.MYSQLAlarmIsReady{
		p.MYSQLAlarmCh <-alarmDbEntity
		p.MYSQLAlarmIsReady =false
		if !p.MYSQLtimer.Stop() {
			<-p.MYSQLtimer.C
		}
		p.MYSQLtimer.Reset(time.Duration(p.MYSQLtimerLimitSec) * time.Second)
	}

	return issafe
}

func Quit(){ac.Quit()}
func (p *AlarmController)Quit(){
	p.quit <- true
	close(p.SMSAlarmCh)
	close(p.MYSQLAlarmCh)
	close(p.quit)
}