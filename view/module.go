package view

type Module struct{
	Name string
	System string
	Nodes [][]byte
}

func (p *Module)AppendNode(node []byte){
	append(nodes,node)
}

func (p *Module)Reset(){
	p.Name =nil
	p.System =nil
	p.Nodes =make([][]byte)
}
