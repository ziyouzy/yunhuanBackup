package service

import(
	"sync"
	"fmt"

	"github.com/ziyouzy/mylib/viperbuilder"
	"github.com/ziyouzy/mylib/nodedobuilder"
	"github.com/ziyouzy/mylib/alarmbuilder"
)


var (
	lock sync.Mutex
)

func WatchingViper(){
	pathch := viperbuilder.ConfigListener()
	go func(){
		for changed := range pathch{
			switch changed{
			case "./widgetsonlyserver.json":
				lock.Lock()
				nodedobuilder.Destory()
				alarmbuilder.Destory()

				nodedobuilder.Load(1, viperbuilder.SelectOneMapFromOneSingleViper("./widgetsonlyserver.json", "test_mainwidget.nodes"))
				BuildNodeDoCh()

				alarmbuilder.Load(viperbuilder.SelectOneMapFromOneSingleViper("./widgetsonlyserver.json", "test_mainwidget.alarms.tty1-serial"))
				alarmbuilder.GenerateSMSbyteCh()
				alarmbuilder.GenerateMYSQLAlarmCh()

				lock.Unlock()
				fmt.Println("更新了nodedobuilder与alarmbuilder的单例模式")
			}
		}
	}()
}