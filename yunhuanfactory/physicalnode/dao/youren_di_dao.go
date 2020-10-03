//详细介绍请看yunhuanfactory/physicalnode/dao/physicalnode_yourendodao.go
//其并不会包含evolver的指针，而只是会作为一个传递媒介，将evolver指针通过方法直接依赖注入最终的物理节点实体
package dao

import (
	"strconv"
	"strings"

	"github.com/ziyouzy/mylib/yunhuanfactory/dbentity"

	"github.com/ziyouzy/mylib/yunhuanfactory/physicalnode"
	"github.com/ziyouzy/mylib/yunhuanfactory/physicalnode/evolver"
)

type YouRenDIDao struct{
	NodeType string
	Edition string

	//接口为引用类型，所以不需要“*”修饰符
	dbEntity dbentity.DBEntity
}

//获取一个初始的entity实体
func (p *YouRenDIDao)CreatePhysicalNode(evolver evolver.Evolver) physicalnode.PhysicalNode{
	var di physicalnode.YouRenDI

	di.NodeType  =p.NodeType
	di.Evolver = evolver
	di.DI_id, di.DI_name, di.DI_value, di.DI_inputTime, di.DI_ip =p.dbEntity.GetAll()

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

//获取一个初始的entity实体
func (p *YouRenDIDao)CreateEvolvedPhysicalNode(evolver evolver.Evolver) physicalnode.PhysicalNode{
	//这里创建的是个接口，所以就不能直接调用其内部字段了
	//如果想通过dao来Evolved其对应的物理实体，目前看来还是将Evolved实体的指针依赖注入到物理节点实体本身比较好
	//而dao本身并不需要拥有Evolved的指针，因为对他来说没什么用
	node :=p.CreatePhysicalNode(evolver)
	node.Evolve()
	return node
}