//该包工厂函数的作用需要与协议包的功能解耦合
//需要分清哪些是协议包的工作内容，哪些是当前包的工作内容
//当前工厂函数的功能边界是
//1.传入数据源是[]byte还是某种结构体(如oldMysqlEntity)
//2.其nodeType参数决定了具体创建哪种类型的对象(youren_di?youren_do?)
//3.其protocolType参数确保了业务流程思路的明确与清晰，也确保了业务的灵活性与可拓展性
package physicalnode

import (

)

type  PhysicalNode interface{
	GetNodeType() string
	GetRaw() (string,string,string,string,string,string)
	FullOf()
}

func NewPhysicalNodeFromBytes(b []byte,nodetype string,protocoltype string) PhysicalNode, error{
	switch (protocoltype){
	case "YUNHUAN20200924":
		switch (nodetype){
		case "DO20200924":
			ip,time,tag,buf :=bytes.Fields(b)
			physicalnode := DO_YOUREN_USRIO808EWR_20200924{
				NodeType :nodetype,
				ProtocolType:protocoltype,

				Tag:string(tag),
				ImportTime:string(time),
				Value:string(buf),
				Ip:string(ip),

				Handler:string(buf[:14]),
			}
			physicalnode.FullOf()
			return &physicalnode
		case "DI20200924":
			ip,time,tag,buf :=bytes.Fields(b)
			physicalnode := DI_YOUREN_USRIO808EWR_20200924{
				NodeType :nodetype,
				ProtocolType:protocoltype,

				Tag:string(tag),
				ImportTime:string(time),
				Value:string(buf),
				Ip:string(ip),

				Handler:string(buf[:14]),
			}
			physicalnode.FullOf()
			return physicalnode
		}//switch (nodeType)
	}//switch (protocolType)
}
