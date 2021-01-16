package service

import(
	"github.com/ziyouzy/mylib/physicalnode"
	"github.com/ziyouzy/mylib/connserver"
)

func BuildPNCh(){
	rawCh :=connserver.RawCh()
	physicalnode.ToPhysicalNodeCh(rawCh, nil)
}

RawParser