//设计该接口的目的是让版本更新时方便进行新旧版本的更新
//这个接口内的函数对负责版本更新者而言就好比说明书
//设计这个接口和设计实现了这个接口的结构体属于同一层的工作内容
//无论是lv几，以及无论是lv文件夹里的哪个model，都要满足这个接口


package physicalnode

import(
	//evo "github.com/ziyouzy/mylib/yunhuanfactory/evolver/entity/physicalnode"
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

type IYouRenEntity interface{
	//需实现方法:获取节点完整体
	//需实现方法:序列化结构体并返回序列化后的字符串
	Evolve()
}