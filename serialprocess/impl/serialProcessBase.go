package impl

import (
    "fmt"
    "github.com/tarm/serial"
    "errors"
)
//无状态通信模块
type serialProcessBase struct {
    serialClient *serial.Port // 串口客户端
}

/// 构造函数
func NewSerialService(Name string, Baud int) *serialProcessBase {
    c := serial.Config{Name: Name, Baud: Baud}
    //打开串口
    s, err := serial.OpenPort(&c)

    if err != nil {
        fmt.Println(err)
    }
    return &serialProcessBase{serialClient: s}
}

// 读取特定长度的字节数据
func (this *serialProcessBase) read(len int, fullBuf *[]byte) error {
    if len <= 0 {
        return errors.New("len<=0")
    }
    l := 0
    for {
        var buf = make([]byte, 1)
        _, err := this.serialClient.Read(buf)
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
func (this *serialProcessBase) ReadLine(len int) {
    var buf []byte
    err := this.read(len, &buf)
    if err != nil {
        //continue
    }
    fmt.Println(buf)
}

func (this *serialProcessBase) SendLine(arr []byte) {
    fmt.Printf("发送数据：%q\n", arr)
    _, err := this.serialClient.Write(arr)
    if err != nil {
    	//log.Fatal(err)
    }
}

