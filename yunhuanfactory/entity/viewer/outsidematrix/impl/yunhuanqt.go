package impl


/*对应了某一组机柜矩阵*/
type OldYunHuanQtViewerMatrix struct {
	Type string `json:"type"` 
	ServerCabinets []*OldSimpleYunHuanServerCabinet
}





