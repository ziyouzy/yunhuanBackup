//这个真的可以直接扔了，但是保留是为了告诫自己一个核心思想
//那就是串口通信的需求场景是特别好的体现“用通信解决数据共享”这一场景的
//所以如果不用管道技术，拿岂不是瞎胡闹
//留作反面典型
package serialprocess

// import (
//     "fmt"
//     "github.com/tarm/serial"
//     "errors"
// )
// //无状态通信模块
// type SerialService struct {
//     SerialP *serial.Port // 串口客户端
// }

// // 读取特定长度的字节数据
// func (p *SerialService) read(len int, fullBuf *[]byte) error {
//     if len <= 0 {
//         return errors.New("len<=0")
//     }
//     l := 0
//     for {
//         var buf = make([]byte, 1)
//         _, err := p.SerialP.Read(buf)
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

// // 数据读取
// func (p *SerialService) ReadLine(len int) {
//     var buf []byte
//     err := p.read(len, &buf)
//     if err != nil {
//         //continue
//     }
//     fmt.Println(buf)
// }

// func (p *SerialService) SendLine(arr []byte) {
//     fmt.Printf("发送数据：%q\n", arr)
//     _, err := p.SerialP.Write(arr)
//     if err != nil {
//     	//log.Fatal(err)
//     }
// }