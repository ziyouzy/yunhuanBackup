package main

import(
	"fmt"
	_ "github.com/ziyouzy/mylib/serialprocess"
	_ "github.com/ziyouzy/mylib/serialprocess/impl"
	_ "github.com/ziyouzy/mylib/yunhuanfactory/dao/db/dao"
)


func main(){
	fmt.Println("nice")
}