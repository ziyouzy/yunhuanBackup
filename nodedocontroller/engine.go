//这一层不是io的do物理节点，而是代表了一个逻辑层
//也就是do，dto，vo中的do层
//从物理节点获取到数据，获得physicalnode无状态节点后，需要结合json配置文档生成的就是这一层，也就是有状态的do对象
package do

import(
	//"encoding/json"
	"strings"
	//"strconv"
	"fmt"

	"github.com/mitchellh/mapstructure"
	//"time"
	//"sync"
	//"github.com/ziyouzy/mylib/conf"
	
	//"github.com/ziyouzy/mylib/physicalnode"
	"github.com/ziyouzy/mylib/model"
)

type Engine map[string]NodeDo

//会从myvipers包已经做好json文档映射关系的对象中拿数据
//当监测到json文件发生变动时需要再次执行
func NewEngine(base map[string]interface{})Engine{
	var e = make(Engine) 
	for k,v := range base{
		switch strings.Split(k,"-")[3]{
		case "bool":
			var do BoolenNodeDo
			mapstructure.Decode(v, &do)
			e[k] =&do
		case "int":
			var do IntNodeDo
			mapstructure.Decode(v, &do)
			e[k] =&do
		case "float":
			var do FloatNodeDo
			mapstructure.Decode(v, &do)
			e[k] =&do
		case "common", "string":
			var do CommonNodeDo 
			mapstructure.Decode(v, &do)
			e[k] =&do
		default:
			fmt.Println("在创建NodeDo各个缓存时，json字符串中，名为：",k,"中的",strings.Split(k,"-")[3],"类型无法被解析")
			return nil
		}
	}
	return e
}

