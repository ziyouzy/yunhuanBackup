package do

import (
	"github.com/ziyouzy/mylib/yunhuanfactory/model/db"
	doModel "github.com/ziyouzy/mylib/yunhuanfactory/model/entity/node/impl"
	"github.com/ziyouzy/mylib/yunhuanfactory/model/entity/node"
)

func DoYouRenByOldYunHuanMySqlDB(old db.OldYunHuanMySqlDB) (doModel.Do, error){
	var do doModel.Do
	do.Id =oldyunhuanmysqldb.Id
	do.Name =oldyunhuanmysqldb.Node_name
	do.Input_time =oldyunhuanmysqldb.Time
	do.Value =oldyunhuanmysqldb.Value
	
	sl :=strings.Split(this.Value,"|")
	temp :=sl[0]
	//fmt.Println("temp:",temp,",sl[1]:",sl[1])
	if strings.Index(temp, "timeout") ==-1{
		c :=[]byte(temp)
		temp =string([]byte(c)[12:16])
		hex, _:=strconv.ParseInt(temp,16,0)
		temp =strconv.FormatInt(hex,2)
		temp =string([]byte(temp)[1:])

		do.Do8 =/*"顶盖恢复(0为开路;1为闭路)|"+*/string([]byte(temp)[0:1])
		do.Do7 =/*"顶盖开启(0为开路;1为闭路)|"+*/string([]byte(temp)[1:2])
		do.Do6 =/*"后门开关(0为开路;1为闭路)|"+*/string([]byte(temp)[2:3])
		do.Do5 =/*"前门开关(0为开路;1为闭路)|"+*/string([]byte(temp)[3:4])
		do.Do4 =/*"散热风扇(0为开路;1为闭路)|"+*/string([]byte(temp)[4:5])
		do.Do3 =/*"绿色灯带(0为开路;1为闭路)|"+*/string([]byte(temp)[5:6])
		do.Do2 =/*"红色灯带(0为开路;1为闭路)|"+*/string([]byte(temp)[6:7])
		do.Do1 =/*"蓝色灯带(0为开路;1为闭路)|"+*/string([]byte(temp)[7:8])
	}else{
		do.Do8 ="timeout"
		do.Do7 ="timeout"
		do.Do6 ="timeout"
		do.Do5 ="timeout"
		do.Do4 ="timeout"
		do.Do3 ="timeout"
		do.Do2 ="timeout"
		do.Do1 ="timeout"
	}
	return do,nil
}
