package physicalnode

import (
	"fmt"
	"github.com/ziyouzy/mylib/yunhuanfactory/physicalnode/evolver"
)

type YouRenDO struct{
	NodeType string
	Evolver evolver.Evolver
	
	DO_id int	//其所处数据库内所在表对应行的id号
	DO_name string	//唯一标识，很重要，之后很多功能都需要通过他来实现
	DO_input_time string
	DO_value string
	DO_ip string

	DO1 string
	DO2 string
	DO3 string
	DO4 string
	DO5 string
	DO6 string
	DO7 string
	DO8 string
}

func (p *YouRenDO)Evolve(){
	err :=p.Evolver.Evolve("DO1",&p.DO1)
	fmt.Println(err)
	err =p.Evolver.Evolve("DO2",&p.DO2)
	fmt.Println(err)
	err =p.Evolver.Evolve("DO3",&p.DO3)
	fmt.Println(err)
	err =p.Evolver.Evolve("DO4",&p.DO4)
	fmt.Println(err)
	err =p.Evolver.Evolve("DO5",&p.DO5)
	fmt.Println(err)
	err =p.Evolver.Evolve("DO6",&p.DO6)
	fmt.Println(err)
	err =p.Evolver.Evolve("DO7",&p.DO7)
	fmt.Println(err)
	err =p.Evolver.Evolve("DO8",&p.DO8)
	fmt.Println(err)
}

func (p *YouRenDO)GetNodeType() string{
	return p.NodeType
}