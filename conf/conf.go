package conf

import(
	"github.com/spf13/viper"
	"github.com/fsnotify/fsnotify"

	"fmt"
	"strings"
	//"encoding/json"
	"github.com/ziyouzy/mylib/physicalnode"
	"github.com/mitchellh/mapstructure"


)

var (
	confNodeMap map[string]interface{}
	confAlarmMap map[string]interface{}
)

func InitConfMap(){
	confNodeMap =make(map[string]interface{})
	viper.SetConfigName("riverconf") //  设置配置文件名 (不带后缀)
	//viper.AddConfigPath("/workspace/appName/") 
	viper.AddConfigPath(".")               // 比如添加当前目
	viper.SetConfigType("json")
	err := viper.ReadInConfig() // 搜索路径，并读取配置数据
	if err == nil {
		confNodeMap =updatemap("nodes")	
		confAlarmMap =updatemap("alarm")
		// ConfAllNodes :=viper.Get("nodes")
		// if value1, ok1 :=ConfAllNodes.([]interface{});ok1{
		// 	if confNodeMap,ok2 :=value1[0].(map[string]interface{});ok2{
		// 		fmt.Println("init confNodeMap success:",confNodeMap)
		// 	}else{
		// 		panic(fmt.Errorf("Fatal init ConfNodeMap! \n"))
		// 	}
		// }else{
		// 	panic(fmt.Errorf("Fatal init ConfNodeMap! \n"))
		// }

		// ConfAllAlarm :=viper.Get("alarm")
		// if value1, ok1 :=ConfAllAlarm.([]interface{});ok1{
		// 	if confAlarmMap,ok2 :=value1[0].(map[string]interface{});ok2{
		// 		fmt.Println("init ConfAlarmMap success:",confAlarmMap)
		// 	}else{
		// 		panic(fmt.Errorf("Fatal init ConfAlarmMap! \n"))
		// 	}
		// }else{
		// 	panic(fmt.Errorf("Fatal init ConfAlarmMap! \n"))
		// }

		go watching()
	}else{//if err == nil
		panic(fmt.Errorf("Fatal init config file! \n"))
	}//if err == nil end
}

func updatemap(typeString string) map[string]interface{}{
	ConfAllNodes :=viper.Get(typeString)
	if value1, ok1 :=ConfAllNodes.([]interface{});ok1{
		if m,ok2 :=value1[0].(map[string]interface{});ok2{
			fmt.Println("update ConfMap success")
			return m
		}else{
			panic(fmt.Errorf("Fatal init ConfMap! \n"))
		}
	}else{
		panic(fmt.Errorf("Fatal init ConfMap! \n"))
	}
}

func watching() {
	viper.WatchConfig()
	viper.OnConfigChange(func(e fsnotify.Event) {
		fmt.Println("Config file changed:", e.Name)
		err := viper.ReadInConfig() // 搜索路径，并读取配置数据
		if err == nil {
			ConfAllNodes :=viper.Get("Nodes")
			if value1, ok1 :=ConfAllNodes.([]interface{});ok1{
				if confNodeMap,ok2 :=value1[0].(map[string]interface{});ok2{
					fmt.Println("reset ConfNodeMap success:",confNodeMap)
				}else{
					fmt.Println("Fatal reset ConfNodeMap file!")
				}
			}else{
				fmt.Println("Fatal reset ConfNodeMap file!")
			}

			ConfAllSMS :=viper.Get("sms")
			if value1, ok1 :=ConfAllSMS.([]interface{});ok1{
				if confSMSMap,ok2 :=value1[0].(map[string]interface{});ok2{
					fmt.Println("reset ConfSMSMap success:",confSMSMap)
				}else{
					fmt.Println("Fatal reset ConfSMSMap file!")
				}
			}else{
				fmt.Println("Fatal reset ConfSMSMap file!")
			}
		}else{
			fmt.Println("Fatal reset config file!")
			return
		}
	})
}

//----



type ConfNode interface{
	CountPhysicalNode(string, string)
	GetMatrixSystemAndModuleString()(string, string, string)
	GetJson()[]byte
	GetMatrixSystemModuleAndCountJSON(string, string)(string, string, string, []byte)
	JudgeAlarm()string
}

//这个函数目前似乎只能生成一个confNode
//首先，confNodeMap的作用确实是基于他和for循环生成多个ConfNode
//但是当生成了一个只后就立刻返回了，于是map后面的内容都会被遗弃
func NewConfNodeArr(p physicalnode.PhysicalNode) []ConfNode {
	//fmt.Println("confNodeMap in, NewConfNode:",confNodeMap)
	//这里缺少一次判定，也就是某个物理节点是否被在conf中被提到了，没有的话，没必要耗费内存去做下面这些事
	phandlertag :=p.GetHandlerTagForConfNodeMap()
	var confnodearr []ConfNode
	for k,v := range confNodeMap{
		o :=k
		//fmt.Println(o)
		if !strings. Contains(o,phandlertag){
			//fmt.Println(o)
			continue
		}else{
			//fmt.Println(o)
			tempValue :=v
			tempstr	:= strings.Split(o,"-")
			handler :=tempstr[0]
			tag :=tempstr[1]
			nodename :=tempstr[2]
			valuetype :=tempstr[3]
			switch valuetype{
			case "bool":
				var confnode BoolenConfNode
				mapstructure.Decode(tempValue, &confnode)
				/*SeleteOneValueByProtocol会返回两个string，一个是值，一个是时间*/
				pvalue,ptime := p.SeleteOneValueByProtocol(handler, tag, nodename)
				confnode.CountPhysicalNode(pvalue,ptime)
				confnodearr =append(confnodearr, &confnode)
			case "int":
				var confnode IntConfNode
				mapstructure.Decode(tempValue, &confnode)
				pvalue,ptime := p.SeleteOneValueByProtocol(handler, tag, nodename)
				confnode.CountPhysicalNode(pvalue,ptime)
				confnodearr =append(confnodearr, &confnode)
			case "float":
				var confnode FloatConfNode
				mapstructure.Decode(tempValue, &confnode)
				pvalue,ptime := p.SeleteOneValueByProtocol(handler, tag, nodename)
				confnode.CountPhysicalNode(pvalue,ptime)
				confnodearr =append(confnodearr, &confnode)
			case "common", "string":
				var confnode CommonConfNode 
				mapstructure.Decode(tempValue, &confnode)
				pvalue,ptime := p.SeleteOneValueByProtocol(handler, tag, nodename)
				confnode.CountPhysicalNode(pvalue,ptime)
				confnodearr =append(confnodearr, &confnode)
			default:
				
			}
		}
	}
	//此处应该录入log，range  confNodeMap失败了
	return confnodearr
}

//---


