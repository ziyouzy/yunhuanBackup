package protocol

import(
	"github.com/ziyouzy/mylib/conf"
	"github.com/ziyouzy/mylib/view"
	"fmt"

	"encoding/json"
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
	module1 =view.Module{
		ModuleName: "环境监测",
		SystemName:"冷通道",
		MatrixName:"矩阵1",
	}
	module2 =view.Module{
		ModuleName: "ups监测",
		SystemName:"智能机柜",
		MatrixName:"矩阵1",
	}
	module3 =view.Module{
		ModuleName: "zndb监测",
		SystemName: "冷通道1",
		MatrixName: "矩阵2",
	}
)

/*代表一位有system级别需求的客户*/
var (
	system1 =view.System{
		SystemName: "冷通道",
		MatrixName:"矩阵1",
	}
	system2 =view.System{
		SystemName: "智能机柜",
		MatrixName:"矩阵1",
	}
	system3 =view.System{
		SystemName: "冷通道1",
		MatrixName: "矩阵2",
	}
)

/*代表一位有matrix级别需求的客户*/
var (
	matrix1 =view.Matrix{
		MatrixName:"矩阵1",
	}
	matrix2 =view.Matrix{
		MatrixName:"矩阵2",
	}
)

var (
	smsticket =0
	mysqlticket =0
)

//拿到主函数去使用,入参是个实体，返回的是个管道
func ProtocolViewNodesHandler_YunHuan20201004(confnodech chan conf.ConfNode)(chan []byte, chan []byte, chan []byte, chan []byte){
	moduleViewCh := make(chan []byte)
	systemViewCh := make(chan []byte)
	matrixViewCh := make(chan []byte)

	alarmSMSCh := make(chan []byte)

	//这个循环alarmsms和alarmdb也会包括在内
	go func(){
		for confnode := range confnodech{
			if confalarm :=conf.NewConfAlram(confnode);confalarm !=nil{
				if(smsticket ==confalarm.SMSSleepMin){
					go func(){
						for _, sms := range confalarm.SMS{
							alarmSMSCh<-[]byte(sms)
							time.Sleep(time.Duration(500)*time.Millisecond)			
						}
					}()
					smsticket =0
				}
			}//循环内的sms服务结束	

			// if(mysqlticket ==confalarm.mysqlticket){
			// 	go func(){
			// 			//mysql的管道里装的不是[]byte，而是可以映射到数据库的结构体
			// 			time.Sleep(time.Duration(500)*time.Millisecond))			
			// 	}()
			// 	mysqlticket =0
			// }	
			//}

			/*------*/
			matrix, system, module := confnode.GetMatrixSystemAndModuleString()
			//由内(module)而外(matrix)的装配
			//会有重复
			//第一次把某一个confnode塞入一个module，会立刻进行创建system并将该module，以及关于system的链式反应
			//第二次虽然将新的confnode塞入了旧的module，但是塞入system时，会变得重复
			//
			switch (module){
			case "环境监测":
				module1.AppendNode(confnode)
				switch(system){
				case "冷通道":
					system1.AppendModule(module1)
					switch(matrix){
					case "矩阵1":
						matrix1.AppendSystem(system1)
					case "矩阵2":
						matrix2.AppendSystem(system1)
					}
				case "智能机柜":
					system2.AppendModule(module1)
					switch(matrix){
					case "矩阵1":
						matrix1.AppendSystem(system2)
					case "矩阵2":
						matrix2.AppendSystem(system2)
					}
				case "冷通道1":
					system3.AppendModule(module1)
					switch(matrix){
					case "矩阵1":
						matrix1.AppendSystem(system3)
					case "矩阵2":
						matrix2.AppendSystem(system3)
					}
				}
				
			case "ups监测":
				module2.AppendNode(confnode)
				switch(system){
				case "冷通道":
					system1.AppendModule(module2)
					switch(matrix){
					case "矩阵1":
						matrix1.AppendSystem(system1)
					case "矩阵2":
						matrix2.AppendSystem(system1)
					}
				case "智能机柜":
					system2.AppendModule(module2)
					switch(matrix){
					case "矩阵1":
						fmt.Println("!@!@!@!")
						matrix1.AppendSystem(system2)
					case "矩阵2":
						matrix2.AppendSystem(system2)
					}
				case "冷通道1":
					system3.AppendModule(module2)
					switch(matrix){
					case "矩阵1":
						matrix1.AppendSystem(system3)
					case "矩阵2":
						matrix2.AppendSystem(system3)
					}
				}

			case "zndb监测":
				module3.AppendNode(confnode)
				switch(system){
				case "冷通道":
					system1.AppendModule(module3)
					switch(matrix){
					case "矩阵1":
						matrix1.AppendSystem(system1)
					case "矩阵2":
						matrix2.AppendSystem(system1)
					}
				case "智能机柜":
					system2.AppendModule(module3)
					switch(matrix){
					case "矩阵1":
						matrix1.AppendSystem(system2)
					case "矩阵2":
						matrix2.AppendSystem(system2)
					}
				case "冷通道1":
					system3.AppendModule(module3)
					switch(matrix){
					case "矩阵1":
						matrix1.AppendSystem(system3)
					case "矩阵2":
						matrix2.AppendSystem(system3)
					}
				}
			}//循环内的view渲染结束
		}//循环结束
	}()//该线程函数结束

		//处理有module级别需求的客户：
	go func (){
		for{
			data1, _ := json.Marshal(module1)
			moduleViewCh<-data1

			data2, _ := json.Marshal(module2)
			moduleViewCh<- data2

			data3, _ := json.Marshal(module3)
			moduleViewCh<- data3

			module1.Reset()
			module2.Reset()
			module3.Reset()

			time.Sleep(time.Duration(500*6)*time.Millisecond)
		}
	}()//循环外的module"阀门"

	//处理有system级别需求的客户：
	go func (){
		for{
			data1, _ := json.Marshal(system1)
			systemViewCh<-data1

			data2, _ := json.Marshal(system2)
			systemViewCh<- data2

			data3, _ := json.Marshal(system3)
			systemViewCh<- data3

			system1.Reset()
			system2.Reset()
			system3.Reset()
			
			time.Sleep(time.Duration(500*6)*time.Millisecond)
		}
	}()//循环外的system"阀门"

	//处理有matrix级别需求的客户：
	go func (){
		for{
			data1, _ := json.Marshal(matrix1)
			matrixViewCh<-data1

			data2, _ := json.Marshal(matrix2)
			matrixViewCh<- data2

			matrix1.Reset()
			matrix2.Reset()
			
			time.Sleep(time.Duration(500*6)*time.Millisecond)
		}
	}()//循环外的matrix"阀门"

	//实时更新sms的ticket，以分钟为单位
	go func (){
		for{
			smsticket =smsticket+1
			time.Sleep(time.Second)
		}
	}()//循环外的sms发送控制

		//实时更新db的ticket，以分钟为单位
		// go func (){
		// 	for{
		// 		mysqlticket =mysqlticket+1
		// 		time.Sleep(time.Second)
		// 	}
		// }
	return moduleViewCh, systemViewCh, matrixViewCh, alarmSMSCh
}