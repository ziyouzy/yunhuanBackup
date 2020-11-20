package service

import(
	//"bytes"
	//"fmt"

	"github.com/ziyouzy/mylib/connserver"
)


//可拿到未设定消费者也未指定生产者的RawCh，只要确保其在消费的时候拥有独立的子携程即可
func RawCh()chan []byte{return connserver.LoadSingletonPatternRecvCh()}

//实现上面RawCh的生产者
func ConnServerListenAndCollect(){connserver.LoadSingletonPatternListenAndCollect()}
