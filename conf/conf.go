package conf

import(
	"fmt"
	"sync"

	"github.com/ziyouzy/mylib/mysql"
	"github.com/ziyouzy/mylib/myvipers"
	"github.com/ziyouzy/mylib/alarmbuilder"

	"github.com/ziyouzy/mylib/nodedo"
	"github.com/ziyouzy/mylib/nodedobuilder"
)

var (
	TcpModbus1 = [][]byte{
		{0xf1,0x01,0x00,0x00,0x00,0x08,0x29,0x3c,},
		{0xf1,0x02,0x00,0x20,0x00,0x08,0x6c,0xf6,},
	}
	SnmpOids1 = [][]byte{
		[]byte("test1"),
		[]byte("test2"),
	}

	NodeDoCh = make(chan nodedo.NodeDo) 

	lock sync.Mutex
)


//拿到可以全局使用的viper变量
func Load(){
	//SingleViper是文件级的，个体拥有独立的chan bool管道，从而告诉上级json文档是否发生更新
	//也就是说一个文件对应一个configischange的管道，因此在这里就可以实现点对点的触发机制
	mysql.ConnectMySQL("yunhuan_api:13131313@tcp(127.0.0.1:3306)/yh?charset=utf8")
	connserver.LoadSingletonPatternListenAndCollect()


	
	lock.Lock()
	//myvipers可以独立的去自我实现更新
	//Load所返回的管道是个独立的管道，实现了每个SingleViper的扇入汇总
	Confofwidgets_testIsChange := myvipers.Load(/*,/abc/def/ghi.json*/"./widgetsonlyserver.json")
	nodedobuilder.LoadSingletonPattern(1, myvipers.SelectOneMap("./widgetsonlyserver.json", "test_mainwidget.nodes"))
	alarmbuilder.LoadSingletonPattern(myvipers.SelectOneMap("./widgetsonlyserver.json", "test_mainwidget.alarms.tty1-serial"))
	
	ch =nodedobuilder.GenerateNodeDoCh()
	go func(){
		//该上游会自动关闭
		for nodedo := range ch{
			NodeDoCh<-nodedo
		}
	}()
	lock.Unlock()
	fmt.Println("初始化了nodedobuilder与alarmbuilder的单例模式")

	Watching(Confofwidgets_testIsChange)
}

func Watching(Confofwidgets_testIsChange chan string){
	go func(){
		defer close(Confofwidgets_testIsChange)
		for changed := range Confofwidgets_testIsChange{
			switch changed{
			case "./widgetsonlyserver.json":
				lock.Lock()
				nodedobuilder.Quit()
				alarmbuilder.Quit()

				nodedobuilder.LoadSingletonPattern(1, myvipers.SelectOneMap("./widgetsonlyserver.json", "test_mainwidget.nodes"))
				alarmbuilder.LoadSingletonPattern(myvipers.SelectOneMap("./widgetsonlyserver.json", "test_mainwidget.alarms.tty1-serial"))

				//内层已对管道设计好了析构逻辑
				ch =nodedobuilder.GenerateNodeDoCh()
				go func(){
					//该上游会自动关闭
					for nodedo := range ch{
						NodeDoCh<-nodedo
					}
				}()
				lock.Unlock()
				fmt.Println("更新了nodedobuilder与alarmbuilder的单例模式")
			}
		}
	}()
}


