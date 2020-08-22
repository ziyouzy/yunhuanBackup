package impl

import (
    "fmt"
    "github.com/tarm/serial"
)

//有状态通信模块
type serialProcess struct {
    serialProcessBase
    SendCh       chan []byte  // 发送通道
    RecvCh       chan []byte  // 接收通道
}


func NewSerialProcess(Name string, Baud int) *serialProcess {
    c := serial.Config{Name: Name, Baud: Baud}
    //打开串口
    s, err := serial.OpenPort(&c)

    if err != nil {
        fmt.Println(err)
    }
    sendCh := make(chan []byte, 1000)
    recvCh := make(chan []byte, 1000)
    return &serialProcess{serialProcessBase:serialProcessBase{serialClient: s}, SendCh: sendCh, RecvCh: recvCh}
}

// 读取特定长度的字节数据
// func (sp *serialProcess) read(len int, fullBuf *[]byte) error {
//     if len <= 0 {
//         return types.Error{}
//     }
//     l := 0
//     for {
//         var buf = make([]byte, 1)
//         _, err := sp.serialClient.Read(buf)
//         if err != nil {
//             fmt.Println("err", err)
//             return err
//         }
//         l += 1
//         *fullBuf = append(*fullBuf, buf...)
//         if l == len {
//             return nil

//         }

//     }
// }

// 数据读取
func (this *serialProcess) readThread() {
    go func() {
        for {
            var buf []byte
            err := this.read(16, &buf)
            if err != nil {
                continue
            }
            this.RecvCh <- buf
        }
    }()
}

// 数据发送
func (this *serialProcess) sendThread() {
    go func() {
        for x := range this.SendCh {
            fmt.Printf("发送数据：%q\n", x)
            _, err := this.serialClient.Write(x)
            if err != nil {
                //log.Fatal(err)
            }
            //time.Sleep(time.Second)
        }
    }()
}

// 启动服务
func (this *serialProcess) Start() {
    go this.readThread()
    go this.sendThread()
}