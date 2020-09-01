package impl//这个包交叉引用了父目录包physicalnode

import (
	"strconv"
	"strings"
	dbreflectorlv1 "github.com/ziyouzy/mylib/yunhuanfactory/dbreflector/mysql/lv1/impl"

	physicalnodeentitylv1  "github.com/ziyouzy/mylib/yunhuanfactory/entity/physicalnode/lv1/impl"
	physicalnodeentity "github.com/ziyouzy/mylib/yunhuanfactory/entity/physicalnode"
)

type Di_YouRenEntityDaoImpl struct{
	NodeType string
	Edition string

	Old *dbreflectorlv1.OldNodeReflectorImpl
}
//获取一个初始的entity实体
func (this *Di_YouRenEntityDaoImpl)CreateNodeEntityFromOldNodeEntity() physicalnodeentity.IYouRenEntity{
	var di physicalnodeentitylv1.Di_YouRenEntityImpl
	di.Di_id =this.Old.Id
	di.Di_name =this.Old.Node_name//后期从全局变量获取
	di.Di_input_time =this.Old.Time
	di.Di_value =this.Old.Value
	
	tempStr :=strings.Split(this.Old.Value,"|")[0]
	//fmt.Println("temp:",temp,",sl[1]:",sl[1])
	if strings.Contains(tempStr, "timeout"){
		di.Di8 ="timeout"
		di.Di7 ="timeout"
		di.Di6 ="timeout"
		di.Di5 ="timeout"
		di.Di4 ="timeout"
		di.Di3 ="timeout"
		di.Di2 ="timeout"
		di.Di1 ="timeout"
	}else if strings.Index(tempStr,"494f")==0{
		c :=[]byte(tempStr)
		tempStr =string([]byte(c)[12:16])
		hex, _:=strconv.ParseInt(tempStr,16,0)
		tempStr =strconv.FormatInt(hex,2)
		tempStr =string([]byte(tempStr)[1:])

		di.Di8 =string([]byte(tempStr)[0:1])//“”的内容都是读取配置文档获得
		di.Di7 =string([]byte(tempStr)[1:2])//如这里的di7，this.Di7是string类型，这里给的值“1”也是string类型，调用warningMgr包中对应的函数进行加工
		di.Di6 =string([]byte(tempStr)[2:3])
		di.Di5 =/*"市电断电(0为开路;1为闭路)|"+*/string([]byte(tempStr)[3:4])
		di.Di4 =string([]byte(tempStr)[4:5])
		di.Di3 =string([]byte(tempStr)[5:6])
		di.Di2 =/*"顶盖状态(0为开路;1为闭路)|"+*/string([]byte(tempStr)[6:7])
		di.Di1 =/*"顶盖全打开(0为开路;1为闭路)|"+*/string([]byte(tempStr)[7:8])
	}else{
		di.Di8 = "err"
		di.Di7 = "err" 
		di.Di6 = "err"
		di.Di5 = "err"
		di.Di4 = "err"
		di.Di3 = "err"
		di.Di2 = "err"
		di.Di1 = "err"
	}

	return &di
}

