//虽然也是dao，但是只是基于dbEntity的反射器拿到数据，所以并不需要注入依赖于gorm.DB
//但是需要nodeType参数，从而确认以及标记获取哪种设备
//同dbentity/dao包一样，只有工厂函数来负责完整的对象创建流程，包括nodeType、edition的相关配置
//同样的，physicalnode本身的工厂函数禁止进行对版本的配置，这是为了实现高耦合低内聚思想
//nodeType需要与具体物理节点的结构体名称一致

//20200905
//现将工厂函数的第三个参数由"某个结构体"更改为"这个结构体所实现的接口"
//具体的说，是由dbentity.OldMySQLNode修改为dbentity.DBEntity这是为了让代码更具灵活性              
//输入一个接口，返回的也是一个接口 
//具体来说，输入一个db实体的接口，返回一个物理节点DAO的接口，这个返回的接口可以生成一个物理节点的接口    
package dao

import (
	"fmt"
	"errors"

	 "github.com/ziyouzy/mylib/yunhuanfactory/dbentity"

	 "github.com/ziyouzy/mylib/yunhuanfactory/physicalnode"
	 "github.com/ziyouzy/mylib/yunhuanfactory/physicalnode/evolver"
)
type PhysicalNodeDaoCreater interface{
	CreatePhysicalNode(evolver evolver.Evolver) physicalnode.PhysicalNode
	CreateEvolvedPhysicalNode(evolver evolver.Evolver) physicalnode.PhysicalNode
}


func NewPhysicalNodeDao(nodeType string, edition string,eArr ...dbentity.DBEntity) (PhysicalNodeDaoCreater  ,error){
	switch (edition){
	case "lv1":
		switch (nodeType){
		case "YouRenDO":
			if len(eArr) !=1{
				err :=errors.New(fmt.Sprintf("创建DoDao时，Old反射器实体不为1，而为%v",len(eArr)))
				return nil, err
			}

			//这里是physicalnode/dao的工厂函数，所以他的功能必然是生成dao的接口，因此
			//他必然使用到同级目录（physicalnode/dao/physicalnode_yourendodao.go）下的dao结构体类型
			//生成接口体后，return给同级目录的接口类型返回值
			dao := YouRenDODao{
				NodeType :nodeType,
				Edition:edition,
				dbEntity :eArr[0],
			}

			return &dao,nil

		case "YouRenDI":
			if len(eArr) !=1{
				err :=errors.New(fmt.Sprintf("创建DIDao时，Old反射器实体不为1，而为%v",len(eArr)))
				return nil, err
			}

			dao := YouRenDIDao{
				NodeType :nodeType,
				Edition:edition,
				dbEntity :eArr[0],
			}

			return &dao,nil

		case "YouRenIO":
			if len(eArr) !=2{
				err :=errors.New(fmt.Sprintf("创建DoDao时，Old反射器实体不为1，而为%v",len(eArr)))
				return nil,err
			}

			dao := YouRenIODao{
				NodeType :nodeType,
				Edition :edition,
				dbEntityArr :eArr,
			}

			return &dao,nil

		default:
			return nil,errors.New(fmt.Sprintf("创建DoDao时所时输入的%s是未知的nodetype类型",nodeType))
		}
	case "lv2":
		return nil,errors.New(fmt.Sprintf("创建DoDao时所时输入的%s是未定以的版本号",edition))
	default:
		return nil,errors.New(fmt.Sprintf("创建DoDao时所时输入的%s是未定以的版本号",edition))
	}
}