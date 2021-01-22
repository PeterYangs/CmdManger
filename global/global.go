package global

import "sync"

type Status struct {
	CmdList []map[string]string //name: 123 cmd :123

}

//全局数据存储变量
var GlobalStatus Status

//全局锁
var GlobalLock sync.Mutex
