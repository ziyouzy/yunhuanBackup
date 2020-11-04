package protocol

import(
	"github.com/ziyouzy/mylib/conf"
	//"github.com/ziyouzy/mylib/view" 取代他的是do包
	"github.com/ziyouzy/mylib/model"

	//"fmt"
	"encoding/json"
	"sync"
	"time"
)



var (
	smsticket float64
	smsticketlimit float64

	mysqlticket float64
	mysqlticketlimit float64
)

func ProtocolEnding_YunHuan20201101(nodedoentitych chan do.NodeDo)(chan []byte, chan []byte, chan *model.AlarmEntity){
	nodeDoCh := make(chan []byte)//用来发送给前端软件,其实是将nodedo的实体再经过判定告警的操作后再转化为字节数组
	alarmSMSCh := make(chan []byte)//用来232串口的短信报警
	alarmEntityCh := make(chan *model.AlarmEntity)//用来mysql的告警信息录入

	smsticketlimit =3600*24
	smsticket =3600*24

	mysqlticketlimit =3600*24
	mysqlticket =3600*24

	//这个循环alarmsms和alarmdb也会包括在内
	go func(){
		for entity := range nodedoentitych{
			//如果==nil则说明当前nodeDo实体不存在超限问题
			//也就不需要用到smsch和mysqlch这两个管道了
			//而在这if内部其实是先创建一个confalarm
			//再进行基于这个对象的一系列操作，从而实现周期告警的功能
			//他应该是个常驻于缓存中的拦截器
			//有点像是do.cache但是do.cache是一个常驻于缓存中的value object
			//代码的实现方式相同，但是意义不同
			if confalarm :=conf.NewConfAlram(entity);confalarm !=nil{
				smsticketlimit =confalarm.SMSSleepMin*60
				mysqlticketlimit =confalarm.MySQLSleepMin*60

				//第一次的时候必然会大于，因为smsticket=3600*24
				if(smsticket >=smsticketlimit){
					go func(){
						for _, sms := range confalarm.SMS{
							alarmSMSCh<-[]byte(sms)
							time.Sleep(time.Duration(500)*time.Millisecond)			
						}
					}()
					smsticket =0
				}

				
				//第一次的时候必然会大于，因为mysqlticket=3600*24
				if(mysqlticket >=mysqlticketlimit){
					go func(){
						//装配一个alarmentity并放入管道
						ae := 	model.AlarmEntity{
							Name : confalarm.MySQLNameString,
							Value : confalarm.MySQLValueString,
							Unit : confalarm.MySQLUnitString,
							Content : confalarm.MySQLContentString,
						}
						alarmEntityCh<-&ae
					}()
					mysqlticket =0
				}
			}//confalarm :=conf.NewConfAlram(entity);confalarm !=nil end

			nodeDoCh<-entity.GetJson()
		}
	}()

	//实时更新sms的ticket，以分钟为单位
	go func (){
		for{
			if smsticket<=smsticketlimit{
				smsticket =smsticket+1
			}
			time.Sleep(time.Second)
		}
	}()//循环外的sms发送控制

	//实时更新db的ticket，以分钟为单位
	go func (){
		for{
			if mysqlticket<=mysqlticketlimit{
				mysqlticket =mysqlticket+1
			}
			time.Sleep(time.Second)
		}
	}()

	return nodeDoCh, alarmSMSCh, alarmEntityCh
}