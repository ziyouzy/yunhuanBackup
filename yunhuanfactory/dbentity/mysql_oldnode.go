//此源文件包含OldMySQLNode实体的结构体
//该结构体所实现的接口存在于同级目录下的dbentity.go源文件
//这里说明一下gorm的反射操作，可以确认数据库中的id是int类型，因为他是自动增益的主键
//之前在使用grom的反射操作时，如下OldMySQLNode.id为string类型
//因此可以得出结论，gorm可以自动完成有效的数据转换，如数据库中为int，取出后变为了string
//但是其他的字段就不建议借助gorm进行隐式转换了
//如value的格式很复杂
//而time在数据库内是string类型，如果他是datatime，则会比较放心
//所以就这样吧，只把id设置成int

package dbentity


type OldMySQLNode struct {
	id		int		`gorm:"column:id"`
	nodeName		string		`gorm:"column:node_name"`
	value		string		`gorm:"column:value"`
	time		string		`gorm:"column:time"`
	ip		string		`gorm:"column:ip"`
}

func (p *OldMySQLNode)GetId() int{
	return p.id
}

func (p *OldMySQLNode)GetNodeName() string{
	return p.nodeName
}

func (p *OldMySQLNode)GetValue() string{
	return p.value
}

func (p *OldMySQLNode)GetTime() string{
	return p.time
}

func (p *OldMySQLNode)GetIp() string{
	return p.ip
}

func (p *OldMySQLNode)GetAll() (int,string,string,string,string){
	return p.id, p.nodeName, p.value, p.time, p.ip
}


func (p *OldMySQLNode)InitForHealthTest(){
	p.id = 1
	p.nodeName ="494f3031f10201"
	p.value ="494f3031f1010201XXXXXtimeout|timeout"
	p.time ="20200829"
	p.ip ="192.168.10.254"
}

func (p *OldMySQLNode)InitForHealthTest_494f3031f10101(){
	p.id = 1
	p.nodeName ="494f3031f10101"
	p.value ="494f3031f1010105a28b|::ffff:192.168.10.2"
	p.time ="20200829"
	p.ip ="192.168.10.254"
}

func (p *OldMySQLNode)InitForHealthTest_494f3031f10201(){
	p.id = 1
	p.nodeName ="494f3031f10201"
	p.value ="494f3031f10201009288|:ffff:192.168.10.2"
	p.time ="20200831"
	p.ip ="192.168.10.254"
}

func (p *OldMySQLNode)InitForHealthTest_494f3031110308_1(){
	p.id = 1
	p.nodeName ="494f3031110308"
	p.value ="494f303111030808db7fff7fff00009b83|::ffff:192.168.10.2"
	p.time ="20200831"
	p.ip ="192.168.10.254"
}

func (p *OldMySQLNode)InitForHealthTest_494f3031110308_2(){
	p.id = 1
	p.nodeName ="494f3031110308"
	p.value ="494f303111030808d77fff7fff00005183|::ffff:192.168.10.2"
	p.time ="20200831"
	p.ip ="192.168.10.254"
}

func (p *OldMySQLNode)InitForHealthTest_494f3031110302(){
	p.id = 1
	p.nodeName ="494f3031110302"
	p.value ="494f303111030200007987|::ffff:192.168.10.2"
	p.time ="20200831"
	p.ip ="192.168.10.254"
}

func (p *OldMySQLNode)InitForHealthTest_494f3031110304(){
	p.id = 1
	p.nodeName ="494f3031110304"
	p.value ="494f3031110304ffffffffea66|::ffff:192.168.10.2"
	p.time ="20200831"
	p.ip ="192.168.10.254"
}