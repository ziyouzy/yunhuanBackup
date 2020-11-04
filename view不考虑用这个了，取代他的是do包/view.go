//node，module，system，matrix每个级别都会有独立的管道与其对应
//虽然节点地位不同，但是管道的级别是平等的，他们的下游都是被websocket或socket直接发送函数，或者redis模块的直接录入函数
//此外告警sms、告警mysql录入这两个管道的下游分别是serial模块的直接发送函数以及mysql模块的直接录入函数
//上面所提到的管道地位都是平等的，并且都是数据流动环节的最后一个管道工序

/*node，module，system，matrix的具体使用方式*/
//他们4个是后者包含前者的关系，算是个默认的功能性模块
//如果客户需求是node级别，那就直接发送就好
//如果是system级别，那么就需要为module内部的module字段与system字段赋值，matrix可以不赋值
//换句话说，客户需求到了那个级别，则低于这一级别的结构体，以及字段名都需要完整的赋值
//这样才能保证数据发送过去后，前端知道每个结构体的内容要去与那个组件做匹配
//这道工序一般会在protocol相关的扇出函数内完成，这道工序是自定义协议的一部分，需要根据需求实时作好迭代更新

//各个结构体需要做到线程安全

package view

type View interface{
	AppendModule([]byte)
	Reset()
}