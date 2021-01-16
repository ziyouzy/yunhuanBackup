package main 

import (
	"fmt"
	//"errors"
)

type S struct{
	a int
	b int
}

type newS S

var (
	x int
	Y int
	s S
	NewS newS
)

var z int =1
var Z int
//Z=3

//zz :=123


var s2 S =S{a :1,b:2}

func (p newS)test(){
	fmt.Println("mytest")
}

type I interface{
	test()
}


func main() {
	var aa,bb,cc string; aa,bb,cc="aa","bb","cc"
	var ZX int =2
	fmt.Println(ZX,z,Z)

	fmt.Println(s2,aa,bb,cc)
}