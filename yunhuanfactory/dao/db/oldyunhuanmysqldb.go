package dao

import(
	"github.com/ziyouzy/mylib/yunhuanfactory/model/db/model"
)
/*
这里设计的dao层不是为了操作model实体
而是只为了从数据库拿到model实体
*/

func GetNodefromMySql(table_name string) (*OldYunHuanMySqlDB,error) {
	var model OldYunHuanMySqlDB
	DB.Table(table_name).Order("id desc").Limit(1).Scan(&model)
	if model.Id !=""{
		return model,nil
	}else{
		return model, errors.New("Scanf OldYunHuanMySqlDB Msyql Error")
	}
}