// 详细介绍请看yunhuanfactory/physicalnode/dao/physicalnode_yourendodao.go
// 似乎不需要依赖注入yunhuanfactory/physicalnode/evolver.Evolver
// 因为每个独立的dbEntity也被依赖注入了evolver evolver.Evolver

//其并不会包含evolver的指针，而只是会作为一个传递媒介，将evolver指针通过方法直接依赖注入最终的物理节点实体
package dao

import (
	"fmt"
	//"errors"
	"strings"

	"github.com/ziyouzy/mylib/yunhuanfactory/dbentity"

	"github.com/ziyouzy/mylib/yunhuanfactory/physicalnode"
	"github.com/ziyouzy/mylib/yunhuanfactory/physicalnode/evolver"
)

type YouRenIODao struct{
	NodeType string
	Edition string

	dbEntityArr []dbentity.DBEntity
}

//获取一个初始的entity实体
func (p *YouRenIODao)CreatePhysicalNode(evolver evolver.Evolver) physicalnode.PhysicalNode{
	var io physicalnode.YouRenIO 
	for _,value :=range p.dbEntityArr{
		fmt.Println(value.GetNodeName())
		if strings.LastIndex(value.GetNodeName() ,"f10201")==8{
			didao := YouRenDIDao{
				dbEntity: value,
			}
			io.DI=didao.CreatePhysicalNode(evolver)
		}

		if strings.LastIndex(value.GetNodeName() ,"f10101")==8{
			dodao := YouRenDODao{
				dbEntity: value,
			}
			io.DO=dodao.CreatePhysicalNode(evolver)
		}
	}
	return &io
}

//这里需要传入evolver evolver.Evolver是因为需要对内部两个子物理节点实体进行依赖注入
func (p *YouRenIODao)CreateEvolvedPhysicalNode(evolver evolver.Evolver) physicalnode.PhysicalNode{
	node :=p.CreatePhysicalNode(evolver)
	node.Evolve()
	return node
}