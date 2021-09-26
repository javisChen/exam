package db

import (
	"database/sql"
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

func SelectOne(sql string, container interface{}, args ...interface{}) error {
	return GetOrm().Raw(sql, args).QueryRow(container)
}

func SelectList(sql string, container interface{}, args ...interface{}) (int64, error) {
	rows, err := funcName(sql, args).QueryRows(container)
	return rows, err
}

func funcName(sql string, args []interface{}) orm.RawSeter {
	return GetOrm().Raw(sql, args)
}

func Exec(sql string, args ...interface{}) (sql.Result, error) {
	result, err := GetOrm().Raw(sql, args).Exec()
	return result, err
}

func ExecTx(txOrm orm.TxOrmer, sql string, args ...interface{}) (sql.Result, error) {
	result, err := txOrm.Raw(sql, args).Exec()
	return result, err
}
