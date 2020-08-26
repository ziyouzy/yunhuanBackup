package node

import (
	"errors"
	"github.com/ziyouzy/mylib/yunhuanfactory/entity/impl"
)

/*
	获取完整体和从old_db结构体转换成entity结构体是两个不同的
	从old_db转换成entity是dao层的事务(dao结构体的方法，或独立的函数)
	获取完整体是model层某个entity自身的事务(model.entity某个结构体的方法)
*/

/*
	节点具有哈希属性
	一个节点名称，如494f3031f10201，会对应一个具体的node结构体序列化后所得的字符串
	而IEntity是不同设备node结构体的统一接口
*/

type IEntity interface{
	//需实现方法:获取节点完整体
	//需实现方法:序列化结构体并返回序列化后的字符串
}