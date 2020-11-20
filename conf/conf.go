package conf

import(
	//"github.com/spf13/viper"
	//"github.com/fsnotify/fsnotify"

	"fmt"
	"sync"
	//"strings"
	//"encoding/json"
	//"os"
	"github.com/ziyouzy/mylib/conf/myvipers"
	"github.com/ziyouzy/mylib/alarmcontroller"
	"github.com/ziyouzy/mylib/nodedocontroller"
)

//使用type只是用来为其设计update方法
//无论是NodeDoVO还是AlarmVO都是一个不能直接使用的原型
//之后的使用方式是使用类似依赖注入的方式作为上层结构体对象的“引擎”
//这里也是体现了分层的设计思路，借鉴与最初tcp/ip协议的设计思路
//也就是如果不分层而把所有功能都“压”在一起，设计起来就太复杂了
//type ConfValueObjectMap map[string]interface{}

// var(
// 	NodeDoCache do.NodeDoValueObject
// 	AlarmFilterCache alarm.AlarmFilterObject
// )

//这个函数似乎不该属于这一层,这个要属于conf层
//func NewConfValueObjectMapOREngine(path string,key string)map[string]interface{}{
	//m :=myvipers.Vipers[path].V.Get(key)
//	m :=myvipers.SelectOne(path).Get(key)

	// if value, ok :=m.(map[string]interface{});ok{
	// 	return value
	// }else{
	// 	fmt.Println("NewConfValueObjectMapOREngine fail, path is:", path,"key is:",key)
	// 	return nil
	// }
//}

//拿到可以全局使用的viper变量
func Load(){
	var lock sync.Mutex
	//SingleViper是文件级的，也就是说一个文件对应一个configischange的管道，因此在这里就可以实现点对点的触发机制
	Confofwidgets_testIsChange :=make(chan bool) 

	lock.Lock()
	myvipers.Load(Confofwidgets_testIsChange,/*,/abc/def/ghi.json*/"./widgetsonlyserver.json")
	nodedocontroller.LoadSingletonPattern(5, myvipers.SelectOneMap("./widgetsonlyserver.json", "test_mainwidget.nodes"))
	alarmcontroller.LoadSingletonPattern(myvipers.SelectOneMap("./widgetsonlyserver.json", "test_mainwidget.alarms.tty1-serial"))
	lock.Unlock()
	fmt.Println("初始化了nodedocontroller与alarmcontroller的单例模式")
	go func(){
		for{
			select{
			case <-Confofwidgets_testIsChange:
				//Confofwidgets_testIsChange :=make(chan bool) 

				lock.Lock()
				//myvipers.Load(Confofwidgets_testIsChange,/*,/abc/def/ghi.json*/"./widgetsonlyserver.json")
				nodedocontroller.LoadSingletonPattern(5, myvipers.SelectOneMap("./widgetsonlyserver.json", "test_mainwidget.nodes"))
				alarmcontroller.LoadSingletonPattern(myvipers.SelectOneMap("./widgetsonlyserver.json", "test_mainwidget.alarms.tty1-serial"))
				lock.Unlock()
				fmt.Println("更新了nodedocontroller与alarmcontroller的单例模式")
			}
		}
	}()
}


