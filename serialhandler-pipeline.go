package main

import (
	"fmt"
	"net"
	"sync"
	"strings"
	
	"github.com/ziyouzy/mylib/serial"
)




type pipelineSerialHandler struct{
	ConnMap map[string]*serial.PipelineSerialConn

	ProtocolPortsName []string
	ProtocolPortsBaud []int
	ProtocolPortsReadTimeout []int
	ProtocolPortsNeedCRC []bool
}


func (p *pipelineSerialHandler)ListenAndGenerateRecvCh()chan([]byte){	
	p.ConnMap =make(map[string]*serial.PipelineSerialConn)
	ch := make(chan []byte)

	collectOneClientMsg := func(oneClientCh chan []byte,portname string) {
		_ =portname
		for b := range oneClientCh{/*阀门*/
			ch <-b
			//fmt.Println("设备在线：",p.ConnMap[strings.Split(ip,":")[0]])
		}
	}

	for key, portname := range p.ProtocolPortsName{
		portname,serialconn :=serial.NewPipelineSerialConn(portname,p.ProtocolPortsBaud[key],
			p.ProtocolPortsReadTimeout[key],p.ProtocolPortsNeedCRC[key])

		p.ConnMap[portname] =serialconn
		go collectOneClientMsg((*serialconn).GenerateRecvCh(),portname)
	}

	return ch
}