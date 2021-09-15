package core

import (
	orm "github.com/beego/beego/v2/client/orm"
	"sync"
)

var o orm.Ormer
var lock = sync.Mutex{}

func GetOrm() orm.Ormer {
	if o == nil {
		lock.Lock()
		o = orm.NewOrm()
		lock.Unlock()
	}
	return o
}
