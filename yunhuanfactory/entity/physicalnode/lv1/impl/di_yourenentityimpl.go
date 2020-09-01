package impl

import(
	"fmt"
	evo "github.com/ziyouzy/mylib/yunhuanfactory/evolver/entity/physicalnode"
)

type Di_YouRenEntityImpl struct{
	Di_id string	//其所处数据库内所在表对应行的id号
	Di_name string	//唯一标识，很重要，之后很多功能都需要通过他来实现
	Di_input_time string
	Di_value string

	Di1 string 
	Di2 string
	Di3 string
	Di4 string
	Di5 string
	Di6 string
	Di7 string
	Di8 string
}

/*目前的问题在于，我不可能每次evolve时都去硬盘调用一次配置文件*/
/*更确切的说，配置文件应该只有在程序启动时调用一次*/
/*所以说，无论如何，还是需要设计依赖注入*/
/*同时evolve包需要和gorm.DB一样不能被销毁*/
	
	//传入Evolver的接口
func (this *Di_YouRenEntityImpl)Evolve(){
	fmt.Println("prepare evolve")

	Ev := evo.NewPhysicalNodeEntityEvolver("lv1")
	this.Di1 =Ev.Evolver("Di1",this.Di1)
	fmt.Println("Last Di1:",this.Di1)

	this.Di2 =Ev.Evolver("Di2",this.Di2)
	fmt.Println("Last Di2:",this.Di2)
}
