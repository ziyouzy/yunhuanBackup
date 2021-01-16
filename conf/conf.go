package conf

import(
	//"fmt"
	"sync"

	"github.com/ziyouzy/mylib/mysql"
	"github.com/ziyouzy/mylib/viperbuilder"
	"github.com/ziyouzy/mylib/alarmbuilder"
	"github.com/ziyouzy/mylib/connserver"
	"github.com/ziyouzy/mylib/nodedobuilder"
	"github.com/ziyouzy/mylib/physicalnode"
)

var lock sync.Mutex


//拿到可以全局使用的viper变量
func Load(){
	//SingleViper是文件级的，个体拥有独立的chan bool管道，从而告诉上级json文档是否发生更新
	//也就是说一个文件对应一个configischange的管道，因此在这里就可以实现点对点的触发机制
	mysql.ConnectMySQL("yunhuan_api:13131313@tcp(127.0.0.1:3306)/yh?charset=utf8")
	connserver.ListenAndCollect()
	physicalnode.Load()

	lock.Lock()
	//viperbuilder的核心是个map，每个value都是一个viper，每个viper都可以独立的去自我实现更新
	//但是需要监听另一个管道字段，从而实时了解哪个viper需要重置
	viperbuilder.Load(/*,/abc/def/ghi.json*/"./widgetsonlyserver.json")

	//第一个参数代表了每隔几秒发送给前端ui数据
	nodedobuilder.Load(1, viperbuilder.SelectOneMapFromOneSingleViper("./widgetsonlyserver.json", "test_mainwidget.nodes"))

	alarmbuilder.Load(viperbuilder.SelectOneMapFromOneSingleViper("./widgetsonlyserver.json", "test_mainwidget.alarms.tty1-serial"))
	alarmbuilder.GenerateSMSbyteCh()
	alarmbuilder.GenerateMYSQLAlarmCh()

	lock.Unlock()
}


