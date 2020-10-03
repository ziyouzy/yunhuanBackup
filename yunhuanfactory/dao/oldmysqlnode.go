//此源文件包含OldMySQLNodeDao实体的结构体
//该结构体所实现的接口存在于同级目录下的mysql_dao.go源文件

//OldMySQLNodeDao实体会返回dbentity.DBEntity接口的实体
//整个dao包被设计成了dbentity的子包，他是dbentity的功能模块,是dao与entity之间的衔接桥梁
//通过dao的工厂函数创建出dao的实体，这个dao实体通过Create()方法！！！注意！！！他可以创建出entity的接口，而不是结构体
//这个接口也可以说是一类entity的超集，同时已经装配了从数据库获取的数据

//每个dao会与数据库内的某一张表相对应
//其实也因如此，虽然返回的是一个entity的接口，但是其内部数据依然是与同一张表相对应
//从而便于使用json包进行各类数据操作

//后期会基于OldMySQLNodeDao接口体继续设计删除表的方法

package dao

import(
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"

	"github.com/ziyouzy/mylib/yunhuanfactory/dbentity"
	//"errors"
)

type OldMySQLNodeDao struct{
	Db *gorm.DB
}
/*
这里设计的dao层不是为了操作model实体
而是只为了从数据库拿到model实体
*/
func (p * OldMySQLNodeDao)Create(tableName string) *dbentity.DBentity {
	var n dbentity.OldMySQLNode
	p.Db.Table(tableName).Order("id desc").Limit(1).Scan(&n)
	if n.Id !=""{
		return &n
	}else{
		//err =errors.New("Scanf OldYunHuanMySqlDB Msyql Error")
		//日志
		return nil
	}
}

//func (d * OldMySQLNodeDao)Delete(tableName string)