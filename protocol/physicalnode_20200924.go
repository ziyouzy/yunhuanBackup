package protocol

import(
	"bytes"
	"github.com/ziyouzy/mylib/physicalnode"
)
//由字符串"反序列化"成physicalnode基础实体
//生成核心物理节点,或者说初始化一个核心物理节点
//之后还需要再经过physicalnode自身的fullof方法完成渲染，才能供下一层使用
func ProtocolPreparePhysicalNode_YunHuan20200924(b []byte)physicalnode.PhysicalNode{
	bufarr :=bytes.Fields(b)//按照空白分割
	tag :=string(bufarr[2])
	buf :=bufarr[3]
	//开始实现协议
	switch (tag){
	case "tcpsocket":
		switch {//s1
		case buf[0]==0x49&&buf[1]==0x4f&&buf[2]==0x30&&buf[3]==0x31:
			switch {//s2
			case buf[4]==0xf1:
				switch {//s3
				case buf[5]==0x01&&buf[6]==0x01:
					return physicalnode.NewPhysicalNodeFromBytes(b, tag, "YUNHUAN20200924","DI20200924")
				case buf[5]==0x02&&buf[6]==0x01:
					return physicalnode.NewPhysicalNodeFromBytes(b, tag, "YUNHUAN20200924","DO20200924")
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