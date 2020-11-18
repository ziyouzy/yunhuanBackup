package service

import(
	//"bytes"
	//"fmt"

	"github.com/ziyouzy/mylib/connserver"
)


func RawCh()chan []byte{return connserver.LoadSingletonPatternRecvCh()}
func ConnServerListenAndCollect(){connserver.LoadSingletonPatternListenAndCollect()}
