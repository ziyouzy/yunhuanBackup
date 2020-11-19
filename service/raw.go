package service

import(
	//"bytes"
	//"fmt"

	"github.com/ziyouzy/mylib/connserver"
)


//拿到未设定消费者的RawCh，只要确保其在消费的时候拥有独立的子携程即可
func RawCh()chan []byte{return connserver.LoadSingletonPatternRecvCh()}

func ConnServerListenAndCollect(){connserver.LoadSingletonPatternListenAndCollect()}
