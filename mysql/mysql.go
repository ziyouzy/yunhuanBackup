package mysql

import(
	"fmt"
	"time"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var DB *gorm.DB

func ConnectMySQL(connString string) {
	db, err := gorm.Open("mysql", connString)
	db.LogMode(true)
	if err != nil {
		fmt.Println("数据库连接失败:",err)
	}
	//设置连接池和空闲连接数/最大连接数/超时时间
	db.DB().SetMaxIdleConns(20)
	db.DB().SetMaxOpenConns(50)
	db.DB().SetConnMaxLifetime(time.Second * 30)

	DB = db

	/*-----------------------------------*/
	/** 初始化时自动创建不存在的表，也被称作自动迁移
	  * 也就是自动的表结构迁移，只会创建表，补充缺少的列，缺少的索引
	  * 但并不会更改已经存在的列类型，也不会删除不再用的列，这样设计的目的是为了保护已存在的数据
	  * 可以同时针对多个表进行迁移设置
	  */
	DB.AutoMigrate(&Alarm{})
	/*-----------------------------------*/
}