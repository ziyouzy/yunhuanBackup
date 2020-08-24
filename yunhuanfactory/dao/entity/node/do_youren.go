package do

func DoYouRenByOldYunHuanMySqlDB(OldYunHuanMySqlDB oldyunhuanmysqldb) error err{
	this.Id =oldyunhuanmysqldb.Id
	this.Name =oldyunhuanmysqldb.Node_name
	this.Input_time =oldyunhuanmysqldb.Time
	this.Value =oldyunhuanmysqldb.Value
	
	sl :=strings.Split(this.Value,"|")
	temp :=sl[0]
	//fmt.Println("temp:",temp,",sl[1]:",sl[1])
	if strings.Index(temp, "timeout") ==-1{
		c :=[]byte(temp)
		temp =string([]byte(c)[12:16])
		hex, _:=strconv.ParseInt(temp,16,0)
		temp =strconv.FormatInt(hex,2)
		temp =string([]byte(temp)[1:])

		this.Do8 =/*"顶盖恢复(0为开路;1为闭路)|"+*/string([]byte(temp)[0:1])
		this.Do7 =/*"顶盖开启(0为开路;1为闭路)|"+*/string([]byte(temp)[1:2])
		this.Do6 =/*"后门开关(0为开路;1为闭路)|"+*/string([]byte(temp)[2:3])
		this.Do5 =/*"前门开关(0为开路;1为闭路)|"+*/string([]byte(temp)[3:4])
		this.Do4 =/*"散热风扇(0为开路;1为闭路)|"+*/string([]byte(temp)[4:5])
		this.Do3 =/*"绿色灯带(0为开路;1为闭路)|"+*/string([]byte(temp)[5:6])
		this.Do2 =/*"红色灯带(0为开路;1为闭路)|"+*/string([]byte(temp)[6:7])
		this.Do1 =/*"蓝色灯带(0为开路;1为闭路)|"+*/string([]byte(temp)[7:8])
	}else{
		this.Do8 ="timeout"
		this.Do7 ="timeout"
		this.Do6 ="timeout"
		this.Do5 ="timeout"
		this.Do4 ="timeout"
		this.Do3 ="timeout"
		this.Do2 ="timeout"
		this.Do1 ="timeout"
	}
}
