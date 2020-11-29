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
	//"github.com/ziyouzy/mylib/do" 
	"github.com/ziyouzy/mylib/physicalnode"
	"github.com/ziyouzy/mylib/mysql"

	"fmt"
	//"encoding/json"
	//"sync"
	//"time"
)



//传入的参数不再会是NodeDo的管道，而直接变成了PhysicalNode的管道
//我也考虑了要不要先用NodeDoCache把PhysicalNode转化为Nodedo
//但是似乎不太合理，因为应该有个工具函数将PhysicalNode分别扇出为NodeDo的管道，短信告警的管道，还有告警录入数据库的管道
//这个函数似乎既可以属于主函数，也可以属于protocol包，考虑到以后每个项目到了这里很可能都会有巨大的不同，所以还是先让他属于protocol吧
func ProtocolEnding_YunHuan20201101(physicalnodech chan physicalnode.PhysicalNode)(chan []byte, chan []byte, chan *mysql.Alarm){
	nodeDoBytesCh := make(chan []byte)//用来发送给前端软件,其实是将nodedo的每个实体，先经过判定告警的操作后，再转化为字节数组
	//alarmSMSBytesCh := make(chan []byte)//用来232串口的短信报警，会将string转化为[]byte
	//alarmEntityCh := make(chan *mysql.Alarm)//用来mysql的告警信息录入

	//对physicalnode管道相关的操作
	go func(){
		for pn := range physicalnodech{
			//这里是利用物理节点来渲染NodeDoCache内的各个NodeDo模板
			conf.NodeDoBuilder.Engineing(pn)
		}
	}()


	//拿到NodeDo对象的管道，并基于他是下如下功能：
	//1.拿到后会进入事件驱动模式，触发条件是NodeDoCache的定时发送所有对象实体方法，而非更新某个对象实体的方法
	//2.连锁反映是先判断是否超限，超限的话之后的工作AlarmFilterCanche会自动完成
	//3.将NodeDo对象序列话成字符数组装入管道
	//这里的工作内容似乎应该属于一个service
	//只将filter这一个动作，为其设计一个对应的service，原则上nodeDoBytesCh <-nd.GetJson()不应属于这个service
	//于是服务内部不应包含对nodedoch的创建工作，而是将参数表设计成需要传入单独某个NodeDo的个体
	//那么这个nodedo管道的创建工作就需要在主函数进行了
	//或者再去设计个bytesch转physicalnodech、physicalnodech转nodedoch的一体化service
	//似乎nodedoch的创建以及对其的操作都必须要在主函数完成了

	go func(){
		nodedoch :=conf.NodeDoBuilder.GenerateNodeDoCh()
		for nd := range nodedoch{
			if issafe :=conf.AlarmFilterCache.Filter(nd);!issafe{
				fmt.Println("有NodeDo超限了：",nd)
			}
			nodeDoBytesCh <-nd.GetJson()
		}
	}()

	//conf.AlarmFilterCache.SMSAlarmCh, conf.AlarmFilterCache.MYSQLAlarmCh在conf.AlarmFilterCache时就已经创建成功了
	return nodeDoBytesCh, conf.AlarmFilterCache.SMSAlarmCh, conf.AlarmFilterCache.MYSQLAlarmCh
}