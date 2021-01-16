//服务端的listen操作需要持久话到程序结束，同时所生成的总管道也是如此
package connserver

import (
	"fmt"
	//"net"
	//"time"
	//"sync"
	//"strings"
	
	"github.com/ziyouzy/mylib/connserver/con"
)


var cs *ConnServer
//tcphandler模块的key为ip地址(无端口号)，可与sendmsg对象无缝对接
//后期设计sendcontroller时可以在对象结构体内部将二者组合
//但是需要先完成将tcphandler重构成connhandler这一步
//或者说，ticketsender和connhander做组合，同时door和connhander也作组合
//connhander是个底层，等同于nodedo是alarmbuilder和nodedobuilder的底层一样

//他似乎也需要存在一个全局实体
type ConnServer struct{
	ConnClientMap map[string]con.Con
	FanInRawCh chan []byte //不需要额外设计close事件，而是与程序自身一起开启与关闭
}

func (p *ConnServer)collectClientRecvCh(clientrecvch chan []byte, key string){
	go func(){
		defer delete(p.ConnClientMap,key)
		defer fmt.Println("该设备管道已经关闭，系统将自动认为该连接以关闭，连接将会从ConnClientMap中删除，设备名：",key)

		for b := range clientrecvch{
			p.FanInRawCh<-b
		}
	}()
}

func ClientMap()map[string]con.Con{return cs.ClientMap()}
func (p *ConnServer)ClientMap()map[string]con.Con {
	return p.ConnClientMap
}

func ListenAndCollect(){cs = new(ConnServer);        cs.ConnClientMap =make(map[string]con.Con);        cs.FanInRawCh =make(chan []byte);        cs.ListenAndCollect()}
func (p *ConnServer)ListenAndCollect(){
	p.TcpListenAndCollect("6668")
	//p.SnmpListenAndCollect("192.168.x.x", "161")
}

func RawCh()chan []byte{return cs.RawCh()}
func (p *ConnServer)RawCh()chan []byte{
	return p.FanInRawCh
}

func Destory(){ cs.Destory() }
func (p *ConnServer)Destory(){
	defer close(p.FanInRawCh)

	for key, client := range p.ConnClientMap{
		if client !=nil { p.ConnClientMap[key].Destory() }
		delete(p.ConnClientMap,key)
	}

}