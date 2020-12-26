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
	smsTimerLimitMilliSecond int
	smsTimerStop chan bool
	smsAlarmIsReady bool
	smsAlarmCache map[string][]string
	SMSAlarmCh chan []byte 

	mysqlTimer *time.Timer
	mysqlTimerLimitMilliSecond int
	mysqlTimerStop chan bool
	mysqlAlarmIsReady bool
	mysqlAlarmCache map[string]*mysql.Alarm
	MYSQLAlarmCh chan *mysql.Alarm



	/** 这样设计是在遵顼分层的设计思路，也就是纯粹的为了分层而去采用了组合
	  * 基于能组合就不继承的原则，这里无论是组合还是继承都是合理的，所以既然地位相当那还是优先组合吧
	  * 退一步讲，就算是继承了，唯一的原因也只是为了分层思路而继承
	*/
	e *Engine
}

func Load(sourcefromviper map[string]interface{}){builder =BuildAlarmBuilder(sourcefromviper)}
//这里模仿了time包的NewTimer的设计模式，New出来的对象生命周期很可能为主函数
func BuildAlarmBuilder(sourcefromviper map[string]interface{}) *AlarmBuilder{
	builder := AlarmBuilder{}
	//实例化内部字段（内部的Engine才会真正进行监测某个NodeDo是否超限）
	builder.e, builder.smsTimerLimitMilliSecond, builder.mysqlTimerLimitMilliSecond = NewEngine(sourcefromviper)
	fmt.Println("builder.smsTimerLimitMilliSecond:", builder.smsTimerLimitMilliSecond , "builder.mysqlTimerLimitMilliSecond:", builder.mysqlTimerLimitMilliSecond)


	// builder.e =e
	// builder.smsTimerLimitMilliSecond = smstimerlimitmin * 60
	// builder.mysqlTimerLimitMilliSecond  = mysqltimerlimitmin * 60
	
	builder.newSMSTimer()
	builder.newMYSQLTimer()
	return &builder
}


//func NewSMSTimer(){builder.newSMSTimer()}
func (p *AlarmBuilder)newSMSTimer(){
	if p.smsTimer !=nil{p.smsTimerStop <-true;time.Sleep(time.Second)}

	p.smsTimer =time.NewTimer(time.Duration(p.smsTimerLimitMilliSecond) * time.Millisecond)
	p.smsAlarmIsReady =true
	p.smsTimerStop =make(chan bool)

	go func(){
		defer close(p.smsTimerStop)
		for{
  			select {
  			case <-p.smsTimer.C:
				fmt.Println("p.smsTimer已到，仅仅p.smsAlarmIsReady变为了",p.smsAlarmIsReady,time.Now().Format("20060102150405"))
				p.smsAlarmIsReady =true
				p.smsTimer.Reset(time.Duration(p.mysqlTimerLimitMilliSecond) *time.Millisecond)
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
	p.mysqlTimer =time.NewTimer(time.Duration(p.mysqlTimerLimitMilliSecond) * time.Millisecond)

	go func(){
		defer close(p.mysqlTimerStop)
		for{
			select{
			case <-p.mysqlTimer.C:
				fmt.Println("p.mysqltimer已到，仅仅p.mysqlAlarmIsReady变为了",p.mysqlAlarmIsReady,time.Now().Format("20060102150405"))
				p.mysqlAlarmIsReady =true
				p.mysqlTimer.Reset(time.Duration(p.mysqlTimerLimitMilliSecond)  *time.Millisecond)
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
func StartFilter(ndch chan nodedo.NodeDo, nd nodedo.NodeDo){
	if ndch !=nil&&nd ==nil { builder.StartFilterNodeDoCh(ndch);        return }
	if ndch ==nil&&nd !=nil { builder.StartFilterNodeDo(nd);        return }

	fmt.Println("StartFilter参数填写错误(都非nil或都为nil是都不允许的)")
	return
}
func (p *AlarmBuilder)StartFilterNodeDoCh(ndch chan nodedo.NodeDo){
	go func(){
		for nd := range ndch{
			p.StartFilterNodeDo(nd)
		}
	}()
}

func (p *AlarmBuilder)StartFilterNodeDo(nd nodedo.NodeDo){
	key, issafe, smsarr, alarmdbentity :=p.e.JudgeOneNodeDo(nd)

	if issafe { return }
	/** sms的发送方式和alarmEntity是不同的
	  * 如果同时存在多个传感器异常，则只sms一个
	  * 而对于alarmEntity，则需要将所有的告警节点一并录入
	  * 因此作出如下设计思路，无论是sms还是alarmEntity的缓存，一个周期内都不会保存同一个节点的告警信息
	  */
	
	if p.smsAlarmCache ==nil { p.smsAlarmCache =make(map[string][]string) };        if p.mysqlAlarmCache ==nil { p.mysqlAlarmCache =make(map[string]*mysql.Alarm) }

	if _,ok :=p.smsAlarmCache[key];        !ok { p.smsAlarmCache[key] = smsarr }   
	if _,ok :=p.mysqlAlarmCache[key];        !ok {  p.mysqlAlarmCache[key] = alarmdbentity }

	if p.smsAlarmIsReady{
		fmt.Println("test2-1 smsAlarmIsReady")
		for k,_ := range p.smsAlarmCache  {
			for _, sms := range p.smsAlarmCache[k]{
				p.SMSAlarmCh <-[]byte(sms)
			}
			p.smsAlarmCache[k] =(p.smsAlarmCache[k])[0:0]
		}
		p.smsAlarmCache =make(map[string][]string);        p.smsAlarmIsReady =false
		p.smsTimer.Reset(time.Duration(p.smsTimerLimitMilliSecond) *time.Millisecond)
	}

	if p.mysqlAlarmIsReady{
		fmt.Println("test2-2 mysqlAlarmIsReady")
		for _, v := range p.mysqlAlarmCache  {
			p.MYSQLAlarmCh <-v
		}
		p.mysqlAlarmCache =make(map[string]*mysql.Alarm) ;        p.mysqlAlarmIsReady =false
		p.mysqlTimer.Reset(time.Duration(p.mysqlTimerLimitMilliSecond)  *time.Millisecond)
	}
}


//当前未设定消费者
func GenerateSMSbyteCh(){builder.GenerateSMSbyteCh()}
func(p *AlarmBuilder)GenerateSMSbyteCh(){
	if p.SMSAlarmCh !=nil {
		close(p.SMSAlarmCh)
	}

	p.SMSAlarmCh =make(chan []byte)
}
func GetSMSAlarmbyteCh()chan []byte{ return builder.SMSAlarmCh }

//当前未设定消费者
func GenerateMYSQLAlarmCh(){builder.GenerateMYSQLAlarmCh()}
func(p *AlarmBuilder)GenerateMYSQLAlarmCh(){
	if p.MYSQLAlarmCh !=nil{
		close(p.MYSQLAlarmCh)
	}

	p.MYSQLAlarmCh =make(chan *mysql.Alarm)
}
func GetMYSQLAlarmEntityCh()chan *mysql.Alarm{ return builder.MYSQLAlarmCh }

func Destory(){builder.Destory()}
func (p *AlarmBuilder)Destory(){
	p.mysqlTimerStop<-true//对应的计时器被销毁之后会立刻defer close该管道，因此就不用在这里close了
	p.smsTimerStop<-true//对应的计时器被销毁之后会立刻defer close该管道，因此就不用在这里close了
	close(p.SMSAlarmCh)//是为了返回给上层关闭事件
	close(p.MYSQLAlarmCh)//是为了返回给上层关闭事件
}