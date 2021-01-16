//是crc16验证工具
//摘自https://www.e-learn.cn/content/qita/583479
//摘自https://github.com/elegantm/ElegantBlog/blob/master/crc-16-demo
//自己设计了函数ValidateModbus()
//因为CheckSum()核心功能函数返回的是一个int16整型，所以对比时采用将末四位强转为int16的思路
//但是也要单独设计一个大小端调换的函数

//0x同样可以理解成语法糖，作用只是为了告诉编辑器后面的是个16进制数
//0x不会占用某个字节的内存空间

//int类型的大小为 8 字节
//int8类型大小为 1 字节
//int16类型大小为 2 字节
//int32类型大小为 4 字节
//int64类型大小为 8 字节

//1个字节＝2个16进制字符：
//byte类型大小为2个"字符"，而不是2个"字节"

package utils

import (
	"fmt"
	"bytes"
	"encoding/binary"
)

//观察如下表
//0x代表了数据类型为16进制数
//使用了2个16进制字符表示这个16进制数，一个字符是4个bit，四个是16个bit，可填满两个字节，同时也代表了4位长度的16进制数
//如"0xC181":
//0x语法糖可忽略，之后的“C”是一个ascii字符，占4个bit，之后的"1"、"8"、"1"同理换算。合在一起四个就是16个bit了，填满2个字节

//一切int数据类型容器都能接收16进制数
//只要他装得下这个16进制数就行
//除了int以外,byte也能装16进制数
//但是只能装下2"个"16进制数,容量等同于uint8
//而两者之间的区别在于uint8的默认输出形式是10进制数,byte则是默认输出ACSII码
//4个bit可以描述一"个"16进制数的由0~F
//8个bit可以描述所有的ASCII码
//在使用方式上,byte和所有类型的int都能接收诸如10进制数,8进制数,16进制数,重点在于不要装不下即可
//如果把这里换成了[]byte,唯一的问题只不过是装不下了
const (
	ISBIGENDDIAN = true
	ISLITTLEENDDIAN = false
)

var MbTable = []uint16{
	0X0000, 0XC0C1, 0XC181, 0X0140, 0XC301, 0X03C0, 0X0280, 0XC241,
	0XC601, 0X06C0, 0X0780, 0XC741, 0X0500, 0XC5C1, 0XC481, 0X0440,
	0XCC01, 0X0CC0, 0X0D80, 0XCD41, 0X0F00, 0XCFC1, 0XCE81, 0X0E40,
	0X0A00, 0XCAC1, 0XCB81, 0X0B40, 0XC901, 0X09C0, 0X0880, 0XC841,
	0XD801, 0X18C0, 0X1980, 0XD941, 0X1B00, 0XDBC1, 0XDA81, 0X1A40,
	0X1E00, 0XDEC1, 0XDF81, 0X1F40, 0XDD01, 0X1DC0, 0X1C80, 0XDC41,
	0X1400, 0XD4C1, 0XD581, 0X1540, 0XD701, 0X17C0, 0X1680, 0XD641,
	0XD201, 0X12C0, 0X1380, 0XD341, 0X1100, 0XD1C1, 0XD081, 0X1040,
	0XF001, 0X30C0, 0X3180, 0XF141, 0X3300, 0XF3C1, 0XF281, 0X3240,
	0X3600, 0XF6C1, 0XF781, 0X3740, 0XF501, 0X35C0, 0X3480, 0XF441,
	0X3C00, 0XFCC1, 0XFD81, 0X3D40, 0XFF01, 0X3FC0, 0X3E80, 0XFE41,
	0XFA01, 0X3AC0, 0X3B80, 0XFB41, 0X3900, 0XF9C1, 0XF881, 0X3840,
	0X2800, 0XE8C1, 0XE981, 0X2940, 0XEB01, 0X2BC0, 0X2A80, 0XEA41,
	0XEE01, 0X2EC0, 0X2F80, 0XEF41, 0X2D00, 0XEDC1, 0XEC81, 0X2C40,
	0XE401, 0X24C0, 0X2580, 0XE541, 0X2700, 0XE7C1, 0XE681, 0X2640,
	0X2200, 0XE2C1, 0XE381, 0X2340, 0XE101, 0X21C0, 0X2080, 0XE041,
	0XA001, 0X60C0, 0X6180, 0XA141, 0X6300, 0XA3C1, 0XA281, 0X6240,
	0X6600, 0XA6C1, 0XA781, 0X6740, 0XA501, 0X65C0, 0X6480, 0XA441,
	0X6C00, 0XACC1, 0XAD81, 0X6D40, 0XAF01, 0X6FC0, 0X6E80, 0XAE41,
	0XAA01, 0X6AC0, 0X6B80, 0XAB41, 0X6900, 0XA9C1, 0XA881, 0X6840,
	0X7800, 0XB8C1, 0XB981, 0X7940, 0XBB01, 0X7BC0, 0X7A80, 0XBA41,
	0XBE01, 0X7EC0, 0X7F80, 0XBF41, 0X7D00, 0XBDC1, 0XBC81, 0X7C40,
	0XB401, 0X74C0, 0X7580, 0XB541, 0X7700, 0XB7C1, 0XB681, 0X7640,
	0X7200, 0XB2C1, 0XB381, 0X7340, 0XB101, 0X71C0, 0X7080, 0XB041,
	0X5000, 0X90C1, 0X9181, 0X5140, 0X9301, 0X53C0, 0X5280, 0X9241,
	0X9601, 0X56C0, 0X5780, 0X9741, 0X5500, 0X95C1, 0X9481, 0X5440,
	0X9C01, 0X5CC0, 0X5D80, 0X9D41, 0X5F00, 0X9FC1, 0X9E81, 0X5E40,
	0X5A00, 0X9AC1, 0X9B81, 0X5B40, 0X9901, 0X59C0, 0X5880, 0X9841,
	0X8801, 0X48C0, 0X4980, 0X8941, 0X4B00, 0X8BC1, 0X8A81, 0X4A40,
	0X4E00, 0X8EC1, 0X8F81, 0X4F40, 0X8D01, 0X4DC0, 0X4C80, 0X8C41,
	0X4400, 0X84C1, 0X8581, 0X4540, 0X8701, 0X47C0, 0X4680, 0X8641,
	0X8201, 0X42C0, 0X4380, 0X8341, 0X4100, 0X81C1, 0X8081, 0X4040}


//计算出某个字节数组的crc16校验码
func CheckSum(data []byte, BigEndian bool) []byte {
	var crc16 uint16
	crc16 = 0xffff
	for _, v := range data {
		n := uint8(uint16(v)^crc16)
		crc16 >>= 8
		crc16 ^= MbTable[n]
	}

	bytesBuffer := bytes.NewBuffer([]byte{})	
	if BigEndian{
		binary.Write(bytesBuffer, binary.BigEndian,&crc16)
	}else{
		binary.Write(bytesBuffer, binary.LittleEndian, &crc16)
	}
	return bytesBuffer.Bytes()
}

//拆分需要验证的modbus码
func MidModbus(data []byte) (bytes []byte,crc []byte){
	l :=len(data)
	return data[:l-2], data[l-2:]
}

//验证一个需要验证的modbus码
func CRCCheck(data []byte, BigEndian bool) bool{
	data,crc := MidModbus(data) 
	if bytes.Equal(CheckSum(data, BigEndian), crc){
		return true
	}else if bytes.Equal(CheckSum(data,!BigEndian),crc){
		fmt.Println("大小端反向了")
		return true
	}else{
		return false
	}
}

