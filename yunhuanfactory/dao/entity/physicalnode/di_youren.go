package di

import (
	"github.com/ziyouzy/mylib/yunhuanfactory/model/db"
	node "github.com/ziyouzy/mylib/yunhuanfactory/model/entity/node/impl"
)


//获取一个初始的entity实体
func DiYouRenByOldYunHuanMySqlDB(old db.OldYunHuanMySqlDB) (di node.Di, err error){
	di.Id =old.Id
	di.Name =old.Node_name//后期从全局变量获取
	di.Input_time =old.Time
	di.Value =old.Value
	
	//sl :=strings.Split(this.Value,"|")
	//temp :=sl[0]
	//fmt.Println("temp:",temp,",sl[1]:",sl[1])

			if temp.begin(0,4)=="494f"
	
	if strings.Index(temp,"timeout") ==-1{
		c :=[]byte(temp)
		temp =string([]byte(c)[12:16])
		hex, _:=strconv.ParseInt(temp,16,0)
		temp =strconv.FormatInt(hex,2)
		temp =string([]byte(temp)[1:])

		di.Di8 =string([]byte(temp)[0:1])//“”的内容都是读取配置文档获得
		di.Di7 =string([]byte(temp)[1:2])//如这里的di7，this.Di7是string类型，这里给的值“1”也是string类型，调用warningMgr包中对应的函数进行加工
		di.Di6 =string([]byte(temp)[2:3])
		di.Di5 =/*"市电断电(0为开路;1为闭路)|"+*/string([]byte(temp)[3:4])
		di.Di4 =string([]byte(temp)[4:5])
		di.Di3 =string([]byte(temp)[5:6])
		di.Di2 =/*"顶盖状态(0为开路;1为闭路)|"+*/string([]byte(temp)[6:7])
		di.Di1 =/*"顶盖全打开(0为开路;1为闭路)|"+*/string([]byte(temp)[7:8])
	}else{
		di.Di8 ="timeout"
		di.Di7 ="timeout"
		di.Di6 ="timeout"
		di.Di5 ="timeout"
		di.Di4 ="timeout"
		di.Di3 ="timeout"
		di.Di2 ="timeout"
		di.Di1 ="timeout"
	}
	return di,nil
}