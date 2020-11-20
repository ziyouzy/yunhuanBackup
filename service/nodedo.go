package service

import(
	"github.com/ziyouzy/mylib/physicalnode"
	"github.com/ziyouzy/mylib/nodedocontroller"
	"github.com/ziyouzy/mylib/nodedo"
)


//PhysicalNodeCh的子携程消费者
func UpdateEveryExsitNodeDoTemplate(pnch chan physicalnode.PhysicalNode){
	go func(){
		for ph := range pnch{
			nodedocontroller.Engineing(ph)
		}
	}()
}

//NodeDoCh
//下层已经实现了子携程生产者，上层去实现子携程消费者
func NodeDoCh()chan nodedo.NodeDo{return nodedocontroller.GenerateNodeDoCh()}

