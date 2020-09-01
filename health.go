package main

import(
	"fmt"
	"os"

	//"github.com/ziyouzy/mylib/db/mysql"
	"github.com/joho/godotenv"

	mysqldbreflectorlv1 "github.com/ziyouzy/mylib/yunhuanfactory/dbreflector/mysql/lv1/impl"
	//dbreflectordao "github.com/ziyouzy/mylib/yunhuanfactory/dao/dbreflector/mysql" 
	entitydao "github.com/ziyouzy/mylib/yunhuanfactory/dao/entity/physicalnode"

	//evo "github.com/ziyouzy/mylib/yunhuanfactory/evolver/entity/physicalnode"
)


func main(){
	/*初始化关系型数据库*/
	godotenv.Load()//初始化全局变量工具
	//mysql.Database(os.Getenv("MYSQLCONF"))
	//fmt.Println(os.Getenv("MYSQLCONF"))
	for i:=0;i<=6;i++{

	// 	//testOld :=new(mysqldbreflectorlv1.OldNodeReflectorImpl)
	// 	//testOld.InitForHealthTest()
	// 	//fmt.Println(testOld)
		
	
		//oldReflectorDao,err :=dbreflectordao.NewMysqlDBReflectorDao("lv1",mysql.DB)
		//oldReflectorEntity :=oldReflectorDao.OldNodeEntityFromMySql("494f3031f10201")
		oldReflectorEntity :=new(mysqldbreflectorlv1.OldNodeReflectorImpl)
		switch (i){
			case 0:
				// oldReflectorEntity.InitForHealthTest()
				// /*开门，关门，烟感，水浸这些属于di，属于f10201*/
				// diEntityDao,err :=entitydao.NewPhysicalNodeEntityDao("DI","lv1",oldReflectorEntity)
				// if err !=nil{
				// 	fmt.Println("err:",err)
				// }
		
				//diEntity :=diEntityDao.CreateNodeEntityFromOldNodeEntity()
				//fmt.Println(diEntity)
				//diEntity.Evolve()
			case 1:
				oldReflectorEntity.InitForHealthTest_494f3031f10101()
			case 2:
				oldReflectorEntity.InitForHealthTest_494f3031f10201()
				/*开门，关门，烟感，水浸这些属于di，属于f10201*/
				diEntityDao,err :=entitydao.NewPhysicalNodeEntityDao("DI","lv1",oldReflectorEntity)
				if err !=nil{
					fmt.Println("err:",err)
				}
		
				diEntity :=diEntityDao.CreateNodeEntityFromOldNodeEntity()
				fmt.Println(diEntity)
				fmt.Println(os.Getenv("Di1"))
				diEntity.Evolve()
			case 3:
				oldReflectorEntity.InitForHealthTest_494f3031110308_1()
			case 4:
				oldReflectorEntity.InitForHealthTest_494f3031110308_2()
			case 5:
				oldReflectorEntity.InitForHealthTest_494f3031110302()
			case 6:
				oldReflectorEntity.InitForHealthTest_494f3031110304()
		}
	}
}
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

		!!!20200823-16:18对于第一，第二步的补充说明!!!
		首先要理清思路，dao层，model层，service层各私其职
		第一部需要实现两个model结构体之间的转化，放在dao层就不太合适了
		因为旧model是通过dao层获得的，让dao层基于他创造的东西再去创造新的东西，这存在着逻辑的不合理
		
		如果让model层来做这件事，则必然发生model层内两个实体之间的平行调用
		如在旧结构体的方法内var一个新结构体的实体，这也是不够优雅的，有些层及混乱

		因此得出结论，第一步需要在service层实现
		
		同理，反向思考，由于第二部只涉及到model/entity结构体(实体)自身的字段值更新操作
		第二步应该在model层实现，更确切的说，是model/entity包的每一个实体，都需要设计一个属于他自身的转换方法

		短信模块日后再去讨论

		对于service层是否应该设计在github库的问题，目前看来，获取model.entity里各个实体的service是一定需要在github库内实现的
		在这之前先检查下分层逻辑吧
		
		也不太对，这个可以放在属于entity的dao层实现
	*/
