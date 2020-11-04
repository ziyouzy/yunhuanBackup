package view

import(
	"github.com/ziyouzy/mylib/conf"
)

type Module struct{
	ModuleName string
	SystemName string
	MatrixName string

	ConfNodes []conf.ConfNode
	NodesLen int
}


func (p *Module)AppendNode(node conf.ConfNode){
	p.ConfNodes =append(p.ConfNodes,node)
	p.MatrixName, p.SystemName, p.ModuleName =node.GetMatrixSystemAndModuleString()
	p.NodesLen =p.NodesLen+1
}

func (p *Module)Reset(){
	p.ModuleName =""
	p.SystemName =""
	p.MatrixName =""
	p.ConfNodes =nil
	p.NodesLen =0
}
