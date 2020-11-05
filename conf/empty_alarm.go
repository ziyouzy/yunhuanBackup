//这个空文章只是为了做下标识,有可能以后会用到
//不过最主要还是明确说明一下conf模块会包含对alarm模块的初始化
package conf



// type AlarmValueObjectMap map[string]interface{}


// func InitConfMap(){
// 	confNodeMap =make(map[string]interface{})
// 	viper.SetConfigName("riverconf") //  设置配置文件名 (不带后缀)
// 	//viper.AddConfigPath("/workspace/appName/") 
// 	viper.AddConfigPath(".")               // 比如添加当前目
// 	viper.SetConfigType("json")
// 	err := viper.ReadInConfig() // 搜索路径，并读取配置数据
// 	if err == nil {
// 		NodeDoValueObjectMap  =updatemap("nodes")	
// 		AlarmValueObjectMap =updatemap("alarm")
// 		fmt.Println("confAlarmMap in init:",confAlarmMap )
// 		fmt.Println("confNodeMap in init:", confNodeMap )
// 		go watching()
// 	}else{//if err == nil
// 		panic(fmt.Errorf("Fatal init config file! \n"))
// 	}//if err == nil end
// }