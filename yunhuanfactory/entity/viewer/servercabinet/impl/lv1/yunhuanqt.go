
/*对应了某一台机柜*/
type OldSimpleYunHuanServerCabinet struct{
	Type string `json:"type"` 
	Pods []*OldSimpleYunHuanPod
}