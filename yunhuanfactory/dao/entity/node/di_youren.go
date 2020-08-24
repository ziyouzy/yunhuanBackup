package di


//获取一个初始的entity实体
func DiYouRenByOldYunHuanMySqlDB(OldYunHuanMySqlDB oldyunhuanmysqldb) (entity.IEntity, error err){
	this.Id =oldyunhuanmysqldb.Id
	this.Name =oldyunhuanmysqldb.Node_name//后期从全局变量获取
	this.Input_time =oldyunhuanmysqldb.Time
	this.Value =oldyunhuanmysqldb.Value
	
	sl :=strings.Split(this.Value,"|")
	temp :=sl[0]
	//fmt.Println("temp:",temp,",sl[1]:",sl[1])
	if strings.Index(temp,"timeout") ==-1{
		c :=[]byte(temp)
		temp =string([]byte(c)[12:16])
		hex, _:=strconv.ParseInt(temp,16,0)
		temp =strconv.FormatInt(hex,2)
		temp =string([]byte(temp)[1:])

		this.Di8 =string([]byte(temp)[0:1])//“”的内容都是读取配置文档获得
		this.Di7 =string([]byte(temp)[1:2])//如这里的di7，this.Di7是string类型，这里给的值“1”也是string类型，调用warningMgr包中对应的函数进行加工
		this.Di6 =string([]byte(temp)[2:3])
		this.Di5 =/*"市电断电(0为开路;1为闭路)|"+*/string([]byte(temp)[3:4])
		this.Di4 =string([]byte(temp)[4:5])
		this.Di3 =string([]byte(temp)[5:6])
		this.Di2 =/*"顶盖状态(0为开路;1为闭路)|"+*/string([]byte(temp)[6:7])
		this.Di1 =/*"顶盖全打开(0为开路;1为闭路)|"+*/string([]byte(temp)[7:8])
	}else{
		this.Di8 ="timeout"
		this.Di7 ="timeout"
		this.Di6 ="timeout"
		this.Di5 ="timeout"
		this.Di4 ="timeout"
		this.Di3 ="timeout"
		this.Di2 ="timeout"
		this.Di1 ="timeout"
	}
}