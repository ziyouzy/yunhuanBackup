// 结合源妹的项目，与当前需求做对比，
// serialProcess与serialProcessBase并不是两个地位完全相等的对象，
// 因此并不适合通过方法封装其业务接口，从而实现上层业务的统一调用
// 而conn其实是个很成功的例子，在这里业务逻辑的复杂程度并不去要可以去实现conn接口，
// 但conn实现了对tcp，udp等等的统一封装，值的反复去学习借鉴他的源代码

//在非容器思路下使用方式是依赖注入到yunhuanfactory/physicalnode/alarm包
//这是目前感觉最为合理的使用方式，或者说对接思路

//20200910不存在serialprocess与serialservice之间的区别，因为串口通信必定是无状态的，所以更绝对不会存在processs
//两者的唯一区别在于有无缓存，无缓存用变量收发数据，有缓存则使用管道
//现决定只保留有缓存的版本

//思考下_, err := p.SerialP.Read(buf)这句
//其实可以理解成是从SerialP中取数据，可以把SerialP理解成他是一个管道(读过的内容自动消失)
//因此在read函数会输入一个预定的长度，达到后返回整体
//因此你准备怎么处理管道内数据，就怎么处理SerialP吧
//或者说，这个问题的本质其实是和前端处理websocket所传来的数据差不多是一样的
//也和你到是取操作IO主控传来的数据是一样的
//都需要基于swtich分支解构业务逻辑
//唯一的区别是serial的switch的判定依据是单个字符
//非serial的switch的判定依据是某个字符串
//而read()方法也是合理可行的方式之一，具体用哪种方式需要根据硬件提供商所给的协议决定
//如杰明的项目，则是判断某个字符是否是“81”这个16进制数
//这个都是需要基于读取每个字符从而实现业务判断的
//目前的项目并不会涉及到，甚至连read（）都不会涉及，所以还是先放置在这吧
package serial

import (
	"github.com/tarm/serial"
    "errors"
    "fmt"
)


//由于目前功能较少，暂时并不需要这个接口，以及其对应的工厂函数
//核心功能的实现还是在于其机构体SerialChan
type SerialProcess interface{
	
}

//设计暴露给调用者的唯一接口
func  NewSerialProcess(portName string, portBaud int,edition string) (SerialProcess, error){
	switch edition {
	case "lv1":
		c := serial.Config{Name: portName, Baud: portBaud}
		s, err := serial.OpenPort(&c)//打开串口
		if err != nil {
			return nil,err
		}else{
			sendCh := make(chan []byte, 1000)
			recvCh := make(chan []byte, 1000)
			//SerialProcess为同级目录下的结构体实体
			return &SerialChan {SP: s, SendCh: sendCh, RecvCh: recvCh,},nil
		}
	
	// case "SERVICE":
	// 	c := serial.Config{Name: portName, Baud: portBaud}
	// 	s, err := serial.OpenPort(&c)//打开串口
	// 	if err != nil {
	// 		return nil,err
	// 	}else{
	// 		//SerialService为同级目录下的结构体实体
	// 		return &SerialService {SerialP: s,},nil
	// 	}

	default:
		return nil,errors.New("创建232串口设备实体时输入了不存在的版本号")
	}
}
/*-------------------上面的暂且用不到，但是包名不变---------------------*/

//管道可以在外部包或者主函数实例化
type SerialChan struct {
    SP *serial.Port // 串口客户端
    SendCh       chan []byte  // 发送通道
    RecvCh       chan []byte  // 接收通道
}




//读取特定长度的字节数据
//之后还会有判定某个字符涉及switch的read方法
//或者说，基于某个字符具体是什么值，判定打包之后多长的字符数组长度的read方法
//比如杰明的项目，如果读到了“81”，则需要打包之后64个字符并返回
func (p *SerialChan) read(len int, fullBuf *[]byte) error {
    if len <= 0 {
        return errors.New(fmt.Sprintf("当用SerialProcess独立调用read()方法时，len<=0,[len=%d]",len))
    }
    l := 0
    for {
        var buf = make([]byte, 1)
        _, err := p.SP.Read(buf)
        if err != nil {
            fmt.Println("err", err)
            return err
        }
        l += 1
        *fullBuf = append(*fullBuf, buf...)
        if l == len {
            return nil
        }

    }
}

// 数据读取
func (p *SerialChan) readThread() {
    go func() {
        for {
            var buf []byte
            err := p.read(16, &buf)
            if err != nil {
                continue
            }
            p.RecvCh <- buf
        }
    }()
}

// 数据发送
//串口通信就是这样，发送的时候可以一句一句发送
//但是接收的时候需要一个一个的接收并判断
func (p *SerialChan) sendThread() {
    go func() {
        for x := range p.SendCh {
            fmt.Printf("发送数据：%q\n", x)
            _, err := p.SP.Write(x)
            if err != nil {
                //log.Fatal(err)
            }
            //time.Sleep(time.Second)
        }
    }()
}

// 数据发送
//并不是很有必要设计这个，因为实现基于某个时间间隔发送数据，完全可以放到报警模块（结构体）去实现
//在报警结构体中会组合serialChan结构体从而实现功能
// func (p *SerialChan) SendByTimerThread(sec int) {
//     go func() {
//     }()
// }

// 启动服务
func (p *SerialChan) Start() {
    go p.readThread()
    go p.sendThread()
}

// 只启动发送服务
func (p *SerialChan) StartSendThread() {
    go p.sendThread()
}

