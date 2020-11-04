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
	nd.FlushTicket =time.NewTicker(step * time.Second)
	return nd
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
			case <-done:
				break
			case <-p.FlushTicket.C:
				p.lock.Lock()
				for k,v := range p.NodeDoMap{
					nodeDoCh <-v
				}
				p.lock.UnLock()
			}
		}    
	}()
	return nodeDoCh
}

//NodeDoMap.Key 举例: "494f3031f10201-tcpsocket-do3-bool"
//GetHandlerTagForConfNodeMap()返回值举例："494f3031f10201-tcpsocket"
func (p *NodeDoValueObject)UpdateNodeDoMap(pn physicalnode.PhysicalNode){
	pnhandlerandtag :=pn.GetHandlerTagForConfNodeMap()
	p.lock.Lock()
	for k,v :=range NodeDoMap{
		if !strings. Contains(k,pnhandlerandtag){
			continue
		}

		//原生工具只能实现一一分配，否则你需要自己封装一下相关功能
		tempstr	:= strings.Split(k,"-")
		handler :=tempstr[0]
		tag :=tempstr[1]
		nodename :=tempstr[2]

		//NodeDo所对应的json字符串的字段名称已包含了对应物理节点的完整信息
		//这里的实质是json配置文档的核心字段（上面的“handler”、“tag”、“nodename”）
		//与现有的physicalnode层实现对接，提取对应的节点数据。并不存在与自定义协议层(package protocol)的任何联系
		//最终会用所获得的节点数据实现NodeDo层
		pvalue,ptime := pn.SeleteOneValueAndTime(handler, tag, nodename)
		//使用从物理节点获取的信息渲染对应的NodeDo缓存，实现NodeDo层
		v.CountPhysicalNode(pvalue,ptime)
	}
	p.lock.UnLock()
}

func (p *NodeDoValueObject)Quit(){
	p.done <- true
	close(p.done)
}