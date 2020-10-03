//这个接口本质是各个数据库实体(如OldMySQLNode)的超集，这些数据库实体也存在于当前目录下，如mysql_oldnode.go文件内

//获得已填充好数据的实体有3种方式:
//1.通过下面的工厂函数NewDBEntity()创建，返回值是DBEntity接口，其拥有所对应结构体(如OldMySQLNode)的所有功能
//2.直接去创建某个实体(如OldMySQLNode)
//3.通过dao创建，此方法使用了gorm.DB的依赖注入，创建实体后会从数据库立刻获取并装配数据
//前2种往往用于测试需求，拿到结构体或者超集接口后，需要立刻执行InitForHealthTest()等模拟数据内容的装配函数
//第3种用于实际应用场景，所生成的dao包含create方法，该方法需要传入表名(tablename)参数，实现从数据库调取数据
//表名诸如“494f3031f10201”，表名=tablename=节点名=nodename

//后期dbentity往往只充当测试用的数据源

//20200918目前看来，他有作为测试模块的潜质
//能想到的核心需求是，首先他需要被设计成一个数据源源，也就是数据流动的上游
//需要基于timer或者ticket设计个功能模块或方法，让他定时主动发送给下游数据
//当然这个功能模块也可以是一个独立的工具包，内部采用组合的编程思想
//与这个组合，实现需求

package dbentity


type DBEntity interface{
	InitForHealthTest()

	GetId() int
	GetNodeName() string
	GetValue() string
	GetTime() string
	GetIp() string
	GetAll() (int,string,string,string,string)
}

//这个工厂函数往往只用来实现测试需求
//n.InitForHealthTest_494f3031f10201()这一行可根据需求任意替换
//就不去单独设计个testMode参数了
func NewDBEntity(typeName string,/*testMode string*/) DBEntity {
	switch (typeName){
	case "OldMySQLNode":
		n :=new(OldMySQLNode)
		/*switch (testMode)*/
		n.InitForHealthTest_494f3031f10201()//根据需求任意替换
		return n
	default:
		return nil
	}
}