/*目前的问题在于，我不可能每次evolve时都去硬盘调用一次配置文件*/
/*更确切的说，配置文件应该只有在程序启动时调用一次*/
/*所以说，无论如何，还是需要设计依赖注入*/
//不需要在最外层进行依赖注入，因为环境变量可以在程序初始化时加载
//但是evolver包还是首先家在到内存，设计成持久化
//或者说，将他和环境变量一同初始化，让他在整个程序的声明周期内只创建一次才更能提升运行效率
//但是他的确只是用来针对处理physicanode各实体内的字段值的
//所以他的文件目录暂时设计到yunhuanfactory/physicalnode/evolver下

//工厂函数包含一个版本号参数，以及一个type参数（type参数诸如PHYSICALNODE）

//他真的有必要独立设计成一个包吗，其实让他以evolver.go的源文件形式存在于yunhuanfactory/physicalnode目录下也是可以的
//这样可以让代码更加精简
//对比而言，alarm的不能这样进行设计，因为他们包含了两个独立的功能模块（serialconn与webcamera）

//20200916对于该包的线程安全问题:
//evolver实体目前不存在变量，也就更不会存在需要“资源共享的内容了”
//他只是从外部获取数据后立刻加工并返回罢了
//但是可以确定的是，他的声明周期必定只存在于某一个子携程中
//因此对于线程安全问题，应该存在于他的业务上游
//当从utp或tcpsocket获取数值，则会立刻开启携程
//这里会涉及到主线程的套接字获取对象于携程对象之间的资源共享
//显而易见的是套接字获取对象结构体内部应该用一个管道接收
//并将数据“流动”到子携程对象
//这里是个重点，子携程对象并不该用管道接收数据
//而是应该先加工数据后，将成品装入另一个管道
//这个先后顺序不要搞错，想想就觉得很完美，开心
package evolver


type Evolver interface{
	Evolve(string, *string) error
}

func NewEvolver(edition string, evolverType string) Evolver{
	switch (edition){
		case "lv1":
			switch (evolverType){
			case "PHYSICALNODE":
				evolver :=new(PhysicalNode)
				return evolver
			default:
				return nil
			}
		default:
			return nil
	}
}