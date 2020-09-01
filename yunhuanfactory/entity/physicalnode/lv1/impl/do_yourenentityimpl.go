package impl

import(
	"fmt"
)

type Do_YouRenEntityImpl struct{
	Do_id string	//其所处数据库内所在表对应行的id号
	Do_name string	//唯一标识，很重要，之后很多功能都需要通过他来实现
	Do_input_time string
	Do_value string

	Do1 string
	Do2 string
	Do3 string
	Do4 string
	Do5 string
	Do6 string
	Do7 string
	Do8 string
}

func (this *Do_YouRenEntityImpl)Evolve(){
	fmt.Println("prepare evolve")
}