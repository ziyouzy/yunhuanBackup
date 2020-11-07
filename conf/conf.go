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

//这个函数似乎不该属于这一层,这个要属于conf层
func NewConfValueObjectMap(path string,key string)map[string]interface{}{
	m :=viperlistener.Vipers[path].V.Get(key)

	if value, ok :=m.(map[string]interface{});ok{
		return value
	}else{
		fmt.Println("CreateConfValueObjectMap fail, path is:", path,"key is:",key)
		return nil
	}
}

//拿到可以全局使用的viper变量
func Load(){
	viperlistener.Load("./widgets_test.json"/*,/abc/def/ghi.json*/)
	NodeDoConf :=NewConfValueObjectMap("./widgets_test.json","test_mainwidget.nodes")
	AlarmFilterConf :=NewConfValueObjectMap("./widgets_test.json","test_mainwidget.alarm")

	NodeDoController :=nodedocontroller.BuildNodeDoController(3, do.NewNodeDoValueObjectMap(nodeDoConf))
	AlarmController :=alarmcontroller.BuildAlarmController(alarmFilterConf)

	fmt.Println("初始化NodeDoCache成功,alarmFilterConf:",NodeDoCache)
	fmt.Println("初始化AlarmFilterCache成功,alarmFilterConf:",AlarmFilterCache)
}


