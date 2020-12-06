package do

import(
	"strings"
	"strconv"
	//"bytes"
	//"fmt"
	//"encoding/binary"
	"github.com/imroc/biu"
)

type DO_YOUREN_USRIO808EWR_20200924 struct{
	NodeType string
	ProtocolType string

	TimeUnixNano int64
	Raw []byte
	//Mark string

	Handler string
	Tag string

	DO1 string 
	DO2 string
	DO3 string
	DO4 string
	DO5 string
	DO6 string
	DO7 string
	DO8 string
}

func (p *DO_YOUREN_USRIO808EWR_20200924)FullOf(){
	if bytes.Contains(p.Raw, []byte("timeout")){
		p.DO8 ="timeout"
		p.DO7 ="timeout"
		p.DO6 ="timeout"
		p.DO5 ="timeout"
		p.DO4 ="timeout"
		p.DO3 ="timeout"
		p.DO2 ="timeout"
		p.DO1 ="timeout"
		return
	}
	
	var binStr string
	if bytes.Index(p.Raw,{0x49,0x4f})==0&&strings.Compare(p.Tag,"tcpsocket")==0{
		binStr =biu.BytesToBinaryString(p.Raw[12:16])
	}else if strings.Compare(p.Tag,"serial")==0{
		binStr =biu.BytesToBinaryString(p.Raw[8:12])
	}

	if len(binStr) ==8{
		p.DO8 =binStr[0]
		p.DO7 =binStr[1]
		p.DO6 =binStr[2]
		p.DO5 =binStr[3]
		p.DO4 =binStr[4]
		p.DO3 =binStr[5]
		p.DO2 =binStr[6]
		p.DO1 =binStr[7]
	}else{
		p.DO8 = "undefined"
		p.DO7 = "undefined"
		p.DO6 = "undefined"
		p.DO5 = "undefined"
		p.DO4 = "undefined"
		p.DO3 = "undefined"
		p.DO2 = "undefined"
		p.DO1 = "undefined"
	}

	return

	// if strings.Compare(p.Tag,[]byte("serial"))==0{
	// 	binStr :=biu.BytesToBinaryString(p.Raw[8:12])
	// 	//hex,err := strconv.ParseInt(p.Value[8:12],16,0)
	// 	if len() ==nil{
	// 		tempStr :=string([]byte(strconv.FormatInt(hex,2)[1:]))
	// 		p.DO8 =string([]byte(tempStr)[0:1])
	// 		p.DO7 =string([]byte(tempStr)[1:2])
	// 		p.DO6 =string([]byte(tempStr)[2:3])
	// 		p.DO5 =string([]byte(tempStr)[3:4])
	// 		p.DO4 =string([]byte(tempStr)[4:5])
	// 		p.DO3 =string([]byte(tempStr)[5:6])
	// 		p.DO2 =string([]byte(tempStr)[6:7])
	// 		p.DO1 =string([]byte(tempStr)[7:8])
	// 		return
	// 	}
	// }


}

// func (p *DO_YOUREN_USRIO808EWR_20200924)GetNodeType() string{
// 	return p.NodeType
// }

// func (p *DO_YOUREN_USRIO808EWR_20200924)GetHandler() string{
// 	return p.Handler
// }

func (p *DO_YOUREN_USRIO808EWR_20200924)SelectHandlerAndTag() (string, string){
	return p.Handler, p.Tag
}

// func (p *DO_YOUREN_USRIO808EWR_20200924)GetRaw() (string,string,string,string,string,string,string){
// 	return p.NodeType, p.ProtocolType, p.Tag, p.InputTime, p.Value, p.Mark, p.Handler
// }


func (p *DO_YOUREN_USRIO808EWR_20200924)SelectOneValueAndTime(nodedohandler string, nodedotag string, nodedoname string) (string,int64){
	if strings.Compare(p.Handler,nodedohandler)!=0||strings.Compare(p.Tag, nodedotag)!=0{
		return "",""
	}
	switch (nodedoname){
	case "do8":
		return p.DO8,p.TimeUnixNano
	case "do7":
		return p.DO7,p.TimeUnixNano
	case "do6":
		return p.DO6,p.TimeUnixNano
	case "do5":
		return p.DO5,p.TimeUnixNano
	case "do4":
		return p.DO4,p.TimeUnixNano
	case "do3":
		return p.DO3,p.TimeUnixNano
	case "do2":
		return p.DO2,p.TimeUnixNano
	case "do1":
		return p.DO1,p.TimeUnixNano
	default :
		return "",""
	}
}

	/*-----------old:-----------*/
	//创建物理节点结构体实体（非dao结构体实体），其和其所实现的方法都存在于上层目录
	//填充数据后，接收他的返回值是physicalnode.PhysicalNode这一接口类型
	//从而实现封装性
	// var do physicalnode.YouRenDO

	// do.NodeType  =p.NodeType
	// do.Evolver = evolver
	// do.DO_id, do.DO_name, do.DO_value, do.DO_input_time, do.DO_ip =p.dbEntity.GetAll()

	// tempStr :=strings.Split(do.DO_value,"|")[0]
	// if strings.Contains(tempStr, "timeout"){
	// 	do.DO8 ="timeout"
	// 	do.DO7 ="timeout"
	// 	do.DO6 ="timeout"
	// 	do.DO5 ="timeout"
	// 	do.DO4 ="timeout"
	// 	do.DO3 ="timeout"
	// 	do.DO2 ="timeout"
	// 	do.DO1 ="timeout"
	// }else if strings.Index(tempStr,"494f")==0{
	// 	c :=[]byte(tempStr)
	// 	tempStr =string([]byte(c)[12:16])
	// 	hex, _:=strconv.ParseInt(tempStr,16,0)
	// 	tempStr =strconv.FormatInt(hex,2)
	// 	tempStr =string([]byte(tempStr)[1:])

	// 	do.DO8 =/*"顶盖恢复(0为开路;1为闭路)|"+*/string([]byte(tempStr)[0:1])
	// 	do.DO7 =/*"顶盖开启(0为开路;1为闭路)|"+*/string([]byte(tempStr)[1:2])
	// 	do.DO6 =/*"后门开关(0为开路;1为闭路)|"+*/string([]byte(tempStr)[2:3])
	// 	do.DO5 =/*"前门开关(0为开路;1为闭路)|"+*/string([]byte(tempStr)[3:4])
	// 	do.DO4 =/*"散热风扇(0为开路;1为闭路)|"+*/string([]byte(tempStr)[4:5])
	// 	do.DO3 =/*"绿色灯带(0为开路;1为闭路)|"+*/string([]byte(tempStr)[5:6])
	// 	do.DO2 =/*"红色灯带(0为开路;1为闭路)|"+*/string([]byte(tempStr)[6:7])
	// 	do.DO1 =/*"蓝色灯带(0为开路;1为闭路)|"+*/string([]byte(tempStr)[7:8])
	// }else{
	// 	do.DO8 ="err"
	// 	do.DO7 ="err"
	// 	do.DO6 ="err"
	// 	do.DO5 ="err"
	// 	do.DO4 ="err"
	// 	do.DO3 ="err"
	// 	do.DO2 ="err"
	// 	do.DO1 ="err"
	// 	fmt.Println(tempStr)
	// }

	// return &do
//}

