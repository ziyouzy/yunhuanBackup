package service

import(
	"sync"
	"fmt"

	"github.com/ziyouzy/mylib/viperbuilder"
	"github.com/ziyouzy/mylib/nodedobuilder"
	"github.com/ziyouzy/mylib/nodedo"
	"github.com/ziyouzy/mylib/physicalnode"
	"github.com/ziyouzy/mylib/connserver"
	//"github.com/ziyouzy/mylib/alarmbuilder"
)

var (
	NodeDoCh chan nodedo.NodeDo
	//PhysicalNodeCh chan physicalnode.PhysicalNode
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
				//alarmbuilder.Destory()

				nodedobuilder.Load(1, viperbuilder.SelectOneMapFromOneSingleViper("./widgetsonlyserver.json", "test_mainwidget.nodes"))
				//nodedobuilder.GenerateNodeDoCh()
				//nodedobuilder.Engineing(PhysicalNodeCh)
				//alarmbuilder.Load(viperbuilder.SelectOneMapFromOneSingleViper("./widgetsonlyserver.json", "test_mainwidget.alarms.tty1-serial"))

				BuildNodeDoCh()

				lock.Unlock()
				fmt.Println("更新了nodedobuilder与alarmbuilder的单例模式")
			}
		}
	}()
}

func BuildPNCh(){
	rawCh :=connserver.RawCh()
	physicalnode.RawChToPhysicalNodeCh(rawCh)
	// go func(){
	// 	for pn :=range physicalnode.PhysicalNodeCh{
	// 		fmt.Println("pn:",pn)
	// 		PhysicalNodeCh <-pn //rawCh的消费者和physicalNodeCh的创建者和生产者
	// 	}
	// }()
}

func BuildNodeDoCh(){
	nodedobuilder.GenerateNodeDoCh()
	nodedobuilder.Engineing(physicalnode.PhysicalNodeCh)

	if NodeDoCh ==nil { NodeDoCh = make(chan nodedo.NodeDo) }

	ch :=nodedobuilder.GetNodeDoCh()//内层自动关闭
	go func(){
		for nodedo := range ch {
			//NodeDoCh<-nodedo
			fmt.Println(string(nodedo.GetJson()))
		}
	}()

}

func DestoryNodeDoCh(){
	if NodeDoCh !=nil { close(NodeDoCh) }
}