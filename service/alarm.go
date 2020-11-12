package service


func ServiceFilterOneNodeDoIsSafe(nd *NodeDo){
	if issafe :=conf.AlarmFilterCache.Filter(nd);!issafe{
		fmt.Println("有NodeDo超限了：",nd)
	}
}

func SendSerialAlarmSMSToNouthBound(){

}

func CreateAlarmToMYSQL(){

}

// go func(){
// 	nodedoch :=conf.NodeDoController.GenerateNodeDoCh()
// 	for nd := range nodedoch{
		
// 		nodeDoBytesCh <-nd.GetJson()
// 	}
// }()