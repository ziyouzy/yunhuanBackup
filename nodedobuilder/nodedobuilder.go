// 从engine到nodedobuilder的转化本质上是运用了“组合”的编程思想
// 当监测到viper配置文档发生改变时，当前整个NodeDoBuilder整体都会被重置，而不是内部某个字段被重置
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

	TicketStep int
	FlushTicket *time.Ticker
	lock *sync.Mutex

	stopNodeCh chan bool

	NodeDoCh chan nodedo.NodeDo
}


func Load(step int, sourcefromviper map[string]interface{}){builder =BuildNodeDoBuilder(step, sourcefromviper)}
//这里模仿了time包的NewTimer的设计模式，New出来的对象生命周期为主函数
func BuildNodeDoBuilder(step int,sourcefromviper map[string]interface{}) *NodeDoBuilder{
	builder :=NodeDoBuilder{}
	builder.e =NewEngine(sourcefromviper)

	builder.TicketStep =step
	builder.lock =new(sync.Mutex)
	builder.stopNodeCh =make(chan bool)
	return &builder
}


//json文档改变后需要从新获得该管道
func GenerateNodeDoCh(){builder.GenerateNodeDoCh()}
//结合定时器生成NodeDo管道，里面的每个NodeDo都是最终的结果
//上层会基于这一结果进行告警判定，以及用字符串的形式发送字节数组给前端的操作
func (p *NodeDoBuilder)GenerateNodeDoCh(){
	p.FlushTicket =time.NewTicker(time.Duration(p.TicketStep) * time.Second)
	p.NodeDoCh = make(chan nodedo.NodeDo)
	//ticker的消费者与ch的生产者都在这个子携程中
	//ticker在NewTicker的同时，其自身就是生产者，所以消费者必然在其之后
	//select是ch的生产者，消费者会在上层实现
	go func(){
		//当done管道收到true时，在这里优雅的关闭该管道即可，因为他不是结构体的字段，无法在Quit方法内关闭
		for{ 
			select {
			case <-p.FlushTicket.C:
				p.lock.Lock()
				for _,v := range p.e{

					/** 超时的衡量标准在于p.FlushTicket.C的激活时间与旧时间戳之间的差值
					  * 也就是说，必须要在p.FlushTicket.C这一刻进行对超时的判断操作
					  */

					v.JudgeTimeOut()
					p.NodeDoCh <-v
				}
				p.lock.Unlock()

			case stop :=<-p.stopNodeCh:
				if stop{ 
					goto CLEANUP
				}

			}
		}

		CLEANUP:
		if len(p.FlushTicket.C)>0{
			fmt.Println("清空nodedobuilder.FlushTicker.C管道中的残留内容：",<-p.FlushTicket.C)
		}
		p.FlushTicket.Stop()  
		close(p.NodeDoCh)
	}()
}
func GetNodeDoCh()chan nodedo.NodeDo{ return builder.NodeDoCh }

/*-------------------------------*/
/** 运行该函数前
  * 需确保结构体内部的engine字段以实例化
  * (p.e即为engine)
  */
func StartEngine(pnch chan physicalnode.PhysicalNode, pn physicalnode.PhysicalNode){ 
/*-------------------------------*/
	if pnch !=nil&&pn ==nil { builder.StartEnginePhysicalNodeCh(pnch);        return }
	if pnch ==nil&&pn !=nil { builder.StartEnginePhysicalNode(pn);        return }

	fmt.Println("StartEngine参数填写错误(都非nil或都为nil是都不允许的)")
	return
}

/** Engine是个map
  * key 举例: "494f3031f10201-tcpsocket-do3-bool"
  * value则是实实在在的NodeDo
  * StartEngine函数的意义在于基于获取PhysicalNode节点所发来的PhysicalNode数据接口对象
  * 其发送的频率就是更新核心map的频率
  */

func (p *NodeDoBuilder)StartEnginePhysicalNodeCh( pnch chan physicalnode.PhysicalNode ){
	go func(){
		for pn :=range pnch{
			p.StartEnginePhysicalNode(pn)
		}
	}()
}

func (p *NodeDoBuilder)StartEnginePhysicalNode( pn physicalnode.PhysicalNode ){
	/* PhysicalNode.SelectHandlerAndTage返回值举例："494f3031f10201","tcpsocket"*/
	handler, tag :=pn.SelectHandlerAndTag()
	p.lock.Lock()
	for k,_ :=range p.e{

		/** 每当传来一个physicalnode实体时会判定这个实体在json文档实体中有没有实现对应的关系
		  * 这个判定的过程中每一个physicalnode都会对应一次engine对象的for循环
		  * 同时一个physicalnode可以在他所对应的for循环结束前多次触发nodedo的更新事件
		  * 这一特性体现在如下语句：
		  * if !strings. Contains(k, fmt.Sprintf("%s-%s",handler,tag)){ continue }
		  */

		/* PhysicalNode.SelectOneValueAndTime返回值举例："value","time"*/
		nodeName :=strings.Split(k,"-")[2]
		nodeValue, timeUnixNano := pn.SelectOneValueAndTimeUnixNano(handler, tag, nodeName)
		p.e[k].UpdateOneNodeDo(nodeValue, timeUnixNano)
	}
	p.lock.Unlock()
}

func Destory(){builder.Destory()}
func (p *NodeDoBuilder)Destory(){
	p.stopNodeCh <- true//只负责关闭返回给上层的NodeDoCh管道
	close(p.stopNodeCh) 

	p.lock.Lock()
	for key, _ := range p.e{
		/* 在这里清空所有旧NodeDo，其实不用清空，只要不再有不引用只向这个结构体就行了*/
		delete(p.e, key)
	}
	p.lock.Unlock()
}