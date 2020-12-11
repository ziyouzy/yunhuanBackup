//服务端的listen操作需要持久话到程序结束，同时所生成的总管道也是如此
package connserver

import (
	//"fmt"
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
	RawCh chan []byte //不需要额外设计close事件，而是与程序自身一起开启与关闭
}

func ClientMap()map[string]con.Con{return cs.ClientMap()}
func (p *ConnServer)ClientMap()map[string]con.Con {
	return p.ConnClientMap
}

func LoadSingletonPatternListenAndCollect(){cs.ConnClientMap =make(map[string]con.Con);        cs.RawCh =make(chan []byte);        cs.ListenAndCollect()}
func (p *ConnServer)ListenAndCollect(){
	p.TcpListenAndCollect(":6668")
	//p.InitSnmp(":161")
}

func RawCh()chan []byte{return cs.RawCh}
func (p *ConnServer)RawCh()chan []byte{
	return p.RawCh
}

func Quit(){ cs.Quit() }
func (p *ConnServer)Quit(){
	defer close(RawCh)

	for key, client := range p.ConnClientMap{
		if client !=nil { p.ConnClientMap[key].Quit() }
		delete(p.ConnClientMap,key)
	}

}