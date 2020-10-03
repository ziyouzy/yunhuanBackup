package view

type Matrix struct{
	systems [][]byte
} 

func (p *Matrix)AppendSystem(system []byte){
	append(systems,system)
}

func (p *Matrix)Reset(){
	p.Name =nil
	p.System =nil
	p.Nodes =make([][]byte)
}