package protocol

//扇出包括sms在内的所有管道
func ProtocolPrepareFinalJsonMgr_YunHuan20200924(confNodeCh conf.ConfNode chan)(nodeViewCh chan []byte,
	moduleViewCh chan []byte,systemViewCh chan []byte,nodeViewCh chan []byte,alarminmysqlCh []string,smsCh chan []byte){
	//"matrix":"矩阵红的每个元素（index；智能机柜test；智能机柜2；冷通道1）都可以单独用vipe取出并形成独立的matrix并序列化后通过websokcet发送",
	//"matrix":{  
	//"index":{"mode":"PUE","systems":["智能机柜test","智能机柜2","冷通道1","冷通道test"]},
	//"智能机柜test":["环境监测1","环境监测2"],
	//"智能机柜2":null,
	//"冷通道1":null,
	//"冷通道test":["环境监测1","环境监测2","环境监测3"]
	//},

	nodeViewCh := make(chan []byte)
	
	moduleViewCh := make(chan []byte)
	module1 :=conf.Module{
		Name: "环境监测1",
	}
	module2 :=conf.Module{
		Name: "环境监测2",
	}
	go func (){
		defer close(moduleViewCh)
		tick := time.Tick(1 * time.Minute)
		for _ = range tick {
			moduleViewCh<- json(module1)
			clear(module1)
			moduleViewCh<- json(module2)
			clear(module2)
		}
	}


	systemViewCh := make(chan []byte)
	system1 :=conf.Module{
		Name: "智能机柜1",
	}
	go func (){
		defer close(systemViewCh)
		tick := time.Tick(1 * time.Minute)
		for _ = range tick {
			moduleViewCh<- json(system1)
			clear(system1)
		}
	}

	matrixViewCh := make(chan []byte)
	//本协议不需发送matrix级别的数据

	sms :=[]byte
	smsCh := make(chan []byte)
	go func (){
		defer close(smsCh)
		tick := time.Tick(1 * time.Minute)
		for _ = range tick {
			smsCh<- sms
			clear(sms)
		}
	}

	
	alarmmsg :=[]byte
	dbAlarmCh :=make(chan byte)
	go func (){
		defer close(dbAlarmCh)
		tick := time.Tick(1 * time.Minute)
		for _ = range tick {
			dbAlarmCh<- alarmmsg
			clear(alarmmsg)
		}
	}
	

	for {开始监听conf.ConfNode管道}
	if (strings.Compare(node.System,"")==0&&strings.Compare(node.Module,"")==0&&strings.Compare(node.Matrix,"")==0){
		nodevViewCh<-node.JSON//no ticket	
	}

	if (strings.Compare(node.System,"")==0&&strings.Compare(node.Module,"")!=0&&strings.Compare(node.Matrix,"")==0){
		go func(){
			switch (node.Module){
			case "环境监测1":
				module1.AppendNode(json(node)
			case "环境监测2":
				module2.AppendNode(json(node)
			default:			
			}
		}()
	}

	if (strings.Compare(node.System,"")!=0&&strings.Compare(node.Module,"")!=0&&strings.Compare(node.Matrix,"")==0){
		switch (node.System){
		case "冷通道test":
			switch (node.Module){
			case "环境监测1":
				module1.AppendNode(json(node)
				system1.AppendModule(json(module1))
			case "环境监测2":
				module2.AppendNode(json(node)
				system1.AppendModule(json(module2))
			default:	
			}
		case "智能机柜test":
			switch(node.Module){
			case "环境监测1":
				module3.AppendNode(json(node)
				system2.AppendModule(json(module3))
			case "环境监测2":
				module4.AppendNode(json(node)
				system2.AppendModule(json(module4))
			case "环境监测3":
				module5.AppendNode(json(node)
				system2.AppendModule(json(module5))
			default:	
			}
		}
	}

	if (strings.Compare(node.System,"")!=0&&strings.Compare(node.Module,"")!=0&&strings.Compare(node.Matrix,"")!=0){
		//当前协议无需求
	}
}