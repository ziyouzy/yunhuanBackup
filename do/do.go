//这一层不是io的do物理节点，而是代表了一个逻辑层
//也就是do，dto，vo中的do层
//从物理节点获取到数据，获得physicalnode无状态节点后，需要结合json配置文档生成的就是这一层，也就是有状态的do对象
package do

import(
	"encoding/json"
	"strings"
	"strconv"
	"fmt"
	"time"
	"sync"
	
	//"github.com/ziyouzy/mylib/physicalnode"
)



type NodeDo interface{
	CountPhysicalNode(string, string)
	GetMatrixSystemAndModuleString()(string, string, string)
	GetJson()[]byte
	GetMatrixSystemModuleAndCountJSON(string, string)(string, string, string, []byte)
	JudgeAlarm()string
	PrepareMYSQLAlarm()(string,string,string,string)
}

//会结合conf中的NodeDoVO生成当前配置文件所描述的所有物理节点VO的缓存map
func NewNodeDoValueObjectMap(base conf.ConfValueObjectMap)m map[string]*NodeDo{
	for k,v := range base{
		switch strings.Split(k,"-")[3]{
		case "bool":
			var do BoolenNodeDo
			mapstructure.Decode(v, &do)
			m =append(m, &do)
		case "int":
			var do IntNodeDo
			mapstructure.Decode(v, &do)
			m =append(m, &do)
		case "float":
			var do FloatNodeDo
			mapstructure.Decode(v, &do)
			m =append(m, &do)
		case "common", "string":
			var do CommonNodeDo 
			mapstructure.Decode(v, &do)
			m =append(m, &do)
		default:
			fmt.Println("在创建NodeVo各个缓存时，json字符串中，名为：",k,"中的",strings.Split(k,"-")[3],"类型无法被解析")
		}
	}
	return
}

//这个函数目前似乎只能生成一个confNode
//首先，confNodeMap的作用确实是基于他和for循环生成多个ConfNode
//但是当生成了一个只后就立刻返回了，于是map后面的内容都会被遗弃
// func NewNodeDoArr(p physicalnode.PhysicalNode) []NodeDo {
// 	//fmt.Println("confNodeMap in, NewConfNode:",confNodeMap)
// 	//这里缺少一次判定，也就是某个物理节点是否被在conf中被提到了，没有的话，没必要耗费内存去做下面这些事
// 	phandlertag :=p.GetHandlerTagForConfNodeMap()
// 	var confnodearr []ConfNode
// 	for k,v := range confNodeMap{
// 		o :=k
// 		//fmt.Println(o)
// 		if !strings. Contains(o,phandlertag){
// 			//fmt.Println(o)
// 			continue
// 		}else{
// 			//fmt.Println(o)
// 			tempValue :=v
// 			tempstr	:= strings.Split(o,"-")
// 			handler :=tempstr[0]
// 			tag :=tempstr[1]
// 			nodename :=tempstr[2]
// 			valuetype :=tempstr[3]
// 			switch valuetype{
// 			case "bool":
// 				var confnode BoolenConfNode
// 				mapstructure.Decode(tempValue, &confnode)
// 				/*SeleteOneValueByProtocol会返回两个string，一个是值，一个是时间*/
// 				pvalue,ptime := p.SeleteOneValueByProtocol(handler, tag, nodename)
// 				confnode.CountPhysicalNode(pvalue,ptime)
// 				confnodearr =append(confnodearr, &confnode)
// 			case "int":
// 				var confnode IntConfNode
// 				mapstructure.Decode(tempValue, &confnode)
// 				pvalue,ptime := p.SeleteOneValueByProtocol(handler, tag, nodename)
// 				confnode.CountPhysicalNode(pvalue,ptime)
// 				confnodearr =append(confnodearr, &confnode)
// 			case "float":
// 				var confnode FloatConfNode
// 				mapstructure.Decode(tempValue, &confnode)
// 				pvalue,ptime := p.SeleteOneValueByProtocol(handler, tag, nodename)
// 				confnode.CountPhysicalNode(pvalue,ptime)
// 				confnodearr =append(confnodearr, &confnode)
// 			case "common", "string":
// 				var confnode CommonConfNode 
// 				mapstructure.Decode(tempValue, &confnode)
// 				pvalue,ptime := p.SeleteOneValueByProtocol(handler, tag, nodename)
// 				confnode.CountPhysicalNode(pvalue,ptime)
// 				confnodearr =append(confnodearr, &confnode)
// 			default:
				
// 			}
// 		}
// 	}
// 	//此处应该录入log，range  confNodeMap失败了
// 	return confnodearr
// }
