package db

//model
type OldYunHuanMySqlDB struct {
	Id		string		`gorm:"column:id"`
	Node_name		string		`gorm:"column:node_name"`
	Value		string		`gorm:"column:value"`
	Time		string		`gorm:"column:time"`
	Ip		string		`gorm:"column:ip"`
}
