//从engine到nodedocontroller的转化本质上是运用了“组合”的编程思想
package nodedocontroller

import(
	"time"
	"sync"
	"strings"
	"fmt"

	"github.com/ziyouzy/mylib/physicalnode"
	"github.com/ziyouzy/mylib/nodedo"
)

var ndc *NodeDoController 
type NodeDoController struct{
	e Engine
	TicketStep int

	FlushTicket *time.Ticker
	lock *sync.Mutex
	quit chan bool
}


func LoadSingletonPattern(step int, base map[string]interface{}){ndc =BuildNodeDoController(step, base)}
//这里模仿了time包的NewTimer的设计模式，New出来的对象生命周期为主函数
func BuildNodeDoController(step int,base map[string]interface{}) *NodeDoController{
	ndc :=NodeDoController{}
	ndc.e =NewEngine(base)
	ndc.TicketStep =step
	ndc.lock =new(sync.Mutex)
	ndc.quit =make(chan bool)
	return &ndc
}


func GenerateNodeDoCh()chan nodedo.NodeDo{return ndc.GenerateNodeDoCh()}
//结合定时器生成NodeDo管道，里面的每个NodeDo都是最终的结果
//上层会基于这一结果进行告警判定，以及用字符串的形式发送字节数组给前端的操作
func (p *NodeDoController)GenerateNodeDoCh()chan nodedo.NodeDo{
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
			fmt.Println("清空nodedocontroller.FlushTicker.C管道中的残留内容：",<-p.FlushTicket.C)
		}
		p.FlushTicket.Stop()   
	}()
	return nodeDoCh
}



func Engineing(pn physicalnode.PhysicalNode){ndc.Engineing(pn)}
//Engine是个map，key 举例: "494f3031f10201-tcpsocket-do3-bool"，而value则是实实在在的NodeDo
//Engineing函数的意义在于基于获取PhysicalNode节点所发来的频率更新核心map
//而当前包也会负责根据所设定的频率生成并发送NodeDo的独立结构体
//PhysicalNode.SelectHandlerAndTage返回值举例："494f3031f10201","tcpsocket"
//PhysicalNode.SelectOneValueAndTime返回值举例："value","time"
func (p *NodeDoController)Engineing(pn physicalnode.PhysicalNode){
	handler, tag :=pn.SelectHandlerAndTag()
	p.lock.Lock()
	for k,_ :=range p.e{
		if !strings. Contains(k,fmt.Sprintf("%s-%s",handler,tag)){
			continue
		}
		nodename :=strings.Split(k,"-")[2]

		pvalue,ptime := pn.SelectOneValueAndTime(handler, tag, nodename)
		fmt.Println("pvalue,ptime :",pvalue,ptime)
		p.e[k].UpdateOneNodeDo(pvalue,ptime)
	}
	p.lock.Unlock()
}

func Quit(){ndc.Quit()}
func (p *NodeDoController)Quit(){
	fmt.Println("为啥NodeDo就关闭了?")
	p.quit <- true
	close(p.quit)
}