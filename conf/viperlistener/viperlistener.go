//viper是可以初始化多个的:
//testMainWidgetViper =viper.New()
//testMainWidgetViper.SetConfigName("widgets_test")等完成初始化
//之后就是使用了：
//和当前的区别在于，NewConfValueObjectMapByType()的参数表必须传入type和testMainWidgetViper作为依赖注入
//从而生成独立的缓存，供上层使用
//这个包需要重构，从viper过渡到vippers
//从而让他适应3个应用场景：
//1.存在多个.json文件
//2.一个viper对应一个矩阵级设备
//3.某个viper所在的json被改动时，立刻更新
package viperlistener

import(
	"github.com/spf13/viper"
	"github.com/fsnotify/fsnotify"

	"fmt"
)

//使用type只是用来为其设计update方法
//无论是NodeDoVO还是AlarmVO都是一个不能直接使用的原型
//之后的使用方式是使用类似依赖注入的方式作为上层结构体对象的“引擎”
//这里也是体现了分层的设计思路，借鉴与最初tcp/ip协议的设计思路
//也就是如果不分层而把所有功能都“压”在一起，设计起来就太复杂了
//type ConfValueObjectMap map[string]interface{}

var(
	ConfigIsChange chan bool
)

func LoadViper(){
	viper.SetConfigName("widgets_test") //  设置配置文件名 (不带后缀)
	//viper.AddConfigPath("/workspace/appName/") 
	viper.AddConfigPath(".")               // 比如添加当前目
	viper.SetConfigType("json")
	err := viper.ReadInConfig() // 搜索路径，并读取配置数据

	if err == nil {
		go watching()
	}else{
		panic(fmt.Errorf("Fatal init config file:",err))
	}
}

func NewConfValueObjectMapByType(typeString string)map[string]interface{}{
	m :=viper.Get(typeString)
	if value, ok :=m.(map[string]interface{});ok{
		return value
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
			ConfigIsChange <-true
			return
		}else{
			fmt.Println("Fatal reset config file:",err)
			return
		}
	})
}


