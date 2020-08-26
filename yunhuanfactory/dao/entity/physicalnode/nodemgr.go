package entity

func NewRAW(string mode, OldYunHuanMySqlDB oldyunhuanmysqldb) error err,i interface{}{
	switch mode{
	// case "IO":
	// 	io := IO
	// 	if err =io.assembleFromOldYunHuanMySqlDB(oldyunhuanmysqldb);err =nil{
	// 		i =io
	// 	}
	// 	return
	case "DO":
		do := Do
		if err =do.assembleFromOldYunHuanMySqlDB(oldyunhuanmysqldb);err =nil{
			i =io
		}
		return
	case "DI":
		di := Di
		if err =di.assembleFromOldYunHuanMySqlDB(oldyunhuanmysqldb);err =nil{
			i =di
		}
		return
	default:
		err =errors.New("在生成IO时输入了未知的mode")
	}
}

