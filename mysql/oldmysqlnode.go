package mysql

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