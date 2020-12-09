//从engine到nodedobuilder的转化本质上是运用了“组合”的编程思想
//当监测到viper配置文档发生改变时，当前整个NodeDoBuilder整体都会被重置，而不是内部某个字段被重置
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
	//timeOutTriggerMap map[string]*time.Timer

	TicketStep int
	FlushTicket *time.Ticker
	lock *sync.Mutex

	stopNodeCh chan bool
}


func LoadSingletonPattern(step int, sourcefromviper map[string]interface{}){builder =BuildNodeDoBuilder(step, sourcefromviper)}
//这里模仿了time包的NewTimer的设计模式，New出来的对象生命周期为主函数
func BuildNodeDoBuilder(step int,sourcefromviper map[string]interface{}) *NodeDoBuilder{
	builder :=NodeDoBuilder{}
	builder.e =NewEngine(sourcefromviper)
	// builder.timeOutTriggerMap =make(map[string]*time.Timer)

	// for key, nodedo :=range builder.e{
	// 	timer :=time.NewTimer(time.Duration(nodedo.GetTimeOutSec()) * time.Second)//timeOutTriggerMap
	// 	go func(){
	// 		for{
	// 			if nodedo ==nil{goto CLEANUP}//这里的break只是为了配合Quit()函数实现最后一个析构步骤,也就是只负责跳出这个循环而不负责当前nodedo所对应timer的销毁
	// 			select{
	// 			case <-timer.C:
	// 				//只有一种情况nodedo会被销毁，那就是就的builder.e被销毁的时候，同时新的builder.e还未完成实例化，也就是json配置文档热更新的过程中
	// 				//这个过程中也会先调用nodedobuilder.Quit()，该函数会销毁每一个nodedo所对应的timer的管道，从而实现解开在这里的引用
	// 				nodedo.TimeOut(); timer.Reset(time.Duration(nodedo.GetTimeOutSec()) *time.Second)
	// 			}
	// 		}

	// 		CLEANUP:
	// 			if len(timer.C)>0{
	// 				<-timer.C
	// 			}
	// 			timer.Stop()
	// 			fmt.Println("nodedo.Timeout")

	// 	}()
	// 	//上边的NewEngine先实例化了builder.e，而timeOutTriggerMap的键名是从builder.e直接拿到的
	// 	builder.timeOutTriggerMap[key] =timer
	//}

	builder.TicketStep =step
	builder.lock =new(sync.Mutex)
	builder.stopNodeCh =make(chan bool)
	return &builder
}


//json文档改变后需要从新获得该管道
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
		for{ 
			select {
			case stop :=<-p.stopNodeCh:
				if stop{ 
					goto CLEANUP
				}
			case <-p.FlushTicket.C:
				p.lock.Lock()
				for _,v := range p.e{
					v.JudgeTimeOut()
					nodeDoCh <-v
				}
				p.lock.Unlock()
			}
		}

	CLEANUP:
		fmt.Println("stop NodeCh wtf1???")
		if len(p.FlushTicket.C)>0{
			fmt.Println("清空nodedobuilder.FlushTicker.C管道中的残留内容：",<-p.FlushTicket.C)
		}
		fmt.Println("stop NodeCh wtf2???")
		p.FlushTicket.Stop()  
		close(nodeDoCh)
		close(p.stopNodeCh) 
		fmt.Println("stop NodeCh wtf3???")
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
		if !strings. Contains(k,fmt.Sprintf("%s-%s",handler,tag)){ continue }
		//PhysicalNode.SelectOneValueAndTime返回值举例："value","time"
		nodename :=strings.Split(k,"-")[2];    physicalBytesValue, physicalTimeUnixNano := pn.SelectOneValueAndTimeUnixNano(handler, tag, nodename)
		p.e[k].UpdateOneNodeDo(string(physicalBytesValue), physicalTimeUnixNano)
		//go p.timeOutTriggerMap[k].Reset(time.Duration(sec) *time.Second)
	}
	p.lock.Unlock()
}

func Quit(){builder.Quit()}
func (p *NodeDoBuilder)Quit(){
	fmt.Println("0")
	p.stopNodeCh <- true//只负责关闭返回给上层的NodeDoCh管道
	fmt.Println("1")
	p.lock.Lock()
	for key, _ := range p.e{
		delete(p.e, key)
		// if len(p.timeOutTriggerMap[key].C)>0 {
		// 	<-p.timeOutTriggerMap[key].C
		// 	p.timeOutTriggerMap[key].Stop()
		// }
		// delete(p.timeOutTriggerMap, key)
	}
	fmt.Println("2")
	p.lock.Unlock()
}