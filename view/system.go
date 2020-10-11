package view

type System struct{
	SystemName string
	MatrixName string
	
	Modules []Module
	ModulesLen int
}

func (p *System)AppendModule(module Module){
	p.Modules =append(p.Modules,module)
	p.SystemName =module.SystemName
	p.MatrixName =module.MatrixName

	p.ModulesLen =p.ModulesLen+1
}

func (p *System)Reset(){
	p.SystemName =""
	p.MatrixName =""
	p.Modules =nil
	p.ModulesLen = 0
}