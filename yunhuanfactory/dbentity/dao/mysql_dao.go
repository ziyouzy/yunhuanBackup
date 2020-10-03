//这个接口本质是各个MySQLDao实体(如OldMySQLNodeDao)的超集，这些数据库实体也存在于当前目录下，如mysql_oldnodedao.go文件内

//工厂函数NewMySQLDaoCreater()是最标准的创建dao对象的方式，创建过程中，对版本号和实体类型都需要进行配置
//需要明确需要创建哪种类型的dao，如用来针对mysql_oldnode实体的，还是别的
//所有的dao实体都会包含Create方法，这个方法用来获取对象（dbentity根目录下）的实体
//要记得每个entity都拥有InitForHealthTest()方法，dbentity.go/DBentity是实现了这个方法的实体对象接口
//工厂函数的第一个参数edition是版本号，如lv1,这里的版本号只是选择dao实体的版本号，而不是entity实体的版本号
//第二个参数是选择去创建哪种dao，如oldnodedao,锁定后会投影到数据库，他会与具体数据库中的某一张表或多张表的联合体对应起来
//第三个参数是依赖注入，db是需要依赖注入的gorm.DB实体，是最终所创建的不同dao实体所必须的共同的数据库引擎
//db（gorm.DB的依赖注入）只是dao用来生成其所对应entity所必须的内部字段，而并不是entity所必须的内部字段
//当前源文件下的MySQLDaoCreater接口以及实现这一接口的结构体都只能用来生成mysql的dao，因为依赖注入的参数是*gorm.DB
//edition只会在NewMySQLDaoCreater()这一工厂函数内被使用，版本的控制只在dao的工厂函数内被执行
//entity本身的工厂函数禁止进行对版本的配置，这是为了实现高耦合低内聚思想
//nodeType需要与具体db实体的结构体名称一致

package dao

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"

	"github.com/ziyouzy/mylib/yunhuanfactory/dbentity"
)

type MySQLDaoCreater interface{
	Creater(string) dbentity.DBentity
}

func NewMySQLDaoCreater(edition string, typeName string,db *gorm.DB) MySQLDaoCreater {
	switch (edition){
	case "lv1":
		switch (typeName){
		case "OldMySQLNode":
			dao :=OldMySQLNodeDao {
				Db: db,
			}
			return &dao
		default:
			return nil
		}
	case "lv2":
		return nil
	default:
		return nil
	}
}