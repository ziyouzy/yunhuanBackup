package service

import(
	"github.com/ziyouzy/mylib/physicalnode"
	"github.com/ziyouzy/mylib/nodedocontroller"
	"github.com/ziyouzy/mylib/nodedo"
)

func NodeDoCh()chan nodedo.NodeDo{return nodedocontroller.GenerateNodeDoCh()}

func UpdateEveryExsitNodeDoTemplate(pnch chan physicalnode.PhysicalNode){
	go func(){
		for ph := range pnch{
			nodedocontroller.Engineing(ph)
		}
	}()
}