package myvipers


func BuildSingleViper(namewithpathandsuffix string)*SingleViper{
	strs :=strings.Split(namewithpathandsuffix, "/")
	namewithsuffix :=strs[len(strs)-1]
	path :=strings.Trim(namewithpathandsuffix,namewithsuffix)
	//对相对路径的判定与额外操作
	if strings.Compare(path,"./")==0{
		path ="."
	}
	strs =strings.Split(namewithsuffix,".")
	name :=temp[0]
	suffix :=temp[1]
	
	v :=viper.New()
	v.SetConfigName(name) 
	v.AddConfigPath(path)   
	v.SetConfigType(suffix)//json yarm

	err := v.ReadInConfig() // 搜索路径，并读取配置数据
	if err == nil {
		sv :=SingleViper{
			Name : name,
			Path : path,
			Suffix :suffix,
			V:v,
		}

		sv.ConfigIsChange =make(chan bool)

		sv.watching()

		return &sv
	}else{
		fmt.Println("初始化配置json失败,名称、路径、后缀名分别为:",name,path,suffix)
		return nil
	}
}

type SingleViper struct{
	Name string
	Path string
	Suffix string

	V *viper.Viper
	ConfigIsChange chan bool
}

func (p *SingleViper)ListenConfigChange(){
	for{
		select {
		case <-p.ConfigIsChange:
			p.V =viper.New()
			p.V.SetConfigName(p.Name) 
			p.V.AddConfigPath(p.Path)   
			p.V.SetConfigType(p.Suffix)//json yarm

			if err := p.V.ReadInConfig();err == nil {
				p.watching()
			}else{
				fmt.Println("Fatal reset config file:",err)
			}
		}
	}
}

func (p *SingleViper)watching() {
	p.V.WatchConfig()
	p.V.OnConfigChange(func(e fsnotify.Event) {
		fmt.Println("Config file changed:", e.Name)
		err := p.V.ReadInConfig() // 搜索路径，并读取配置数据
		if err == nil {
			p.ConfigIsChange <-true
			return
		}else{
			fmt.Println("Fatal reset config file:",err)
			return
		}
	})
}

