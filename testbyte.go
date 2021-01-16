package main

import(
	"fmt"
	"bytes"
	//"github.com/imroc/biu"
)

var testbytes = []byte{0x16,0x17,16,17,'\x16','\x17',/*x16, x17,*(编译器无法识别)/ /*'\0x16','\0x17'(编译器无法识别)*/}
var testbytes2 = []byte{0x41,0x42,41,42,'\x41','\x42', 'A','B',/*A,B(无引号编译器无法识别)*/}

var testbytes3 = []rune{0101,65,0x41, 'A','\x41','\101'}
func main(){
	//fmt.Println(testbytes3)
	//fmt.Println(biu.BytesToBinaryString(testbytes3))
	 _ =testbytes
	 _ =testbytes2
	 _ =testbytes3
	 charnil :=[]byte(" ")
	 char0:=[]byte{0x00}
	 buffer :=bytes.NewBuffer([]byte{});
	 _,_ = buffer.Write([]byte("a"));        _,_ = buffer.Write([]byte(charnil));        _,_ = buffer.Write([]byte("b"));        _,_ = buffer.Write([]byte(char0));        _,_ = buffer.Write([]byte("c"))
	 //fmt.Println(bytes.Fields(buffer.Bytes()))

	 char1:=[]byte{0x65,0x56,71,0x86,16,0x32,0x32,32,0x32,0x66,22,0x41,0x32,73,32,73,84}
	 fmt.Println(bytes.Fields(char1))
	 fmt.Println(bytes.Split(char1,[]byte(" ")))
	 fmt.Println("testing the nil:", []byte(" "))
	 fmt.Println("testing the nil:", []byte(""))
	 fmt.Println("testing the nil:", []byte("0"))
	 fmt.Println("testing the nil:", []byte{0})
	 fmt.Println("testing the nil:", []byte{0x00})
	 fmt.Println("testing the nil:", []byte(nil))
}