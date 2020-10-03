package conf

import(
	"github.com/spf13/viper"
	"github.com/fsnotify/fsnotify"
	"github.com/mitchellh/mapstructure"

	"fmt"
	"strings"
	"encoding/json"
)

var confNodeMap map[string]interface{}

func InitConfMap(){
	confNodeMap =make(map[string]interface{})
	viper.SetConfigName("riverconf") //  设置配置文件名 (不带后缀)
	//viper.AddConfigPath("/workspace/appName/") 
	viper.AddConfigPath(".")               // 比如添加当前目
	viper.SetConfigType("json")
	err := viper.ReadInConfig() // 搜索路径，并读取配置数据
	if err == nil {
		//ConfListarr :=viper.Get("SMS")
		ConfAllNodes :=viper.Get("Nodes")
		if value1, ok1 :=ConfAllNodes.([]interface{});ok1{
			if confNodeMap,ok2 :=value1[0].(map[string]interface{});ok2{
				fmt.Println("init ConfNodeMap success!")
				go watching()
				return
			}
		}
		panic(fmt.Errorf("Fatal init config file! \n")
	}else{
		panic(fmt.Errorf("Fatal init config file! \n")
	}
}


func watching() {
	viper.WatchConfig()
	viper.OnConfigChange(func(e fsnotify.Event) {
		fmt.Println("Config file changed:", e.Name)
		err := viper.ReadInConfig() // 搜索路径，并读取配置数据
		if err == nil {
			//ConfListarr :=viper.Get("SMS")
			ConfAllNodes :=viper.Get("Nodes")
			if value1, ok1 :=ConfAllNodes.([]interface{});ok1{
				if confNodeMap,ok2 :=value1[0].(map[string]interface{});ok2{
					fmt.Println("reset ConfNodeMap success!")
					return
				}
			}
			fmt.Println("Fatal reset config file!")
		}else{
			fmt.Println("Fatal reset config file!")
			return
		}
	})
}