package conf

import(
	"fmt"
	"sync"

	"github.com/ziyouzy/mylib/mysql"
	"github.com/ziyouzy/mylib/conf/myvipers"
	"github.com/ziyouzy/mylib/alarmbuilder"
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
)


//拿到可以全局使用的viper变量
func Load(){
	var lock sync.Mutex
	//SingleViper是文件级的，个体拥有独立的chan bool管道，从而告诉上级json文档是否发生更新
	//也就是说一个文件对应一个configischange的管道，因此在这里就可以实现点对点的触发机制
	mysql.ConnectMySQL("yunhuan_api:13131313@tcp(127.0.0.1:3306)/yh?charset=utf8")
	lock.Lock()
	//myvipers可以独立的去自我实现更新
	//Load所返回的管道是个独立的管道，实现了每个SingleViper的扇入汇总
	Confofwidgets_testIsChange := myvipers.Load(/*,/abc/def/ghi.json*/"./widgetsonlyserver.json")

	nodedobuilder.LoadSingletonPattern(1, myvipers.SelectOneMap("./widgetsonlyserver.json", "test_mainwidget.nodes"))
	alarmbuilder.LoadSingletonPattern(myvipers.SelectOneMap("./widgetsonlyserver.json", "test_mainwidget.alarms.tty1-serial"))
	lock.Unlock()
	fmt.Println("初始化了nodedobuilder与alarmbuilder的单例模式")

	go func(){
		for{
			select{
			case <-Confofwidgets_testIsChange:
				lock.Lock()
				nodedobuilder.LoadSingletonPattern(1, myvipers.SelectOneMap("./widgetsonlyserver.json", "test_mainwidget.nodes"))
				alarmbuilder.LoadSingletonPattern(myvipers.SelectOneMap("./widgetsonlyserver.json", "test_mainwidget.alarms.tty1-serial"))
				lock.Unlock()
				fmt.Println("更新了nodedobuilder与alarmbuilder的单例模式")
			}
		}
	}()
}


