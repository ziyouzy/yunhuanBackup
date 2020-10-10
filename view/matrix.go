package view

type Matrix struct{
	MatrixName string
	Systems []System
} 

func (p *Matrix)AppendSystem(system System){
	p.Systems =append(p.Systems,system)
	p.MatrixName =system.MatrixName
}

func (p *Matrix)Reset(){
	p.MatrixName =""
	p.Systems =nil
}