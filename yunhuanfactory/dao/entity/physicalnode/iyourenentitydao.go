package physicalnode

import(
	 physicalnodeentity "github.com/ziyouzy/mylib/yunhuanfactory/entity/physicalnode"
)
type IYouRenEntityDao interface{
	CreateNodeEntityFromOldNodeEntity() physicalnodeentity.IYouRenEntity
}