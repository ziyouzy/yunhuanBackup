package service

import(
	"fmt"

	"github.com/ziyouzy/mylib/nodedo"
)

func SendNodeDoBytesToSouthBound(nd nodedo.NodeDo){
	fmt.Println("c.准备通过connserver.ClientMap()['127.0.0.1']发送,nd.GetJson():"/*,nd.GetJson()*/)
	//connserver.ClientMap()["127.0.0.1"].SendBytes(nd.GetJson())
}