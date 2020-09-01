package impl

//reflector_model
type OldNodeReflectorImpl struct {
	Id		string		`gorm:"column:id"`
	Node_name		string		`gorm:"column:node_name"`
	Value		string		`gorm:"column:value"`
	Time		string		`gorm:"column:time"`
	Ip		string		`gorm:"column:ip"`
}

func (this *OldNodeReflectorImpl)InitForHealthTest(){
	this.Id ="1"
	this.Node_name ="494f3031010201"
	this.Value ="494f3031010201XXXXXtimeout|timeout"
	this.Time ="20200829"
	this.Ip ="192.168.10.254"
}

func (this *OldNodeReflectorImpl)InitForHealthTest_494f3031f10101(){
	this.Id ="1"
	this.Node_name ="494f30310f10101"
	this.Value ="494f3031f1010105a28b|::ffff:192.168.10.2"
	this.Time ="20200829"
	this.Ip ="192.168.10.254"
}

func (this *OldNodeReflectorImpl)InitForHealthTest_494f3031f10201(){
	this.Id ="1"
	this.Node_name ="494f3031f10201"
	this.Value ="494f3031f10201009288|:ffff:192.168.10.2"
	this.Time ="20200831"
	this.Ip ="192.168.10.254"
}

func (this *OldNodeReflectorImpl)InitForHealthTest_494f3031110308_1(){
	this.Id ="1"
	this.Node_name ="494f3031110308"
	this.Value ="494f303111030808db7fff7fff00009b83|::ffff:192.168.10.2"
	this.Time ="20200831"
	this.Ip ="192.168.10.254"
}

func (this *OldNodeReflectorImpl)InitForHealthTest_494f3031110308_2(){
	this.Id ="1"
	this.Node_name ="494f3031110308"
	this.Value ="494f303111030808d77fff7fff00005183|::ffff:192.168.10.2"
	this.Time ="20200831"
	this.Ip ="192.168.10.254"
}

func (this *OldNodeReflectorImpl)InitForHealthTest_494f3031110302(){
	this.Id ="1"
	this.Node_name ="494f3031110302"
	this.Value ="494f303111030200007987|::ffff:192.168.10.2"
	this.Time ="20200831"
	this.Ip ="192.168.10.254"
}

func (this *OldNodeReflectorImpl)InitForHealthTest_494f3031110304(){
	this.Id ="1"
	this.Node_name ="494f3031110304"
	this.Value ="494f3031110304ffffffffea66|::ffff:192.168.10.2"
	this.Time ="20200831"
	this.Ip ="192.168.10.254"
}