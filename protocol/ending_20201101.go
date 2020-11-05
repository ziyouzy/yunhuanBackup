//其实这里主要会用到两个包:
//一个是do，另一个是alarm
//我已经尽最大努力把他们设计成了和viper一样的模式
//也就是在主函数初始化后，包就可以在全局使用了
//因此在这里引用后，直接去使用即可
//这也是一次考验，之前的设计思路过不过关
//ps：实例化的缓存都存在于conf包内，而不是分别存在于do和alarm内
package protocol

import(
	"github.com/ziyouzy/mylib/conf"
	"github.com/ziyouzy/mylib/do" 
	"github.com/ziyouzy/mylib/physicalnode"
	"github.com/ziyouzy/mylib/model"

	//"fmt"
	//"encoding/json"
	//"sync"
	"time"
)



// var (
// 	smsticket float64
// 	smsticketlimit float64

// 	mysqlticket float64
// 	mysqlticketlimit float64
// )

//传入的参数不再会是NodeDo的管道，而直接变成了PhysicalNode的管道
//我也考虑了要不要先用NodeDoCache把PhysicalNode转化为Nodedo
//但是似乎不太合理，因为应该有个工具函数将PhysicalNode分别扇出为NodeDo的管道，短信告警的管道，还有告警录入数据库的管道
//这个函数似乎既可以属于主函数，也可以属于protocol包，考虑到以后每个项目到了这里很可能都会有巨大的不同，所以还是先让他属于protocol吧
func ProtocolEnding_YunHuan20201101(physicalnodech chan physicalnode.PhysicalNode)(chan []byte, chan []byte, chan *model.AlarmEntity){
	nodeDoBytesCh := make(chan []byte)//用来发送给前端软件,其实是将nodedo的每个实体，先经过判定告警的操作后，再转化为字节数组
	alarmSMSBytesCh := make(chan []byte)//用来232串口的短信报警，会将string转化为[]byte
	alarmEntityCh := make(chan *model.AlarmEntity)//用来mysql的告警信息录入

	NodeDoEntityCh :=conf.NodeDoCache.CreateNodeDoCh

	//处理physicalnode的管道
	go func(){
		for pn := range physicalnodech{
			//这里是利用物理节点来渲染NodeDoCache内的各个NodeDo模板
			conf.NodeDoCache.UpdateNodeDoMap(pn)


	// 		if confalarm :=conf.NewConfAlram(entity);confalarm !=nil{
	// 			smsticketlimit =confalarm.SMSSleepMin*60
	// 			mysqlticketlimit =confalarm.MySQLSleepMin*60

	// 			//第一次的时候必然会大于，因为smsticket=3600*24
	// 			if(smsticket >=smsticketlimit){
	// 				go func(){
	// 					for _, sms := range confalarm.SMS{
	// 						alarmSMSCh<-[]byte(sms)
	// 						time.Sleep(time.Duration(500)*time.Millisecond)			
	// 					}
	// 				}()
	// 				smsticket =0
	// 			}

				
	// 			//第一次的时候必然会大于，因为mysqlticket=3600*24
	// 			if(mysqlticket >=mysqlticketlimit){
	// 				go func(){
	// 					//装配一个alarmentity并放入管道
	// 					ae := 	model.AlarmEntity{
	// 						Name : confalarm.MySQLNameString,
	// 						Value : confalarm.MySQLValueString,
	// 						Unit : confalarm.MySQLUnitString,
	// 						Content : confalarm.MySQLContentString,
	// 					}
	// 					alarmEntityCh<-&ae
	// 				}()
	// 				mysqlticket =0
	// 			}
	// 		}//confalarm :=conf.NewConfAlram(entity);confalarm !=nil end

	// 		nodeDoCh<-entity.GetJson()
	// 	}
	// }()

	// //实时更新sms的ticket，以分钟为单位
	// go func (){
	// 	for{
	// 		if smsticket<=smsticketlimit{
	// 			smsticket =smsticket+1
	// 		}
	// 		time.Sleep(time.Second)
	// 	}
	// }()//循环外的sms发送控制

	// //实时更新db的ticket，以分钟为单位
	// go func (){
	// 	for{
	// 		if mysqlticket<=mysqlticketlimit{
	// 			mysqlticket =mysqlticket+1
	// 		}
	// 		time.Sleep(time.Second)
	// 	}
	}()

	return nodeDoCh, alarmSMSCh, alarmEntityCh
}