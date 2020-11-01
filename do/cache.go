//把缓存单独提取出来了，有些像是属于主函数的那几个handler(tcphandler,udphandler等等)
package do

type NodeDoCacheObject struct{
	NodeDoMap map[string]*NodeDo
	FlushTicket *time.Ticker
	lock *sync.Mutex
}

func (p *NodeDoCacheObject)Load(step int,m map[string]*NodeDo){
	p.NodeDoMap =m
	p.FlushTicket =time.NewTicker(step * time.Second)
}

//NodeDoMap.Key 举例: "494f3031f10201-tcpsocket-do3-bool"
//GetHandlerTagForConfNodeMap()返回值举例："494f3031f10201-tcpsocket"
func (p *NodeDoCacheObject)UpdateNodeDoMap(pn physicalnode.PhysicalNode){
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
		pvalue,ptime := pn.SeleteOneValueByProtocol(handler, tag, nodename)
		//使用从物理节点获取的信息渲染对应的NodeDo缓存
		v.CountPhysicalNode(pvalue,ptime)
	}
	p.lock.UnLock()
}

//结合定时器生成可供tcpsocket直接发送的字节管道
func (p *NodeDoCacheObject)CreateBytesCh()chan []byte{
	nodeDoBytesCh := make(chan []byte)
	go func(){
		for {
			select {
			case <-p.FlushTicket.C:
				p.lock.Lock()
				for k,v := range p.NodeDoMap{
					nodeDoBytesCh <-v.GetJson()
					p.NodeDoMap[k].IsTimeOut =true
				}
				p.lock.UnLock()
			}
		}    
	}()
}