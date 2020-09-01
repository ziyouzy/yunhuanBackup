package impl//这个包交叉引用了父目录包physicalnode

import (
	"strconv"
	"strings"
	dbreflectorlv1 "github.com/ziyouzy/mylib/yunhuanfactory/dbreflector/mysql/lv1/impl"

	physicalnodeentitylv1  "github.com/ziyouzy/mylib/yunhuanfactory/entity/physicalnode/lv1/impl"
	physicalnodeentity "github.com/ziyouzy/mylib/yunhuanfactory/entity/physicalnode"
)

type Do_YouRenEntityDaoImpl struct{
	NodeType string
	Edition string

	Old *dbreflectorlv1.OldNodeReflectorImpl
}

func (this *Do_YouRenEntityDaoImpl)CreateNodeEntityFromOldNodeEntity() physicalnodeentity.IYouRenEntity{
	var do physicalnodeentitylv1.Do_YouRenEntityImpl
	do.Do_id =this.Old.Id
	do.Do_name =this.Old.Node_name
	do.Do_input_time =this.Old.Time
	do.Do_value =this.Old.Value
	
	tempStr :=strings.Split(this.Old.Value,"|")[0]
	//fmt.Println("temp:",temp,",sl[1]:",sl[1])
	if strings.Contains(tempStr, "timeout"){
		do.Do8 ="timeout"
		do.Do7 ="timeout"
		do.Do6 ="timeout"
		do.Do5 ="timeout"
		do.Do4 ="timeout"
		do.Do3 ="timeout"
		do.Do2 ="timeout"
		do.Do1 ="timeout"
	}else if strings.Index(tempStr,"494f")==0{
		c :=[]byte(tempStr)
		tempStr =string([]byte(c)[12:16])
		hex, _:=strconv.ParseInt(tempStr,16,0)
		tempStr =strconv.FormatInt(hex,2)
		tempStr =string([]byte(tempStr)[1:])

		do.Do8 =/*"顶盖恢复(0为开路;1为闭路)|"+*/string([]byte(tempStr)[0:1])
		do.Do7 =/*"顶盖开启(0为开路;1为闭路)|"+*/string([]byte(tempStr)[1:2])
		do.Do6 =/*"后门开关(0为开路;1为闭路)|"+*/string([]byte(tempStr)[2:3])
		do.Do5 =/*"前门开关(0为开路;1为闭路)|"+*/string([]byte(tempStr)[3:4])
		do.Do4 =/*"散热风扇(0为开路;1为闭路)|"+*/string([]byte(tempStr)[4:5])
		do.Do3 =/*"绿色灯带(0为开路;1为闭路)|"+*/string([]byte(tempStr)[5:6])
		do.Do2 =/*"红色灯带(0为开路;1为闭路)|"+*/string([]byte(tempStr)[6:7])
		do.Do1 =/*"蓝色灯带(0为开路;1为闭路)|"+*/string([]byte(tempStr)[7:8])
	}else{
		do.Do8 ="err"
		do.Do7 ="err"
		do.Do6 ="err"
		do.Do5 ="err"
		do.Do4 ="err"
		do.Do3 ="err"
		do.Do2 ="err"
		do.Do1 ="err"
	}

	return &do
}
