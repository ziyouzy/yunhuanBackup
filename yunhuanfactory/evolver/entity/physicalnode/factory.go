package physicalnode

import(
	allnodesevolverlv1 "github.com/ziyouzy/mylib/yunhuanfactory/evolver/entity/physicalnode/lv1/impl"
)

func NewPhysicalNodeEntityEvolver(edition string) IAllNodesEvolver{
	switch (edition){
		case "lv1":
			evolver :=new(allnodesevolverlv1.AllNodesEvolverImpl)
			return evolver
		default:
			return nil
	}
}