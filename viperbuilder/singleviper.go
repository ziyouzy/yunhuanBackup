//每个singleViper只会被new一次
//之后都只是更新
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

		sv.OneViperConfigIsChangeAndUpdateFinishCh =make(chan string)
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
	OneViperConfigIsChangeAndUpdateFinishCh chan string
}

func (p *SingleViper)watching() {
	p.V.WatchConfig()
	p.V.OnConfigChange(func(e fsnotify.Event) {
		fmt.Println("Config file changed:", e.Name)
		err := p.V.ReadInConfig() // 搜索路径，并读取配置数据
		if err == nil {
			p.OneViperConfigIsChangeAndUpdateFinishCh <-p.NameWithPathAndSuffix
			close(p.OneViperConfigIsChangeAndUpdateFinishCh)
			fmt.Println("Success reset config file")
			return
		}else{
			fmt.Println("Fatal reset config file:",err)
			return
		}
	})
}




