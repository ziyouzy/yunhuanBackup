//这里要用到timer而不是tikcer
//出于对节省资源的考虑，这里会先判定是否超时，如超时才会去进行下一步对NodeDo的相关计算
package alarm

import(
	"time"
)


//这里模仿了time包的NewTimer的设计模式，New出来的对象生命周期很可能为主函数
func NewAlarmFilterObject(base map[string]interface{}) *AlarmFilterObject{
	af := AlarmFilterObject{}
	at, smstimerlimitmin, mysqltimerlimitmin := NewAlramTemplate(base)

	//实例化内部字段（AT用来监测某个NodeDo是否超限）
	af.AT =at

	af.SMStimerLimitSec = smstimerlimitmin * 60
	af.MYSQLtimerLimitSec  = mysqltimerlimitmin * 60

	af.SMSAlarmIsReady =true
	af.MYSQLAlarmIsReady =true

	af.SMSAlarmCh =make(chan []byte)
	af.MYSQLAlarmCh =make(chan *model.AlarmEntity)
	
	af.InitSMSTimer()
	af.InitMYSQLTimer()
	return af
}

type AlarmFilterObject struct{
	SMStimer time.Timer
	SMStimerLimitSec float64
	SMSAlarmIsReady bool
	SMSAlarmCh chan []byte 

	MYSQLtimer time.Timer
	MYSQLtimerLimitSec float64
	MYSQLAlarmIsReady bool
	MYSQLAlarmCh chan *model.AlarmEntity

	//这里采用了组合的设计思路
	//AlarmTemplate结构体只有一个方法，判断一个NodeDo是否超限
	//这样设计是在遵顼分层的设计思路，也就是纯粹的为了分层而去采用了组合
	//基于能组合就不继承的原则，这里无论是组合还是继承都是合理的，所以既然地位相当那还是优先组合吧
	//退一步讲，就算是继承了，唯一的原因也只是为了分层思路而继承
	at AlarmTemplate

	done chan bool
}

func (p *AlarmFilterObject)InitSMSTimer(){
	p.SMStimer =time.NewTimer(p.SMStimerLimitSec * time.Second)
	go func(){
		for {
			select {
			case <-done:
				break
			case <-p.SMStimer.C:
				p.SMSAlarmIsReady =true				
			}
		}    
	}()
}

func (p *AlarmFilterObject)InitMYSQLTimer(){
	p.MYSQLtimer =time.NewTimer(p.MYSQLtimerLimitSec * time.Second)
	go func(){
		for {
			select {
			case <-done:
				break
			case <-p.MYSQLtimer.C:
				p.MYSQLAlarmIsReady =true				
			}
		}    
	}()
}

func (p *AlarmFilterObject)Filter(nd NodeDo)bool{
	issafe :=false
	newalarm :=p.at.CreateAlarm(nd)
	if newalarm ==nil{
		issafe =true
	}

	if p.SMSAlarmIsReady{
		for _,v :=range newalarm.SMS{
			p.SMSAlarmCh <-v
		} 
		p.SMSAlarmIsReady =false
		if !p.SMStimer.Stop() {
			<-p.SMStimer.C
		}
		p.SMStimer.Reset(p.SMStimerLimitSec * time.Second)
	}

	if p.MYSQLAlarmIsReady{
		p.MYSQLAlarmCh <-&(newalarm.AlarmEntityTemplate)
		p.MYSQLAlarmIsReady =false
		if !p.MYSQLtimer.Stop() {
			<-p.MYSQLtimer.C
		}
		p.MYSQLtimer.Reset(p.MYSQLtimerLimitSec * time.Second)
	}

	return
}

func (p *AlarmFilterObject)Quit(){
	p.done <- true
	close(p.SMSAlarmCh)
	close(p.MYSQLAlarmCh)
	close(p.done)
}