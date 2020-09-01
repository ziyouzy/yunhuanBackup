package impl//这个包交叉引用了父目录包physicalnode

import (
	//"fmt"
	//"errors"
	"strings"
	mysqldbreflectorlv1 "github.com/ziyouzy/mylib/yunhuanfactory/dbreflector/mysql/lv1/impl"

	physicalnodeentitylv1 "github.com/ziyouzy/mylib/yunhuanfactory/entity/physicalnode/lv1/impl"
	physicalnodeentity "github.com/ziyouzy/mylib/yunhuanfactory/entity/physicalnode"
)

type IO_YouRenEntityDaoImpl struct{
	NodeType string
	Edition string

	Olds []*mysqldbreflectorlv1.OldNodeReflectorImpl
}

//获取一个初始的entity实体
func (this *IO_YouRenEntityDaoImpl)CreateNodeEntityFromOldNodeEntity() physicalnodeentity.IYouRenEntity{
	var io physicalnodeentitylv1.IO_YouRenEntityImpl 
	// if len(this.Olds) !=2{
	// 	err :=errors.New(fmt.Sprintf("装配有人io的缓存实体时，输入的节点总数不为两个，而是%v个",len(this.Olds)))
	// 	return nil,err
	// }

	for _,value :=range this.Olds{
		if strings.LastIndex(value.Node_name,"010201")==9{
			didao := Di_YouRenEntityDaoImpl{
				Old: value,
			}
			io.Di=didao.CreateNodeEntityFromOldNodeEntity()
		}

		if strings.LastIndex(value.Node_name,"010101")==9{
			dodao := Do_YouRenEntityDaoImpl{
				Old: value,
			}
			io.Do=dodao.CreateNodeEntityFromOldNodeEntity()
		}
	}

	// if io.Di ==nil&&io.Do==nil{
	// 	err :=errors.New(fmt.Sprintf("装配有人io的缓存实体为空"))
	// 	return nil, err
	// }

	return &io
}