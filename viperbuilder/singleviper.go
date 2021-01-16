//每个singleViper只会被new一次，监听到文件被修改后会立刻通过v.ReadInConfig()更新，同时向上层汇报更新的文件路径名称
//一个singleViper对应了一个json的文本文件/路径
package viperbuilder

import(
	"github.com/spf13/viper"
	"github.com/fsnotify/fsnotify"

	"strings"
	"fmt"
)

//SingleViper是文件级的
func BuildSingleViper(namewithpathandsuffix string)*SingleViper{
	strs :=strings.Split(namewithpathandsuffix, "/");    namewithsuffix :=strs[len(strs)-1];    path :=strings.Replace(namewithpathandsuffix,namewithsuffix,"",1)
	fmt.Println("path:",path,"namewithpathandsuffix:",namewithpathandsuffix,"namewithsuffix:",namewithsuffix)
	//对相对路径的判定与额外操作
	if strings.Compare(path,"./")==0 { path ="." }

	strs =strings.Split(namewithsuffix,".");        name :=strs[0];        suffix :=strs[1]
	
	v :=viper.New()
	v.SetConfigName(name) 
	v.AddConfigPath(path)   
	v.SetConfigType(suffix)//json yarm

	err := v.ReadInConfig() // 搜索路径，并读取配置数据

	if err == nil {
		sv :=SingleViper{
			NameWithPathAndSuffix: namewithpathandsuffix,
			Name : name,
			Path : path,
			Suffix :suffix,

			V:v,
		}

		sv.OneViperConfigChangedCh =make(chan string)
		sv.watching()
		return &sv
	}else{
		fmt.Println("初始化配置json失败,名称、路径、后缀名分别为:",name,path,suffix)
		return nil
	}
}

type SingleViper struct{
	NameWithPathAndSuffix string
	Name string
	Path string
	Suffix string

	V *viper.Viper
	OneViperConfigChangedCh chan string
}

func (p *SingleViper)watching() {
	/** 不能直接去p.V.OnConfigChange = func(e fsnotify.Event) { fmt.Println("Config file changed:", e.Name)}
	  * 因为V的内部根本不存在 OnConfigChange变量
	  * 只存在私有的onConfigChange变量，以及OnConfigChange方法
	  * 因此大神的设计模式是这样的：
	  * 私有变量通过名称相同，但首字母大写的公有方法实现初始化
	  * 同时为了匹配单例模式，单例函数(非方法)与viper结构体的OnConfigChange方法名称一致，我一直也在用这种设计模式

	  * 大神这种设计模式的目的只有一个，就是给使用者暴露一个自定义函数
	  * 为了实现这个目的这个函数必须定义在结构体内部，而不能把他设计成这个结构体的方法
	  * 同时这也是借助了闭包的特性实现的，因为当使用者具象话这个函数的具体功能后
	  * 该函数就会以函数变量的形式存在于结构体内部了，也就是一个内部字段
	  * 等同于在一个方法内部使用了闭包
	  * 或者说，一个方法内部的闭包函数对象，都可以把转移成结构体自身的内部字段
	  * 这样会变得十分优雅，也是函数变量的意义与价值所在
	  */

	p.V.OnConfigChange(func(e fsnotify.Event) {
			  
	    /** OnConfigChange方法虽然没有在变量名上表达出Set、Init等初始化意图
	      * 但是其的功能只有一个，那就是初始化viper.onConfigChange这一私有字段
	      * 以后我也这么写就好，变量名不用加Set、Init之类的
		  */
		  
		fmt.Println("Config file changed:", e.Name)

		err := p.V.ReadInConfig() // 搜索路径，并读取配置数据

		if err == nil {
			p.OneViperConfigChangedCh <-p.NameWithPathAndSuffix
			fmt.Println("Success reset config file")
			return
		}else{
			fmt.Println("Fatal reset config file:",err)
			return
		}
	})

p.V.WatchConfig()
}

func (p *SingleViper)Destory(){
	close(p.OneViperConfigChangedCh)
}




