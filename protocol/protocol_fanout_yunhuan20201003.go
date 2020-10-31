//注意！每一个客户需求彼此之间的数据不能共享
//同时一个客户对应一个携程，对应一个携程函数
//这里存在多个用户，其实只是为了测试
//这里是数据共享的设计方式，所以最终需要枷锁
package protocol

import(
	"github.com/ziyouzy/mylib/conf"
	"github.com/ziyouzy/mylib/view"
	"github.com/ziyouzy/mylib/model"

	//"fmt"
	"encoding/json"
	"sync"
	"time"
)


	/*以下是自定义协议本体*/
	//假设有三个客户需求:
	//1.客户需要接收到单个或相互独立的多个module，因此需要创建module管道
	//2.客户需要接收到单个相互独立的多个system，因此需要创建system管道
	//3.客户需要接收到单个相互独立的多个matrix，因此需要创建matrix管道
	//现在我们将需求反过来看，创建matrix的话我们需要让其子结构体都有要明确的name标识
	//否则前端就无法与组件做匹配了

/*代表一位有module级别需求的客户*/
var (
	lock_1 *sync.Mutex
	testyunhuan20201010_1_module1 =view.Module{}
	testyunhuan20201010_1_module2 =view.Module{}
	testyunhuan20201010_1_module3 =view.Module{}
)

/*代表一位有system级别需求的客户*/
var (
	lock_2 *sync.Mutex
	testyunhuan20201010_2_module1 =view.Module{}
	testyunhuan20201010_2_module2 =view.Module{}
	testyunhuan20201010_2_module3 =view.Module{}

	testyunhuan20201010_2_system1 =view.System{}
	testyunhuan20201010_2_system2 =view.System{}
	testyunhuan20201010_2_system3 =view.System{}
)

/*代表一位有matrix级别需求的客户*/
var (
	lock_3 *sync.Mutex
	testyunhuan20201010_3_module1 =view.Module{}
	testyunhuan20201010_3_module2 =view.Module{}
	testyunhuan20201010_3_module3 =view.Module{}

	testyunhuan20201010_3_system1 =view.System{}
	testyunhuan20201010_3_system2 =view.System{}
	testyunhuan20201010_3_system3 =view.System{}

	testyunhuan20201010_3_matrix1 =view.Matrix{}
	testyunhuan20201010_3_matrix2 =view.Matrix{}
)

var (
	smsticket float64
	smsticketlimit float64

	mysqlticket float64
	mysqlticketlimit float64
)

//拿到主函数去使用,入参是个实体，返回的是个管道
func ProtocolViewNodesHandler_YunHuan20201004(confnodech chan conf.ConfNode)(chan []byte, chan []byte, chan []byte, chan []byte, chan *model.AlarmEntity){
	moduleViewCh := make(chan []byte)
	systemViewCh := make(chan []byte)
	matrixViewCh := make(chan []byte)

	alarmSMSCh := make(chan []byte)
	alarmEntityCh := make(chan *model.AlarmEntity)

	smsticketlimit =3600*24
	smsticket =3600*24

	mysqlticketlimit =3600*24
	mysqlticket =3600*24

	mysqlticketlimit =3600*24

	lock_1 = new(sync.Mutex)
	lock_2 = new(sync.Mutex)
	lock_3 = new(sync.Mutex)

	//这个循环alarmsms和alarmdb也会包括在内
	go func(){
		for confnode := range confnodech{
			if confalarm :=conf.NewConfAlram(confnode);confalarm !=nil{
				smsticketlimit =confalarm.SMSSleepMin*60
				mysqlticketlimit =confalarm.MySQLSleepMin*60

				if(smsticket >=smsticketlimit){
					go func(){
						for _, sms := range confalarm.SMS{
							alarmSMSCh<-[]byte(sms)
							time.Sleep(time.Duration(500)*time.Millisecond)			
						}
					}()
					smsticket =0
				}

				
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
			}

			/*------*/
			_, _, module := confnode.GetMatrixSystemAndModuleString()
			switch (module){
			case "环境监测":
				go func(){
					lock_1.Lock()
					if testyunhuan20201010_1_module1.NodesLen <3{
						testyunhuan20201010_1_module1.AppendNode(confnode)
					}
					lock_1.Unlock()
				}()

				go func(){
					lock_2.Lock()
					if testyunhuan20201010_2_module1.NodesLen <3{
						testyunhuan20201010_2_module1.AppendNode(confnode)
					}
					lock_2.Unlock()
				}()

				go func(){
					lock_3.Lock()
					if testyunhuan20201010_3_module1.NodesLen <3{
						testyunhuan20201010_3_module1.AppendNode(confnode)
					}
					lock_3.Unlock()
				}()
				
			case "ups监测":
				go func(){
					lock_1.Lock()
					if testyunhuan20201010_1_module2.NodesLen <3{
						testyunhuan20201010_1_module2.AppendNode(confnode)
					}
					lock_1.Unlock()
				}()
			
				go func(){
					lock_2.Lock()
					if testyunhuan20201010_2_module2.NodesLen <3{
						testyunhuan20201010_2_module2.AppendNode(confnode)
					}
					lock_2.Unlock()
				}()
				
				go func(){
					lock_3.Lock()
					if testyunhuan20201010_3_module2.NodesLen <3{
						testyunhuan20201010_3_module2.AppendNode(confnode)
					}
					lock_3.Unlock()
				}()
				
			case "zndb监测":
				go func(){
					lock_1.Lock()
					if testyunhuan20201010_1_module3.NodesLen <3{
						testyunhuan20201010_1_module3.AppendNode(confnode)
					}
					lock_1.Unlock()
				}()
				
				go func(){
					lock_2.Lock()
					if testyunhuan20201010_2_module3.NodesLen <3{
						testyunhuan20201010_2_module3.AppendNode(confnode)
					}
					lock_2.Unlock()
				}()
				
				go func(){
					lock_3.Lock()
					if testyunhuan20201010_3_module3.NodesLen <3{
						testyunhuan20201010_3_module3.AppendNode(confnode)
					}
					lock_3.Unlock()
				}()
			}//循环内的view渲染结束
		}//循环结束
	}()//该线程函数结束

		//处理有module级别需求的客户：
	go func (){
		
		for {
			time.Sleep(time.Duration(2000)*time.Millisecond)

			lock_1.Lock()
			if data, err := json.Marshal(testyunhuan20201010_1_module1);err == nil{
				moduleViewCh<-data
			}

			if data, err := json.Marshal(testyunhuan20201010_1_module2);err == nil{
				moduleViewCh<- data
			}

			if data, err := json.Marshal(testyunhuan20201010_1_module3);err == nil{
				moduleViewCh<- data
			}
			
			testyunhuan20201010_1_module1.Reset()
			testyunhuan20201010_1_module2.Reset()
			testyunhuan20201010_1_module3.Reset()
			lock_1.Unlock()

		}

	}()//循环外的module"阀门"

	//处理有system级别需求的客户：
	go func (){
		for{
			time.Sleep(time.Duration(2000)*time.Millisecond)
	
			lock_2.Lock()
			if (testyunhuan20201010_2_module1.SystemName == "智能机柜"){
				testyunhuan20201010_2_system2.AppendModule(testyunhuan20201010_2_module1)	
			}

			if (testyunhuan20201010_2_module2.SystemName == "智能机柜"){
				testyunhuan20201010_2_system2.AppendModule(testyunhuan20201010_2_module2)	
			}

			if (testyunhuan20201010_2_module3.SystemName == "智能机柜"){
				testyunhuan20201010_2_system2.AppendModule(testyunhuan20201010_2_module3)	
			}

			if data, err := json.Marshal(testyunhuan20201010_2_system1);err ==nil{
				systemViewCh<-data
			}

			if data, err := json.Marshal(testyunhuan20201010_2_system2);err ==nil{
				systemViewCh<- data
			}

			if data, err := json.Marshal(testyunhuan20201010_2_system3);err ==nil{
				systemViewCh<- data
			}

			testyunhuan20201010_2_module1.Reset()
			testyunhuan20201010_2_module2.Reset()
			testyunhuan20201010_2_module3.Reset()

			testyunhuan20201010_2_system1.Reset()
			testyunhuan20201010_2_system2.Reset()
			testyunhuan20201010_2_system3.Reset()
			lock_2.Unlock()
		}
	}()//循环外的system"阀门"

	//处理有matrix级别需求的客户：
	go func (){
		for{
			time.Sleep(time.Duration(2000)*time.Millisecond)

			lock_3.Lock()
			if (testyunhuan20201010_3_module1.SystemName == "智能机柜"){
				testyunhuan20201010_3_system2.AppendModule(testyunhuan20201010_3_module1)		
			}

			if (testyunhuan20201010_3_module2.SystemName == "智能机柜"){
				testyunhuan20201010_3_system2.AppendModule(testyunhuan20201010_3_module2)	
			}

			if (testyunhuan20201010_3_module3.SystemName == "智能机柜"){
				testyunhuan20201010_3_system2.AppendModule(testyunhuan20201010_3_module3)	
			}


			if (testyunhuan20201010_3_system1.MatrixName == "矩阵1"){
				testyunhuan20201010_3_matrix1.AppendSystem(testyunhuan20201010_3_system1)	
			}

			if (testyunhuan20201010_3_system2.MatrixName == "矩阵1"){
				testyunhuan20201010_3_matrix1.AppendSystem(testyunhuan20201010_3_system2)	
			}

			if (testyunhuan20201010_3_system3.MatrixName == "矩阵1"){
				testyunhuan20201010_3_matrix1.AppendSystem(testyunhuan20201010_3_system3)	
			}
			

			if data, err := json.Marshal(testyunhuan20201010_3_matrix1);err == nil{
				matrixViewCh<-data
			}

			if data, err := json.Marshal(testyunhuan20201010_3_matrix2);err == nil{
				matrixViewCh<- data
			}

			testyunhuan20201010_3_module1.Reset()
			testyunhuan20201010_3_module2.Reset()
			testyunhuan20201010_3_module3.Reset()

			testyunhuan20201010_3_system1.Reset()
			testyunhuan20201010_3_system2.Reset()
			testyunhuan20201010_3_system3.Reset()

			testyunhuan20201010_3_matrix1.Reset()
			testyunhuan20201010_3_matrix2.Reset()
			lock_3.Unlock()
		}
	}()//循环外的matrix"阀门"

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

	return moduleViewCh, systemViewCh, matrixViewCh, alarmSMSCh, alarmEntityCh
}