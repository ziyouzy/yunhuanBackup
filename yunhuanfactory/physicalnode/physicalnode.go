//这个接口本质是各个物理节点实体(如YouRenDI)的超集，这些物理实体也存在于当前目录下，如yourenentity_di.go文件内

//思路上要清晰的区分：
//现在是在获取节点物理实体
//这与dbentity包中，负责操作数据库dao实体创建dbentity的数据库实体是两个不同的包，两件不同的事

//当前目录下的各个物理节点实体，需要基于当前目录下的子包（dao包）来实现
//暂时并没有设计工厂函数的需求，但是很可能立刻会发现很有必要

//每个物理节点实体都拥有GetNodeType()方法与Evolve()方法，该方法基于evolver子包，也因此，每个物理节点都实现了当前的PhysicalNode接口

//节点具有哈希属性
//一个节点名称，如494f3031f10201，会对应一个具体的node结构体序列化后所得的字符串
//而IEntity是不同设备node结构体的统一接口

package physicalnode

import (

)
//其只存在接口而不存在工厂函数是因为工厂函数往往只用来测试
//或者说测试的准备工作
//物理节点实体包需要基于的数据源是数据库实体包
//只要数据库实体包拥有工厂函数，能独立生成实体即可
type  PhysicalNode interface{
	//需实现方法:获取节点完整体
	//需实现方法:序列化结构体并返回序列化后的字符串
	Evolve()
	GetNodeType() string
}

// func NewPhysicalNodeFromTcpSocketBytes(b []byte) PhysicalNode{

// }