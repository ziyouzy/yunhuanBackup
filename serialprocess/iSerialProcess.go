package serialprocess

import (
	"errors"
	"github.com/ziyouzy/mylib/serialprocess/impl"
)

/*
    结合源妹的项目，与当前需求做对比，
    serialProcess与serialProcessBase并不是两个地位完全相等的对象，
	因此并不适合通过方法封装其业务接口，从而实现上层业务的统一调用

	而conn其实是个很成功的例子，在这里业务逻辑的复杂程度并不去要可以去实现conn接口，
	但conn实现了对tcp，udp等等的统一封装，值的反复去学习借鉴他的源代码
*/


/*这里所设计的接口虽然符合语法规范，但并没有实际的应用价值，只是用来练习*/
type ISerialProcess interface{
	//readThread()
	//sendThread()
	SendLine(arr []byte) 
	ReadLine(len int) 
	//Start()
}

//设计暴露给调用者的唯一接口
func  CreateNewSerial(Name string, Baud int,mode string) (ISerialProcess, error){
	switch mode {
	case "PROCESS":
		return impl.NewSerialProcess(Name,Baud),nil
	case "SERVICE":
		return impl.NewSerialService(Name,Baud),nil
	default:
		return nil,errors.New("创建232串口设备实体时输入了不存在的模式参数")
	}
}

