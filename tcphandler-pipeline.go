//此pakcage负责对tcpSocket的listen以及将客户端的Conn放入ConnMap
//而Conn使用的是对net.Conn经过了当前业务需求封装的PipelineTcpSocketConn
//该核心函数ListenAndGenerateRecvCh()在整体程序生命周期内只会执行也次
//因此比较适合被once修饰，但是似乎没有必要
//因为逻辑较简单，并不需要去刻意担心他会被二次执行
package main

import (
	"fmt"
	"net"
	"sync"
	"strings"
	
	"github.com/ziyouzy/mylib/tcp"
)

var (
	once sync.Once
)


type pipelineTcpHandler struct{
	ConnMap map[string]*tcp.PipelineTcpSocketConn
}


func (p *pipelineTcpHandler)ListenAndGenerateRecvCh()chan([]byte){	
	p.ConnMap =make(map[string]*tcp.PipelineTcpSocketConn)

	tcpAddr, err := net.ResolveTCPAddr("tcp", ":6668")
	if err != nil {
		fmt.Println(err.Error())
	}

	ch := make(chan []byte)
	
	var wg sync.WaitGroup
	collectOneClientMsg := func(oneClientCh chan []byte,ip string) {
		//某一个客户连接的销毁，该事件最终会让这个函数销毁，但不会影响更上层了
		defer wg.Done()
		defer delete(p.ConnMap,ip)
		defer fmt.Println("该设备5秒无应答，连接将会从map中删除：",ip)
		for b := range oneClientCh{/*阀门*/
			ch <-b
			//fmt.Println("设备在线：",p.ConnMap[strings.Split(ip,":")[0]])
		}
	}

	go func(){
		listener, err := net.ListenTCP("tcp", tcpAddr)
		if err != nil {
			fmt.Println(err.Error())
		}	
		defer listener.Close()
		fmt.Println("tcpHandler now is listening")

		//开始接收数据
		for {
			conn, err := listener.Accept()
			//体验一下once的作用，其实并太不需要他
			if err != nil {
				fmt.Println(err.Error())
			}
		
			ip,tcpconn :=tcp.NewPipelineTcpSocketConn(conn,true)
			p.ConnMap[strings.Split(ip,":")[0]] =tcpconn

			go collectOneClientMsg((*tcpconn).GenerateRecvCh(),ip)

			wg.Add(1)
		}

		go func(){
			wg.Wait()
			close(ch)
		}()
	}()

	return ch
}

