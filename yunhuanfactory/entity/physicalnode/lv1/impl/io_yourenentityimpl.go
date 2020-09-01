package impl

import (
	physicalnodeentity"github.com/ziyouzy/mylib/yunhuanfactory/entity/physicalnode"
	"fmt"
)
type IO_YouRenEntityImpl struct{
	Di physicalnodeentity.IYouRenEntity
	Do physicalnodeentity.IYouRenEntity
}

func (this *IO_YouRenEntityImpl)Evolve(){
	fmt.Println("prepare evolve")
}