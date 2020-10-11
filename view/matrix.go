package view

type Matrix struct{
	MatrixName string
	
	Systems []System
	SystemsLen int
} 

func (p *Matrix)AppendSystem(system System){
	p.Systems =append(p.Systems,system)
	p.MatrixName =system.MatrixName
	p.SystemsLen =p.SystemsLen+1
}

func (p *Matrix)Reset(){
	p.MatrixName =""
	p.Systems =nil
	p.SystemsLen =0
}