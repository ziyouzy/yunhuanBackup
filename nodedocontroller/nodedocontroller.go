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
	go func(){
		//当done管道收到true时，在这里优雅的关闭该管道即可，因为他不是结构体的字段，无法在Quit方法内关闭
		defer close(nodeDoCh)
		for {
			select {
			case <-p.quit:
				break
			case <-p.FlushTicket.C:
				p.lock.Lock()
				for _,v := range p.e{
					nodeDoCh <-v
				}
				p.lock.Unlock()
			}
		}    
	}()
	return nodeDoCh
}

func Engineing(pn physicalnode.PhysicalNode){ndc.Engineing(pn)}
//Engine是个map，key 举例: "494f3031f10201-tcpsocket-do3-bool"
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
		p.e[k].UpdateOneNodeDo(pvalue,ptime)
	}
	p.lock.Unlock()
}

func Quit(){ndc.Quit()}
func (p *NodeDoController)Quit(){
	p.quit <- true
	close(p.quit)
}