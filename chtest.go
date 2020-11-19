//管道往往会围绕着三件事：
//1.管道的创建
//2.管道的生产者
//3.管道的消费者
//生产者和消费者都会是周期性的驱动事件，只是驱动原因不同
//而最大的共通点(特征)是消费者一定会是个基于当前管道自身的for循环
//同时如果消费者放在了生产者之前，就必须在独立的携程内进行才不会造成死锁
//而消费者的自由性就会大很多，形式多变，只要重点去注意且正确的使用遵顼上面生产者的使用原则，就会解决大部分问题
package main

import(
	"fmt"
	"time"
)


func main(){
	ch1 :=make(chan int)
	//go func(){
		for i := range ch1{
		fmt.Println("i=",i)
		}
	//}()


	go func(){
		i :=0
		for{
			ch1<-i
			i++
			time.Sleep(time.Second)
		}
	}()

	for{}


}