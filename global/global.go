package global

import (
	"context"
	"sync"
)

type Status struct {
	CmdList        []map[string]string //name: 123 cmd :123
	CancelFuncList map[string][]*context.CancelFunc
}

//全局数据存储变量
var GlobalStatus Status

//全局锁
var GlobalLock sync.Mutex

const Success = "1"
const Stop = "2"
const Fail = "3"

//通过命令名称查找cmd某个item
func GetCmdListByName(name string) *map[string]string {

	for _, v := range GlobalStatus.CmdList {

		tempName := v["name"]

		if name == tempName {

			return &v
		}

	}

	return nil
}
