//对象所维护的vipers map的底层会实时更新所对应的SingleViper
//一旦ConfigIsChange触发文件发生改变也不必惊慌，在ViperBuilder这一层不需要做任何事
//但是他的上层，如nodedobuilder则需要switch一下ConfigIsChange里所传来的路径名，从而判断其自身是否需要进行相应的底层数据更新
package viperbuilder

import(
	"fmt"
)

var builder *ViperBuilder
type ViperBuilder struct{
	vipers map[string]*SingleViper
	ConfigIsChange chan string
}

func Load(paths ...string){
	var i  interface{};    i =paths

	if res, ok :=i.(string);        ok{ builder = BuildViperBuilder([]string{res,});        return }

	if res, ok :=i.([]string);        ok{ builder = BuildViperBuilder(res);        return }

	fmt.Println("执行ViperBuilder.Load()时监测到配置文件路径填写错误")
}

func BuildViperBuilder(paths []string)*ViperBuilder{
	builder :=ViperBuilder{}
	builder.vipers =make(map[string]*SingleViper)
	builder.ConfigIsChange =make(chan string)

	builder.AddSingleVipers(paths)

	return &builder
}

func ConfigListener()chan string{return builder.ConfigListener()}
func (p *ViperBuilder)ConfigListener()chan string{return p.ConfigIsChange}

func (p *ViperBuilder)AddSingleVipers(paths []string){
	for _, path :=range paths{
		if sv :=BuildSingleViper(path); sv!=nil{
			p.vipers[path] =sv
			go func(){
				//底层OneViperConfigIsChangeAndUpdateFinishCh管道是不会关闭的
				//因为当文件改变时，SingleViper只会进行更新操作，而不是重新创建
				//除非某些情况下主动调用SingleViper.Destory()
				for changedJSONName := range sv.OneViperConfigChangedCh{
					p.ConfigIsChange<-changedJSONName
				}
			}()
		}else{	
			fmt.Println("您设置的json路径[",p,"]格式错误，只支持绝对路径与根目录两种模式")
		}
	}
}

func SelectOneMapFromOneSingleViper(singleviperpath string, keyofmap string)map[string]interface{} { return builder.SelectOneMapFromOneSingleViper(singleviperpath, keyofmap)}
func (p *ViperBuilder)SelectOneMapFromOneSingleViper(singleviperpath string, keyofmap string)map[string]interface{}{
	m :=p.vipers[singleviperpath].V.Get(keyofmap)
	if value, ok :=m.(map[string]interface{});ok{
		return value
	}else{
		fmt.Println("SelectOneMap fail, path is:", singleviperpath,"key is:",keyofmap)
		return nil
	}
}

func (p *ViperBuilder)DeleteOneSingleViper(singleviperpath string){
	p.vipers[singleviperpath].Destory()

	if _, exist := p.vipers[singleviperpath]; exist{ delete(p.vipers, singleviperpath) }
}


//不需要这个方法，写出来只是为了强调OneMapFromOneSingleVipe与OneSingleViper是完全不同的两个东西
//func (p *ViperBuilder)DeleteOneMapFromOneSingleVipe(){}


func Destory(){builder.Destory()}
func (p *ViperBuilder)Destory(){
	close(p.ConfigIsChange)
	for key,_ := range p.vipers{
		p.vipers[key].Destory()

		if _, exist := p.vipers[key]; exist{ delete(p.vipers, key) }
	}
}

