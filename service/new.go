package service

import(
	"sync"
	"fmt"

	"github.com/ziyouzy/mylib/viperbuilder"
	"github.com/ziyouzy/mylib/nodedobuilder"
	"github.com/ziyouzy/mylib/nodedo"
	"github.com/ziyouzy/mylib/physicalnode"
	"github.com/ziyouzy/mylib/connserver"
	"github.com/ziyouzy/mylib/alarmbuilder"
)

var (
	NodeDoCh chan nodedo.NodeDo
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

func BuildPNCh(){
	rawCh :=connserver.RawCh()
	physicalnode.RawChToPhysicalNodeCh(rawCh)
}

func BuildNodeDoCh(){
	nodedobuilder.StartEngine(physicalnode.PhysicalNodeCh,nil)

	nodedobuilder.GenerateNodeDoCh()
	if NodeDoCh ==nil { NodeDoCh = make(chan nodedo.NodeDo) }

	ch :=nodedobuilder.GetNodeDoCh()//内层自动关闭
	go func(){
		for nodedo := range ch {
			/*每个nodedo在上层都已经实现了对超时的判定工作*/
			NodeDoCh<-nodedo
		}
	}()
}

func DestoryNodeDoCh(){
	if NodeDoCh !=nil { close(NodeDoCh) }
}