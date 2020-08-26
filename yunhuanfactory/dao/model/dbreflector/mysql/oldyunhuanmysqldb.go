package db

import(
	"github.com/ziyouzy/mylib/db/mysql"
	"github.com/ziyouzy/mylib/yunhuanfactory/model/db"
	"errors"
)
/*
这里设计的dao层不是为了操作model实体
而是只为了从数据库拿到model实体
*/

func GetNodefromMySql(table_name string) (*db.OldYunHuanMySqlDB,error) {
	var model db.OldYunHuanMySqlDB
	mysql.DB.Table(table_name).Order("id desc").Limit(1).Scan(&model)
	if model.Id !=""{
		return &model,nil
	}else{
		return &model, errors.New("Scanf OldYunHuanMySqlDB Msyql Error")
	}
}