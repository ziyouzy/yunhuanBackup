
type System struct{
	Name string
	Modules [][]byte
}

func (p *System)AppendModule(module []byte){
	append(p.Modules,module)
}

func (p *System)Reset(){
	p.Name =nil
	p.System =nil
	p.Nodes =make([][]byte)
}