//该包工厂函数的作用需要与协议包的功能解耦合
//需要分清哪些是协议包的工作内容，哪些是当前包的工作内容
//当前工厂函数的功能边界是
//1.传入数据源是[]byte还是某种结构体(如oldMysqlEntity)
//2.其nodeType参数决定了具体创建哪种类型的对象(youren_di?youren_do?)
//3.其protocolType参数确保了业务流程思路的明确与清晰，也确保了业务的灵活性与可拓展性
package physicalnode

import (
	"bytes"
	//"time"
	//"fmt"
	"encoding/hex"
	"encoding/binary"

	"github.com/ziyouzy/mylib/physicalnode/di"
	"github.com/ziyouzy/mylib/physicalnode/do"
)

type  PhysicalNode interface{
	FullOf()
	SelectHandlerAndTag() (string,string)
	SelectOneValueAndTimeUnixNano(string, string, string) (string, uint64)
}

func NewPhysicalNodeFromBytes(char []byte, tag string, protocolnodetype string) PhysicalNode{
	switch (protocolnodetype){
	case "PROTOCOLDO20200924":
		arr := bytes.Split(char, []byte(" /-/ "));
		/*ip:arr[0]    time:arr[1]    tag:arr[2]    buf:arr[3]*/
		physicalnode := do.DO_YOUREN_USRIO808EWR_20200924{
			ProtocolNodeType :protocolnodetype,
	
			Tag:string(arr[2]),
			Handler:hex.EncodeToString(arr[3][:7]),

			TimeUnixNano:uint64(binary.BigEndian.Uint64(arr[1])),
			Raw:arr[3],
		}
		physicalnode.FullOf()
		return &physicalnode

	case "PROTOCOLDI20200924":
		arr := bytes.Split(char, []byte(" /-/ "));
		/*ip:arr[0]    time:arr[1]    tag:arr[2]    buf:arr[3]*/
		physicalnode := di.DI_YOUREN_USRIO808EWR_20200924{
			ProtocolNodeType :protocolnodetype,

			Tag:string(arr[2]),
			Handler:hex.EncodeToString(arr[3][:7]),

			TimeUnixNano:uint64(binary.BigEndian.Uint64(arr[1])),
			Raw:arr[3],
		}
		physicalnode.FullOf()
		return &physicalnode
	default:
		return nil
	}//switch (nodeType)
}

func RawChToPhysicalNodeCh(rawch chan []byte)chan PhysicalNode{
	physicalnodech :=make(chan PhysicalNode)
	go func(){
		defer close(physicalnodech)
		for raw := range rawch{
			physicalNode :=buildPhysicalNode_PROTOCOLYUNHUAN20200924(raw)
			physicalnodech<-physicalNode
		}
	}()
	return physicalnodech
}




func buildPhysicalNode_PROTOCOLYUNHUAN20200924(char []byte)PhysicalNode{
	arr :=bytes.Split(char,[]byte(" /-/ "))//按照空白分割
	tag :=string(arr[2])
	buf :=arr[3]
	//开始实现协议
	switch (tag){
	case "tcpsocket":
		switch {//s1
		case buf[0]==0x49&&buf[1]==0x4f&&buf[2]==0x30&&buf[3]==0x31:
			switch {//s2
			case buf[4]==0xf1:
				switch {//s3
				case buf[5]==0x01&&buf[6]==0x01:
					return NewPhysicalNodeFromBytes(char, tag, "PROTOCOLDI20200924")
				case buf[5]==0x02&&buf[6]==0x01:
					return NewPhysicalNodeFromBytes(char, tag, "PROTOCOLDO20200924")
				default:
					return nil
				}//s3
			default:
				return nil	
			}//s2
		default:
			return nil		
		}//s1

	// case "serial":
	// 	switch{
	// 	case buf[0]==0xf1:
	// 		switch {//s1
	// 		case buf[1]==0x01&&buf[2]==0x01:
	// 			//fmt.Println("f10101!")
	// 			return physicalnode.NewPhysicalNodeFromBytes(b, tag, "YUNHUAN20200924","DO20200924")
	// 		case buf[1]==0x02&&buf[2]==0x01:
	// 			//fmt.Println("f10201!")
	// 			return physicalnode.NewPhysicalNodeFromBytes(b, tag, "YUNHUAN20200924","DI20200924")
	// 		default:
	// 			return nil
	// 		}//s1
	// 	default:
	// 		return nil
	// 	}//s2
	// //case "localqt":
	// 	//反序列化buf为某种功能结构体(如physicalnode.NewDoorMgr)
	// 	//return 这个结构体
	 default:
	 	return nil	
	}//tag
}
