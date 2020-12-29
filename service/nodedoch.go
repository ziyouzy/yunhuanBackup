package service

import(
	"github.com/ziyouzy/mylib/nodedobuilder"
	"github.com/ziyouzy/mylib/nodedo"
	"github.com/ziyouzy/mylib/physicalnode"
)

var (
	NodeDoCh chan nodedo.NodeDo
)





func BuildNodeDoCh(){
	nodedobuilder.StartEngine(physicalnode.PhysicalNodeCh,nil)

	nodedobuilder.GenerateNodeDoCh()
	if NodeDoCh ==nil { NodeDoCh = make(chan nodedo.NodeDo) }

	ch :=nodedobuilder.GetNodeDoCh()//内层自动关闭
	go func(){
		for nodedo := range ch {
			/*每个nodedo在上层都已经实现了对超时的判定工作*/
			NodeDoCh<-nodedo
		}
	}()
}

func DestoryNodeDoCh(){
	if NodeDoCh !=nil { close(NodeDoCh) }
}