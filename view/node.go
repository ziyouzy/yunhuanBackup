//node来自于conf.ConfNode接口，conf包内设计了基于配置文档的对先关系映射，从而生成组成view的最小单位node
//而最终真正会拼装成什么样的view，是由protocol包来决定
//view包应像是使用mysql-orm一样使用conf包，从而获取最小单位node
package view

import (
	_"github.com/ziyouzy/mylib/conf"
)

type Node []byte
//一般不会用到这个，而是直接用从conf包返回的[]byte直接作为最小零件拼接，这里只是个形式