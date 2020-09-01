


/*对应了Pod里每一个仪表盘*/
type OldSimpleYunHuanNode struct{
	/*并不是entity.physicalnode内的节点结构体副本，
		而是从entity.physicalnode内的节点结构体中获取所需的字段与字段值后
		将他们装配成这个结构体
	*/

	/*仪表盘类型，用来区别诸如普通数据显示与pue数据显示*/
	model string

	/*数据是否超限或异常*/
	isred bool 

	/*仪表盘下方文字（如：温度，市电电压，前门状态）*/
	theme string

	/*仪表盘内数值，(如25.7%，开/关，正常/异常)*/
	date string

	/*不在这里写入仪表盘素材文件的具体路径
	前端根据model类型去实现具体选择与加载
	*/
}