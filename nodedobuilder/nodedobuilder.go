//从engine到nodedobuilder的转化本质上是运用了“组合”的编程思想
package nodedobuilder

import(
	"time"
	"sync"
	"strings"
	"fmt"

	"github.com/ziyouzy/mylib/physicalnode"
	"github.com/ziyouzy/mylib/nodedo"
)

var builder *NodeDoBuilder 
type NodeDoBuilder struct{
	e Engine
	timeOutTriggerMap map[string]*time.Timer

	TicketStep int
	FlushTicket *time.Ticker
	lock *sync.Mutex

	quit chan bool
}


func LoadSingletonPattern(step int, sourcefromviper map[string]interface{}){builder =BuildNodeDoBuilder(step, sourcefromviper)}
//这里模仿了time包的NewTimer的设计模式，New出来的对象生命周期为主函数
func BuildNodeDoBuilder(step int,sourcefromviper map[string]interface{}) *NodeDoBuilder{
	builder :=NodeDoBuilder{}
	builder.e =NewEngine(sourcefromviper)

	for key, nodedo :=range builder.e{
		t :=time.NewTimer(time.Duration(nodedo.timeoutsec() ) * time.Second)//timeOutTriggerMap
		go func(nd nodedo.Nodedo){
			for{
				select{
				case <-t.C:
					nd.TimeOut()
				}
			}
		}(nodedo)
		builder.timeOutTriggerMap[key] =t
	}

	builder.TicketStep =step
	builder.lock =new(sync.Mutex)
	builder.quit =make(chan bool)
	return &builder
}


func GenerateNodeDoCh()chan nodedo.NodeDo{return builder.GenerateNodeDoCh()}
//结合定时器生成NodeDo管道，里面的每个NodeDo都是最终的结果
//上层会基于这一结果进行告警判定，以及用字符串的形式发送字节数组给前端的操作
func (p *NodeDoBuilder)GenerateNodeDoCh()chan nodedo.NodeDo{
	p.FlushTicket =time.NewTicker(time.Duration(p.TicketStep) * time.Second)
	nodeDoCh := make(chan nodedo.NodeDo)
	//ticker的消费者与ch的生产者都在这个子携程中
	//ticker在NewTicker的同时，其自身就是生产者，所以消费者必然在其之后
	//select是ch的生产者，消费者会在上层实现
	go func(){
		//当done管道收到true时，在这里优雅的关闭该管道即可，因为他不是结构体的字段，无法在Quit方法内关闭
		defer close(nodeDoCh)
		for range p.FlushTicket.C{
			p.lock.Lock()
			for _,v := range p.e{
				nodeDoCh <-v
			}
			p.lock.Unlock()

			select {
			case <-p.quit:
				break
			default:
			}
		}

		if len(p.FlushTicket.C)>0{
			fmt.Println("清空nodedobuilder.FlushTicker.C管道中的残留内容：",<-p.FlushTicket.C)
		}
		p.FlushTicket.Stop()   
	}()
	return nodeDoCh
}


//运行该函数前，需确保结构体内部的engine字段以实例化（p.e即为engine）
func Engineing(pn physicalnode.PhysicalNode){builder.Engineing(pn)}
//Engine是个map，key 举例: "494f3031f10201-tcpsocket-do3-bool"，而value则是实实在在的NodeDo
//Engineing函数的意义在于基于获取PhysicalNode节点所发来的频率更新核心map
func (p *NodeDoBuilder)Engineing(pn physicalnode.PhysicalNode){
	//PhysicalNode.SelectHandlerAndTage返回值举例："494f3031f10201","tcpsocket"
	handler, tag :=pn.SelectHandlerAndTag()
	p.lock.Lock()
	for k,_ :=range p.e{
		//每当传来一个physicalnode实体时，会判定这个实体在json文档实体中，有没有实现对应的关系
		//这个判定的过程中每一个physicalnode都会对应一次engine对象的for循环
		//同时一个physicalnode可以在他所对应的for循环结束前多次触发nodedo的更新事件
		if !strings. Contains(k,fmt.Sprintf("%s-%s",handler,tag)){
			continue
		}
		nodename :=strings.Split(k,"-")[2]

		//PhysicalNode.SelectOneValueAndTime返回值举例："value","time"
		pvalue,ptime := pn.SelectOneValueAndTime(handler, tag, nodename)
		p.e[k].UpdateOneNodeDo(pvalue,ptime)
	}
	p.lock.Unlock()
}

func TimeOut(){

}

func Quit(){builder.Quit()}
func (p *NodeDoBuilder)Quit(){
	fmt.Println("为啥NodeDo就关闭了?")
	p.quit <- true
	close(p.quit)
}