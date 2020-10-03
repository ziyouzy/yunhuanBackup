//当前源代码文件并没有引用到其同级目录下physicalnode_dao.go内的代码
//physicalnode_dao.go只包含当前源代码结构体YouRenDoDao所实现的接口，以及这一接口的工厂函数
//同样的，dbentity/dao/mysql_oldnodedao.go并没有用到dbentity/dao/mysql_dao.go内的元素
//这都是编写代码的思路逻辑和风格

//其并不会包含evolver的指针，而只是会作为一个传递媒介，将evolver指针通过方法直接依赖注入最终的物理节点实体
package dao

import (
	"strconv"
	"strings"
	"fmt"

	"github.com/ziyouzy/mylib/yunhuanfactory/dbentity"

	"github.com/ziyouzy/mylib/yunhuanfactory/physicalnode"
	"github.com/ziyouzy/mylib/yunhuanfactory/physicalnode/evolver"
)

type YouRenDODao struct{
	NodeType string
	Edition string

	//数据源，是一个接口类型，作为当前包所必须的内部字段，其属于dbentity包
	 dbEntity dbentity.DBEntity
}

func (p *YouRenDODao)CreatePhysicalNode(evolver evolver.Evolver) physicalnode.PhysicalNode{
	//创建物理节点结构体实体（非dao结构体实体），其和其所实现的方法都存在于上层目录
	//填充数据后，接收他的返回值是physicalnode.PhysicalNode这一接口类型
	//从而实现封装性
	var do physicalnode.YouRenDO

	do.NodeType  =p.NodeType
	do.Evolver = evolver
	do.DO_id, do.DO_name, do.DO_value, do.DO_input_time, do.DO_ip =p.dbEntity.GetAll()

	tempStr :=strings.Split(do.DO_value,"|")[0]
	if strings.Contains(tempStr, "timeout"){
		do.DO8 ="timeout"
		do.DO7 ="timeout"
		do.DO6 ="timeout"
		do.DO5 ="timeout"
		do.DO4 ="timeout"
		do.DO3 ="timeout"
		do.DO2 ="timeout"
		do.DO1 ="timeout"
	}else if strings.Index(tempStr,"494f")==0{
		c :=[]byte(tempStr)
		tempStr =string([]byte(c)[12:16])
		hex, _:=strconv.ParseInt(tempStr,16,0)
		tempStr =strconv.FormatInt(hex,2)
		tempStr =string([]byte(tempStr)[1:])

		do.DO8 =/*"顶盖恢复(0为开路;1为闭路)|"+*/string([]byte(tempStr)[0:1])
		do.DO7 =/*"顶盖开启(0为开路;1为闭路)|"+*/string([]byte(tempStr)[1:2])
		do.DO6 =/*"后门开关(0为开路;1为闭路)|"+*/string([]byte(tempStr)[2:3])
		do.DO5 =/*"前门开关(0为开路;1为闭路)|"+*/string([]byte(tempStr)[3:4])
		do.DO4 =/*"散热风扇(0为开路;1为闭路)|"+*/string([]byte(tempStr)[4:5])
		do.DO3 =/*"绿色灯带(0为开路;1为闭路)|"+*/string([]byte(tempStr)[5:6])
		do.DO2 =/*"红色灯带(0为开路;1为闭路)|"+*/string([]byte(tempStr)[6:7])
		do.DO1 =/*"蓝色灯带(0为开路;1为闭路)|"+*/string([]byte(tempStr)[7:8])
	}else{
		do.DO8 ="err"
		do.DO7 ="err"
		do.DO6 ="err"
		do.DO5 ="err"
		do.DO4 ="err"
		do.DO3 ="err"
		do.DO2 ="err"
		do.DO1 ="err"
		fmt.Println(tempStr)
	}

	return &do
}

func (p *YouRenDODao)CreateEvolvedPhysicalNode(evolver evolver.Evolver) physicalnode.PhysicalNode{
	node :=p.CreatePhysicalNode(evolver)
	node.Evolve()
	return node
}