//虽然也是dao，但是只是从oldNode的反射器拿到数据，所以并不需要注入依赖于gorm.DB
//但是需要mode参数，从而确认获取哪种设备

package physicalnode//这个包引用了子目录包/impl/lv1

import(
	"errors"
	"fmt"
	mysqldbreflectorlv1 "github.com/ziyouzy/mylib/yunhuanfactory/dbreflector/mysql/lv1/impl"
	physicalnodeentitydaolv1"github.com/ziyouzy/mylib/yunhuanfactory/dao/entity/physicalnode/lv1/impl"
)

func NewPhysicalNodeEntityDao(nodeType string, edition string,oldNodeReflectors ...*mysqldbreflectorlv1.OldNodeReflectorImpl) (IYouRenEntityDao ,error){
	switch (nodeType){
	case "DO":
		if len(oldNodeReflectors) !=1{
			err :=errors.New(fmt.Sprintf("创建DoDao时，Old反射器实体不为1，而为%v",len(oldNodeReflectors )))
			return nil, err
		}

		switch (edition){
			case "lv1":
				dao :=physicalnodeentitydaolv1.Do_YouRenEntityDaoImpl{
					NodeType :nodeType,
					Edition:edition,
					Old :oldNodeReflectors[0],
				}
				return &dao,nil
			default:
				err :=errors.New("未知的api版本号")
				return nil,err
		}

	case "DI":
		if len(oldNodeReflectors) !=1{
			err :=errors.New(fmt.Sprintf("创建DoDao时，Old反射器实体不为1，而为%v",len(oldNodeReflectors )))
			return nil, err
		}

		switch (edition){
			case "lv1":
				dao :=physicalnodeentitydaolv1.Di_YouRenEntityDaoImpl{
					NodeType :nodeType,
					Edition:edition,
					Old :oldNodeReflectors[0],
				}
				return &dao,nil
			default:
				err :=errors.New("未知的api版本号")
				return nil,err
		}
	
	case "IO":
		if len(oldNodeReflectors) !=2{
			err :=errors.New(fmt.Sprintf("创建DoDao时，Old反射器实体不为1，而为%v",len(oldNodeReflectors )))
			return nil,err
		}

		switch (edition){
			case "lv1":
				dao :=physicalnodeentitydaolv1.IO_YouRenEntityDaoImpl{
					NodeType :nodeType,
					Edition:edition,
					Olds :oldNodeReflectors,
				}
				return &dao,nil
			default:
				err :=errors.New("未知的api版本号")
				return nil,err
			}
	
	default:
		err :=errors.New("在生成IO时输入了未知的mode")
		return nil,err
	}
}