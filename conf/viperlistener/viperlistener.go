package viperlistener

import(
	"github.com/spf13/viper"
	"github.com/fsnotify/fsnotify"

	"fmt"
	//"strings"
	//"encoding/json"
	//"github.com/ziyouzy/mylib/physicalnode"
	//"github.com/mitchellh/mapstructure"

	//"os"
	//"github.com/ziyouzy/mylib/do"
	//"github.com/ziyouzy/mylib/alarm"
)

//使用type只是用来为其设计update方法
//无论是NodeDoVO还是AlarmVO都是一个不能直接使用的原型
//之后的使用方式是使用类似依赖注入的方式作为上层结构体对象的“引擎”
//这里也是体现了分层的设计思路，借鉴与最初tcp/ip协议的设计思路
//也就是如果不分层而把所有功能都“压”在一起，设计起来就太复杂了
//type ConfValueObjectMap map[string]interface{}

var(
	//NodeDoVO ConfValueObjectMap
	//AlarmVO ConfValueObjectMap

	//NodeDoCache do.NodeDoValueObject
	//AlarmFilterCache alarm.AlarmFilterObject
	ConfigIsChange chan bool
)

func InitConfMap(){
	viper.SetConfigName("riverconf") //  设置配置文件名 (不带后缀)
	//viper.AddConfigPath("/workspace/appName/") 
	viper.AddConfigPath(".")               // 比如添加当前目
	viper.SetConfigType("json")
	err := viper.ReadInConfig() // 搜索路径，并读取配置数据

	if err == nil {
		// NodeDoVO :=ConfValueObjectMap
		// NodeDoVO.update("nodes")
		// fmt.Println("NodeDoVO in init:",NodeDoVO)
		// NodeDoCache =do.NewNodeDoValueObj(3,NodeDoVO)

		// AlarmVO :=ConfValueObjectMap
		// AlarmVO.update("alarm")
		// fmt.Println("AlarmVO in init:", AlarmVO)
		// AlarmFilterCache =alarm.NewAlarmFilterObject(AlarmVO)

		go watching()
	}else{
		panic(fmt.Errorf("Fatal init config file! \n"))
	}
}

func NewConfValueObjectMapByType(typeString string)map[string]interface{}{
	m :=viper.Get(typeString)
	if value, ok :=m.(map[string]interface{});ok{
		return m
	}else{
		fmt.Println("CreateConfValueObjectMap fail, type is",typeString)
		return nil
	}
}

func watching() {
	viper.WatchConfig()
	viper.OnConfigChange(func(e fsnotify.Event) {
		fmt.Println("Config file changed:", e.Name)
		err := viper.ReadInConfig() // 搜索路径，并读取配置数据
		if err == nil {
		// 	NodeDoVO :=ConfValueObjectMap
		// 	AlarmVO :=ConfValueObjectMap
			
		// 	NodeDoVO.update("nodes")	
		// 	AlarmVO.update("alarm")

		// 	NodeDoCache.Quit()
		// 	NodeDoCache=NewNodeDoValueObj(3,NodeDoVO)

		// 	AlarmFilterCache.Quit()
		// 	AlarmFilterCache =NewAlarmFilterObject(AlarmVO)
			ConfigIsChange <-true
			return
		}else{
			fmt.Println("Fatal reset config file!")
			return
	}
	})
}


