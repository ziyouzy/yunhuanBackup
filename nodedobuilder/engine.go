//核心作用是将一个map[string]interface{}转化为map[string]NodeDo
package nodedobuilder

import(
	"strings"
	"fmt"

	"github.com/ziyouzy/mylib/nodedo"
)

type Engine map[string]nodedo.NodeDo

/** 会从myvipers包已经做好json文档映射关系的对象中拿数据
  * 当监测到json文件发生变动时需要再次执行
  */

func NewEngine(base map[string]interface{})Engine{
	var e = make(Engine) 
	/* k为各个json结构体的名称，是个string，v则是这个结构体，是个interface{}*/
	for key,inter := range base{
		/* 将k拆分后第4个元素可能为:bool、boolen、int、float、string、common*/
		datatype :=strings.Split(key,"-")[3]
		if nd :=nodedo.NewNodeDo(datatype,inter);nd !=nil{
			e[key] =nd
		}else{
			fmt.Println("在创建NodeEngine时，json字符串中，名为：",key,"中的",datatype,"类型无法被解析")
		}
	}
	return e
}

