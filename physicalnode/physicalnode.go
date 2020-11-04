//该包工厂函数的作用需要与协议包的功能解耦合
//需要分清哪些是协议包的工作内容，哪些是当前包的工作内容
//当前工厂函数的功能边界是
//1.传入数据源是[]byte还是某种结构体(如oldMysqlEntity)
//2.其nodeType参数决定了具体创建哪种类型的对象(youren_di?youren_do?)
//3.其protocolType参数确保了业务流程思路的明确与清晰，也确保了业务的灵活性与可拓展性
package physicalnode

import (
	"bytes"
	//"fmt"
	"encoding/hex"
	//"encoding/binary"

	"github.com/ziyouzy/mylib/physicalnode/di"
	"github.com/ziyouzy/mylib/physicalnode/do"
)

type  PhysicalNode interface{
	//GetNodeType() string
	//GetHandler() string
	//GetRaw() (string,string,string,string,string,string,string)
	FullOf()
	SeleteHandlerAndTag() (string,string)
	SeleteOneValueAndTime(string, string, string) (string,string)
}

func NewPhysicalNodeFromBytes(b []byte,tag string,protocoltype string,nodetype string) PhysicalNode{
	//fmt.Println("in NewPhysicalNodeFromBytes,nodetype,protocoltype:",nodetype,protocoltype)
	switch (protocoltype){
	case "YUNHUAN20200924":
		switch (nodetype){
		case "DO20200924":
			b := bytes.Fields(b);
			/*ip:b[0]    time:b[1]    tag:b[2]    buf:b[3]*/
			//hex,err := strconv.ParseInt((p.Value)[12:16],16,0)
			physicalnode := do.DO_YOUREN_USRIO808EWR_20200924{
				NodeType :nodetype,
				ProtocolType:protocoltype,

				Tag:string(b[2]),
				InputTime:string(b[1]),
				Value:string(hex.EncodeToString(b[3])),
				Mark:string(b[0]),

				Handler:string(hex.EncodeToString(b[3][:7])),
			}
			physicalnode.FullOf()
			return &physicalnode
		case "DI20200924":
			b := bytes.Fields(b);
			/*ip:b[0]    time:b[1]    tag:b[2]    buf:b[3]*/
			physicalnode := di.DI_YOUREN_USRIO808EWR_20200924{
				NodeType :nodetype,
				ProtocolType:protocoltype,

				Tag:string(b[2]),
				InputTime:string(b[1]),
				Value:string(hex.EncodeToString(b[3])),
				Mark:string(b[0]),

				Handler:string(hex.EncodeToString(b[3][:7])),
			}
			physicalnode.FullOf()
			return &physicalnode
		default:
			return nil
		}//switch (nodeType)
	default:
		return nil
	}//switch (protocolType)
}
