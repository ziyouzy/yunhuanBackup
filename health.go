package main

import(
	"fmt"
	"github.com/ziyouzy/mylib/db/mysql"
)


func main(){
	/*初始化关系型数据库*/
	mysql.Database("yunhuan_api:13131313@tcp(127.0.0.1:3306)/DB_yunhuan?charset=utf8")
	fmt.Println("nice")

	/*
		到目前为止，已经可以获取旧io的数据库实体(model/db)了，这是通过dao层实现的
		下一步是将该实体转化为model/entity实体，这需要依靠设计service层去实现
		实体只有一个，但是或许需要分两步去获得完整的实体：
		第一步是只获取诸如DO.Do1="1"这样的实体形式
		第二步再去将DO.Do1的值("1")替换成诸如"名称-内容-异常值-报警内容-是否在线-是否报警"这样的通用字符串
		第三步(也可能是第1.5步)也要去实现触发短信报警的功能模块

		第二部在model层实现，因为该步骤的功能实现只属于其自身的范畴，只属于数据转化需求，不需要去调用其他包，及其他功能模块
		model层确实需要写一个实现计算的业务源文件，用来modbus字符串转化成具体的数值
		先把他只设计成model的一部分，后期根据业务需求的变化再去做更加科学的重构
	*/
}