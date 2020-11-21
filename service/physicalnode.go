package service

import(
	"bytes"
	"fmt"
	
	"github.com/ziyouzy/mylib/physicalnode"
)
//1121-11:03 39.39

//创建管道的同时实现了生产者,也是实现了生产者与消费者都在子携程中
func RawChToPhysicalNodeCh(rawch chan []byte)chan physicalnode.PhysicalNode{
	physicalnodech :=make(chan physicalnode.PhysicalNode)
	go func(){
		for raw := range rawch{
			physicalNode :=buildPhysicalNode_PROTOCOL_YUNHUAN20200924(raw)
			fmt.Println("physicalnodech_a")
			physicalnodech<-physicalNode
			fmt.Println("physicalnodech_b")
		}
	}()
	return physicalnodech
}


func buildPhysicalNode_PROTOCOL_YUNHUAN20200924(b []byte)physicalnode.PhysicalNode{
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
