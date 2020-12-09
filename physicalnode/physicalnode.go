//该包工厂函数的作用需要与协议包的功能解耦合
//需要分清哪些是协议包的工作内容，哪些是当前包的工作内容
//当前工厂函数的功能边界是
//1.传入数据源是[]byte还是某种结构体(如oldMysqlEntity)
//2.其nodeType参数决定了具体创建哪种类型的对象(youren_di?youren_do?)
//3.其protocolType参数确保了业务流程思路的明确与清晰，也确保了业务的灵活性与可拓展性
package physicalnode

import (
	"bytes"
	"time"
	"fmt"
	"encoding/hex"
	"encoding/binary"

	"github.com/ziyouzy/mylib/physicalnode/di"
	"github.com/ziyouzy/mylib/physicalnode/do"
)

type  PhysicalNode interface{
	FullOf()
	SelectHandlerAndTag() (string,string)
	SelectOneValueAndTimeUnixNano(string, string, string) ([]byte, uint64)
}

func NewPhysicalNodeFromBytes(char []byte,tag string,protocoltype string,nodetype string) PhysicalNode{
	//fmt.Println("in NewPhysicalNodeFromBytes,nodetype,protocoltype:",nodetype,protocoltype)
	switch (protocoltype){
	case "YUNHUAN20200924":
		switch (nodetype){
		case "DO20200924":
			arr := bytes.Split(char, []byte(" /-/ "));
			/*ip:arr[0]    time:arr[1]    tag:arr[2]    buf:arr[3]*/
			if len(arr[1]) ==7 {/*b[1] =append(b[1],0);*/fmt.Println("do,timeunixnano==7",time.Unix(0, int64(binary.BigEndian.Uint64(append(arr[1],0)))).Format("2006-01-02 15:04:05.000000000"))}
			fmt.Println("DO,",time.Unix(0, int64(binary.BigEndian.Uint64(arr[1]))).Format("2006-01-02 15:04:05.000000000"))
			
			physicalnode := do.DO_YOUREN_USRIO808EWR_20200924{
				NodeType :nodetype,
				ProtocolType:protocoltype,

				TimeUnixNano:uint64(binary.BigEndian.Uint64(arr[1])),
				//hex.EncodeToString(b[3])含义是将一个raw先基于16进制协议转化为16进制数，再将这个16进制数基于utf8协议转化为string
				//但是现在就用不到他了
				Raw:arr[3],
				//Mark:string(b[0]),

				Tag:string(arr[2]),
				Handler:hex.EncodeToString(arr[3][:7]),
			}
			physicalnode.FullOf()
			return &physicalnode
		case "DI20200924":
			arr := bytes.Split(char, []byte(" /-/ "));
			/*ip:arr[0]    time:arr[1]    tag:arr[2]    buf:arr[3]*/
			if len(arr[1]) ==7 {/*arr[1] =append(arr[1],0);*/fmt.Println("di,timeunixnano==7")}
			fmt.Println("DI,",time.Unix(0, int64(binary.BigEndian.Uint64(arr[1]))).Format("2006-01-02 15:04:05.000000000"))
			
			physicalnode := di.DI_YOUREN_USRIO808EWR_20200924{
				NodeType :nodetype,
				ProtocolType:protocoltype,

				TimeUnixNano:uint64(binary.BigEndian.Uint64(arr[1])),
				Raw:arr[3],

				//Mark:string(b[0]),
				Tag:string(arr[2]),
				Handler:hex.EncodeToString(arr[3][:7]),
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
