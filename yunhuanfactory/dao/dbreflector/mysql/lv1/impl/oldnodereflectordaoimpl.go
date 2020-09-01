package impl

import(
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"

	mysqldbreflectorlv1 "github.com/ziyouzy/mylib/yunhuanfactory/dbreflector/mysql/lv1/impl"
	//"errors"
)

type OldNodeReflectorDaoImpl struct{
	Db *gorm.DB
	Edition string
}
/*
这里设计的dao层不是为了操作model实体
而是只为了从数据库拿到model实体
*/
func (this *OldNodeReflectorDaoImpl)OldNodeEntityFromMySql(table_name string) *mysqldbreflectorlv1.OldNodeReflectorImpl {
	var old mysqldbreflectorlv1.OldNodeReflectorImpl
	this.Db.Table(table_name).Order("id desc").Limit(1).Scan(&old)
	if old.Id !=""{
		return &old
	}else{
		//err =errors.New("Scanf OldYunHuanMySqlDB Msyql Error")
		return nil
	}
}