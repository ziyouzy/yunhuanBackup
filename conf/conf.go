package conf

import(
	//"github.com/spf13/viper"
	//"github.com/fsnotify/fsnotify"

	"fmt"
	//"strings"
	//"encoding/json"
	"github.com/ziyouzy/mylib/conf/viperlistener"

	//"os"
	"github.com/ziyouzy/mylib/do"
	"github.com/ziyouzy/mylib/alarm"
)

//使用type只是用来为其设计update方法
//无论是NodeDoVO还是AlarmVO都是一个不能直接使用的原型
//之后的使用方式是使用类似依赖注入的方式作为上层结构体对象的“引擎”
//这里也是体现了分层的设计思路，借鉴与最初tcp/ip协议的设计思路
//也就是如果不分层而把所有功能都“压”在一起，设计起来就太复杂了
//type ConfValueObjectMap map[string]interface{}

var(
	NodeDoCache do.NodeDoValueObject
	AlarmFilterCache alarm.AlarmFilterObject
)

//拿到可以全局使用的viper变量
func Load(){
	viperlistener.LoadViper()

	nodeDoConf :=viperlistener.NewConfValueObjectMapByType("test_mainwidget.nodes")
	alarmFilterConf :=viperlistener.NewConfValueObjectMapByType("test_mainwidget.alarm")

	NodeDoCache :=do.NewNodeDoValueObj(3, do.NewNodeDoValueObjectMap(nodeDoConf))
	AlarmFilterCache :=alarm.NewAlarmFilterObject(alarmFilterConf)

	fmt.Println("初始化NodeDoCache成功,alarmFilterConf:",NodeDoCache)
	fmt.Println("初始化AlarmFilterCache成功,alarmFilterConf:",AlarmFilterCache)

	//通过viper监听配置文件是否被改动:
	go func(){
		for{
			select {
			case <-viperlistener.ConfigIsChange:
				nodeDoConf :=viperlistener.NewConfValueObjectMapByType("nodedo")
				alarmFilterConf :=viperlistener.NewConfValueObjectMapByType("alarm")

				NodeDoCache :=do.NewNodeDoValueObj(3, do.NewNodeDoValueObjectMap(nodeDoConf))
				AlarmFilterCache :=alarm.NewAlarmFilterObject(alarmFilterConf)

				fmt.Println("更新NodeDoCache成功,alarmFilterConf:",NodeDoCache)
				fmt.Println("更新AlarmFilterCache成功,alarmFilterConf:",AlarmFilterCache)
			}
		}
	}()
}


