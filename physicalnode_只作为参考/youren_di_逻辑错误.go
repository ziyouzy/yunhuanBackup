//此源文件包含YouRenDI实体的结构体
//该结构体所实现的接口存在于同级目录下的physicalnode.go源文件
//他需要调用evolver包内的功能对象从而实现自身的数据更新，让数据拥有真正完整的价值与意义
//他不需要pyhsicalnode/dao包，相反的，他需要dao包引入从而是实现物理节点的实体化
package physicalnode

import(
	"fmt"
	"github.com/ziyouzy/mylib/physicalnode/evolver"
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


//不用在这里合成IO实体，后期在说
//这里只专注每一条从上游tcp流过来的数据如何处理
//这里错了，应该是基于接口的工厂函数
func (p *YouRenDI)CreatePhysicalNodeFromTcpBytes(b []byte,e evolver.Evolver) PhysicalNode{
	p.DI_id =-1
	p.DI_name //之后整理
	p.DI_inputTime//之后整理
	p.DI_value =string（b）
	p.DI_ip =ip

	p.Evolver = e
}

//获取一个初始的entity实体
func (p *YouRenDI)toFull() PhysicalNode{
	//var di physicalnode.YouRenDI

	p.NodeType  =p.NodeType
	//p.Evolver = evolver
	p.DI_id, p.DI_name, p.DI_value, p.DI_inputTime, p.DI_ip =p.dbEntity.GetAll()

	tempStr :=strings.Split(di.DI_value,"|")[0]
	if strings.Contains(tempStr, "timeout"){
		di.DI8 ="timeout"
		di.DI7 ="timeout"
		di.DI6 ="timeout"
		di.DI5 ="timeout"
		di.DI4 ="timeout"
		di.DI3 ="timeout"
		di.DI2 ="timeout"
		di.DI1 ="timeout"
	}else if strings.Index(tempStr,"494f")==0{
		c :=[]byte(tempStr)
		tempStr =string([]byte(c)[12:16])
		hex, _:=strconv.ParseInt(tempStr,16,0)
		tempStr =strconv.FormatInt(hex,2)
		tempStr =string([]byte(tempStr)[1:])

		di.DI8 =string([]byte(tempStr)[0:1])//“”的内容都是读取配置文档获得
		di.DI7 =string([]byte(tempStr)[1:2])//如这里的di7，this.Di7是string类型，这里给的值“1”也是string类型，调用warningMgr包中对应的函数进行加工
		di.DI6 =string([]byte(tempStr)[2:3])
		di.DI5 =/*"市电断电(0为开路;1为闭路)|"+*/string([]byte(tempStr)[3:4])
		di.DI4 =string([]byte(tempStr)[4:5])
		di.DI3 =string([]byte(tempStr)[5:6])
		di.DI2 =/*"顶盖状态(0为开路;1为闭路)|"+*/string([]byte(tempStr)[6:7])
		di.DI1 =/*"顶盖全打开(0为开路;1为闭路)|"+*/string([]byte(tempStr)[7:8])
	}else{
		di.DI8 = "err"
		di.DI7 = "err" 
		di.DI6 = "err"
		di.DI5 = "err"
		di.DI4 = "err"
		di.DI3 = "err"
		di.DI2 = "err"
		di.DI1 = "err"
	}

	return &di
}

func (p *YouRenDI)toEvolve(){
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

//获取一个初始的entity实体
func (p *YouRenDI)GetThisPhysicalNode() PhysicalNode{
	p.ToFull(evolver)
	p.ToEvolve()
	return node
}



func (p *YouRenDI)GetNodeType() string{
	return p.NodeType
}
