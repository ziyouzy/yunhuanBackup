// 似乎不需要依赖注入yunhuanfactory/physicalnode/evolver.Evolver
// 因为每个独立的dbEntity也被依赖注入了evolver evolver.Evolver

package physicalnode

import (
	"fmt"
)
type YouRenIO struct{
	NodeType string

	DI PhysicalNode
	DO PhysicalNode
}

func (p *YouRenIO)Evolve(){
	if p.DI !=nil{
		p.DI.Evolve()
	}else{
		fmt.Println("p.DI is nil")
	}

	if p.DO !=nil{
		p.DO.Evolve()
	}else{
		fmt.Println("p.DO is nil")
	}
}

func (p *YouRenIO)GetNodeType() string{
	return p.NodeType
}