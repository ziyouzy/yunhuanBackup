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
//一个viper对应了一个文本文件，核心任务是监听该文件的修改并更新相关上层对象
//对于上层，目前主要是alarm包和do包，在去考虑一个viper对象会生成几个alarmCache和nodedoCache对象
//届时会用到功能与NewConfValueObjectMapByType相近的方法或函数，基于一个viper对象很可能会生成多个Cache对象
//总之这也都是上一层需要去做的，而NewConfValueObjectMapByType很可能会变成SingleViper的方法之一，就不用再去单独设计依赖注入了
package myvipers

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
	Vipers map[string]*SingleViper
)

//只设计两种情况：要么是绝对路径，要么是根目录
func Load(jsonpaths ...string){
	for _, data :=range jsonpath{
		if sv :=NewSingleViper(data); sv!=nil{
			sv.ListenConfigChange()
			Vipers[data] =sv
		}else{	
			fmt.Println("您设置的json路径[",data,"]格式错误，只支持绝对路径与根目录两种模式")
		}
	}
}
