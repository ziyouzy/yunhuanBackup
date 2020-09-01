package mysql

import(
	"fmt"
	"time"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var DB *gorm.DB

func Database(connString string) {
	db, err := gorm.Open("mysql", connString)
	db.LogMode(true)
	if err != nil {
		fmt.Println("connect mysql error:",err)
	}
	//设置连接池和空闲连接数/最大连接数/超时时间
	db.DB().SetMaxIdleConns(20)
	db.DB().SetMaxOpenConns(50)
	db.DB().SetConnMaxLifetime(time.Second * 30)
	DB = db
	//migration()
}