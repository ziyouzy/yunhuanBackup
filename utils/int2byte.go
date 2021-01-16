//不该这么用，对于普通的字段变量，直接用binary.BigEndian结构体或binary.LittleEndian结构体的方法去实现即可
package utils

import(
	"bytes"
	"encoding/binary"
)

func IntToBytes(n int) []byte{
	x :=int64(n)
	bytesBuffer := bytes.NewBuffer([]byte{})
	binary.Write(bytesBuffer, binary.BigEndian, x)
	return bytesBuffer.Bytes()
}

func BytesToInt64(b []byte) int64{
	bytesBuffer := bytes.NewBuffer(b)
	var x int
	binary.Read(bytesBuffer, binary.BigEndian, &x)
	return int64(x)
}