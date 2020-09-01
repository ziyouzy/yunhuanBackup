

/*对应了前端如温湿度，智能电表，ups某一个按钮*/
type OldSimpleYunHuanPod struct{
	Type string `json:"type"`  
	Nodes []*OldSimpleYunHuanNode
}