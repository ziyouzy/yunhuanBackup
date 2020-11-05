//把缓存单独提取出来了，有些像是属于主函数的那几个handler(tcphandler,udphandler等等)
//之所以他是个缓存对象，也就是说实例化后他就会常驻与程序内存中了
//alarm.cache里面设计的并非是个ValueObject(变量容器)
//而是一个Filter(过滤器/拦截器)
//这里设计的很好，do包分成了两层
//一层是从json文档获得do的map
//另一层是生成ValueObject缓存
package do

import(
	"time"
	"sync"
	"strings"
	"fmt"

	"github.com/ziyouzy/mylib/physicalnode"
)

type NodeDoValueObject struct{
	NodeDoMap map[string]NodeDo
	FlushTicket *time.Ticker
	lock *sync.Mutex
	done chan bool
}

//这里模仿了time包的NewTimer的设计模式，New出来的对象生命周期为主函数
func NewNodeDoValueObj(step int,m map[string]NodeDo) *NodeDoValueObject{
	nd :=NodeDoValueObject{}
	nd.NodeDoMap =m
	nd.FlushTicket =time.NewTicker(time.Duration(step) * time.Second)
	return &nd
}

//结合定时器生成NodeDo管道，里面的每个NodeDo都是最终的结果
//上层会基于这一结果进行告警判定，以及用字符串的形式发送字节数组给前端的操作
func (p *NodeDoValueObject)CreateNodeDoCh()chan NodeDo{
	nodeDoCh := make(chan NodeDo)
	go func(){
		//当done管道收到true时，在这里优雅的关闭该管道即可，因为他不是结构体的字段，无法在Quit方法内关闭
		defer close(nodeDoCh)
		for {
			select {
			case <-p.done:
				break
			case <-p.FlushTicket.C:
				p.lock.Lock()
				for _,v := range p.NodeDoMap{
					nodeDoCh <-v
				}
				p.lock.Unlock()
			}
		}    
	}()
	return nodeDoCh
}

//NodeDoMap.Key 举例: "494f3031f10201-tcpsocket-do3-bool"
//GetHandlerTagForConfNodeMap()返回值举例："494f3031f10201-tcpsocket"
func (p *NodeDoValueObject)UpdateNodeDoMap(pn physicalnode.PhysicalNode){
	handler, tag :=pn.SelectHandlerAndTag()
	p.lock.Lock()
	for k,_ :=range p.NodeDoMap{
		if !strings. Contains(k,fmt.Sprintf("%s-%s",handler,tag)){
			continue
		}
		nodename :=strings.Split(k,"-")[2]

		pvalue,ptime := pn.SelectOneValueAndTime(handler, tag, nodename)
		p.NodeDoMap[k].UpdateOneNodeDo(pvalue,ptime)
	}
	p.lock.Unlock()
}

func (p *NodeDoValueObject)Quit(){
	p.done <- true
	close(p.done)
}