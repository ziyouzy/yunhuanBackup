//此源文件包含YouRenDI实体的结构体
//该结构体所实现的接口存在于同级目录下的physicalnode.go源文件
//他需要调用evolver包内的功能对象从而实现自身的数据更新，让数据拥有真正完整的价值与意义
//他不需要pyhsicalnode/dao包，相反的，他需要dao包引入从而是实现物理节点的实体化
package physicalnode

import(
	"fmt"
	"github.com/ziyouzy/mylib/yunhuanfactory/physicalnode/evolver"
)

type YouRenDI struct{
	NodeType string
	Evolver evolver.Evolver

	DI_id int	//其所处数据库内所在表对应行的id号
	DI_name string	//唯一标识，很重要，之后很多功能都需要通过他来实现
	DI_inputTime string
	DI_value string
	DI_ip string

	DI1 string 
	DI2 string
	DI3 string
	DI4 string
	DI5 string
	DI6 string
	DI7 string
	DI8 string
}

func (p *YouRenDI)Evolve(){
	//以后这些错误不会打印，而是写入日志
	err :=p.Evolver.Evolve("DI1",&p.DI1)
	fmt.Println(err)
	err =p.Evolver.Evolve("DI2",&p.DI2)
	fmt.Println(err)
	err =p.Evolver.Evolve("DI3",&p.DI3)
	fmt.Println(err)
	err =p.Evolver.Evolve("DI4",&p.DI4)
	fmt.Println(err)
	err =p.Evolver.Evolve("DI5",&p.DI5)
	fmt.Println(err)
	err =p.Evolver.Evolve("DI6",&p.DI6)
	fmt.Println(err)
	err =p.Evolver.Evolve("DI7",&p.DI7)
	fmt.Println(err)
	err =p.Evolver.Evolve("DI8",&p.DI8)
	fmt.Println(err)
}

func (p *YouRenDI)GetNodeType() string{
	return p.NodeType
}
