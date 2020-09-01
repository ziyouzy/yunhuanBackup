package mysql

import(
	"errors"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"

	//"github.com/ziyouzy/mylib/yunhuanfactory/dao/dbreflector"
	//"github.com/ziyouzy/mylib/yunhuanfactory/model/dbreflector/mysql"
	//olddbreflector  "github.com/ziyouzy/mylib/yunhuanfactory/dbreflector/mysql"
	//mysql_reflector_lv2 "github.com/ziyouzy/mylib/yunhuanfactory/model/dbreflector/mysql/impl/lv2"
	//mysql_reflector_lv3 "github.com/ziyouzy/mylib/yunhuanfactory/model/dbreflector/mysql/impl/lv3"
	dbreflectordaolv1 "github.com/ziyouzy/mylib/yunhuanfactory/dao/dbreflector/mysql/lv1/impl"
)

/*工厂函数可以用来实现版本控制与依赖注入,
imysql的位置为github.com/ziyouzy/mylib/yunhuanfactory/dao/dbreflector
是同一层不需要import
*/

func NewMysqlDBReflectorDao(edition string, db *gorm.DB) (IOldNodeReflectorDao,error) {
	switch (edition){
	case "lv1":
		dao :=dbreflectordaolv1.OldNodeReflectorDaoImpl{
			Edition:edition,
			Db: db,
		}
		return &dao,nil
	default:
		err :=errors.New("未知的api版本号")
		return nil, err
	// case "lv2":
	// 	dao :=mysql_reflector_lv2.mySqlReflector{
	// 		db: db,
	// 	}
	// 	return dao
	// case "lv3":
	// 	dao :=mysql_reflector_lv3.mySqlReflector{
	// 		db: db,
	// 	}
	// 	return dao
	}
}