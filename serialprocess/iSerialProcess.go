package serialprocess

import (
	_"github.com/ziyouzy/mylib/serialprocess/impl"
)
type ISerialProcess interface{
	readThread()
	sendThread()
	Start()
}

