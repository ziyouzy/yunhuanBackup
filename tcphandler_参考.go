
import (
	"github.com/ziyouzy/mylib/tcp"
)
type tcpHandler struct{
	//tcp与udp共用的数据流入管道，需要在整体程序初始化前绑定
	//也就是说，初始化的函数要做两件事，并且先后顺序不能错:
	//1.绑定管道
	//2.listen客户端，并进行之后的handle操作

	//这个是需要与外界绑定的管道
	//但是目前看来他应该是个流入physicalnode接口的管道
	//而不是个纯粹的Recv管道
	//不过我想把这里设计成装饰者模式的核心
	//遇到的第一个问题是每个客户端都有自己的RecvCh是无可厚非的事，但是一个tcpHandler只有一个physicalnode做所有客户端的扇入是否可行呢？
	//第二个问题是，physicalnode是否可以作为装饰者模式的入口关键点，不过似乎如果第一个问题可行，也就没必要在tcphandler内设计装饰者模式了，因为目前没有需求
	//最后一个问题，目前看来似乎还是需要在tcpClinet内部设计装饰者模式的，要不然还有更好的办法往physicalnode里存数据吗？
	//或者说在tcpHandler里并不需要存在RecvCh chan([]byte)或physicalnode接口的Ch
	//只要让tcpHandler取实现tcpclient的装饰者模式即可
	/*RecvCh chan([]byte)*/

	//每当配置新客户端连接时将管道指针依赖注入从而实现上下游通信
	RecvCh chan([]byte)//似乎这个也不需要，而是同样作为返回值
	//ip地址和结构体指针的哈系表
	ConnMap map[string]*tcp.PipelineTcpClient
}

func (p *tcpHandler)bindRecvCh(ch PhysicalNodeCh){
	p.PhysicalNodeCh =ch
}

/*先绑定管道后监听*/
func (p *tcpHandler)BindChanToListenAndRead(ch PhysicalNodeCh, phych){
	p.ConnMap =make(map[string]*tcpClient)//初始化map
	p.bindRecvCh(ch)//绑定

	//监听
	tcpAddr, err := net.ResolveTCPAddr("tcp", ":6668")
	if err != nil {
		fmt.Println(err.Error())
	}
	listener, err := net.ListenTCP("tcp", tcpAddr)
	if err != nil {
		fmt.Println(err.Error())
	}	
	defer listener.Close()
	fmt.Println("tcpHandler now is listening")

	//开始接收数据
	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println(err.Error())
		}
		ip,tcpconn :=newTcpClient(conn/*, p.RecvCh*/)
		//因为是组合，所以必须先实例化一个tcpclient实体
		//然后将这个tcpclient组合为tcpClientWithPhysicalNodeCh，此时不打开tcpclient内部管道，而是先执行StartIntoPhysicalNodeCh()让内外管道联结起来
		//StartIntoPhysicalNodeCh()内部会单独开一个线程for rangeRecvCh,实现数据流动
		//然后装入map
		//最后执行p.ConnMap[ip].tcpClient.Start()
		//问题是真的决定不在最内层tcpclient直接依赖注入PhysicalNodeCh的指针从而跳过RecvCh直接Into核心功能管道呢？
		//似乎不太可行，因为会修改并破坏tcpsocket底层包的代码了，继承不会破坏，组合更不会，他们都是为了完全隔离对tcpsocket底层包的修改
		//其实挺不错的，发现一个问题，那就是什么情况下值的做一次数据流动，目前看来，包与包之间是值得的，这里是tcpsocket包与当前整体项目包
		//而理由是1.是包与包之间；2.底层包与外部通信方式是基于管道的数据流动，如果想与其对接，必然也要用一个管道，目前看来方式很优雅
		tcpclientdec :={tcpclient: tcpconn,PhysicalNodeCh:phych}
		p.ConnMap[ip] =tcpclientdec
		p.ConnMap[ip].Start("RECEIVE&SEND")

		fmt.Println("收到客户端的连接，客户端ip为：",ip)
		fmt.Println("用打印ip地址的方式测试ConnMap是否正常(不带*):",p.ConnMap[ip].Conn.RemoteAddr())
		fmt.Println("用打印ip地址的方式测试ConnMap是否正常(带*):",(*p.ConnMap[ip]).Conn.RemoteAddr())
	}
}





type tcpClientWithPhysicalNodeCh struct{
	tcpclient TcpClient
	PhysicalNodeCh chan (PhysicalNode)
}

func (p *tcpClientWithPhysicalNodeCh)Set(t tcpClient,  physicalnodech chan (PhysicalNode)){
	p.tcpclient = t
	p.PhysicalNodeCh =physicalnodech
}


func (p *tcpClientWithPhysicalNodeCh)StartIntoPhysicalNodeCh {
	//有两种方式：
	//第一种是基于内部tcpclient重写readThread()函数(这会让内部RecvCh名存实亡)

	//第二种是先调用readThread函数，这会让内部RecvCh先流入数据
	//之后在这里用for-range循环获取数据(当然了以可以先准备好其他参数的配置):

	clinetIP :=p.tcpclient.Conn.getip.String()//(会变成tcpclientdec)这里就不能实现了把，以为这里会变成调用接口
	go for tempbytes := p.tcpclient.RecvCh{//(会变成tcpclientdec)这里也不能直接轮训接口内部的对象了
		tempPhyNode := NewPhysicalNodeFromTcpSocketBytes(clinetIP, tempbytes) 
		PhysicalNodeCh <-&tempPhyNode
	} 
}




