package db

import (
	"database/sql"
	"fmt"
	orm "github.com/beego/beego/v2/client/orm"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"sync"
)

var o orm.Ormer
var lock = sync.Mutex{}

var db *gorm.DB

func init() {
	//dbUrl, _ := beego.AppConfig.String("datasource.url")
	db, _ = gorm.Open(mysql.New(mysql.Config{
		DSN:                       "root:root@/exam?charset=utf8&parseTime=true", // data source name
		DefaultStringSize:         256,                                           // default size for string fields
		DisableDatetimePrecision:  true,                                          // disable datetime precision, which not supported before MySQL 5.6
		DontSupportRenameIndex:    true,                                          // drop & create when rename index, rename index not supported before MySQL 5.7, MariaDB
		DontSupportRenameColumn:   true,                                          // `change` when rename column, rename column not supported before MySQL 8, MariaDB
		SkipInitializeWithVersion: false,                                         // auto configure based on currently MySQL version
	}), &gorm.Config{})
}

//func db orm.Ormer {
//	if o == nil {
//		lock.Lock()
//		o = orm.NewOrm()
//		lock.Unlock()
//	}
//	return o
//}

func SelectOne(sql string, container interface{}, args ...interface{}) error {
	result := db.Raw(sql, args).Scan(container)
	return result.Error
}

func SelectList(sql string, container interface{}, args ...interface{}) (int64, error) {
	result := db.Raw(sql, args).Scan(container)
	return result.RowsAffected, result.Error
}

func Exec(sql string, args ...interface{}) (sql.Result, error) {
	err := db.Exec(sql, args...)
	//exec := db.Exec("UPDATE orders SET shipped_at = ? WHERE id IN ?", args)
	fmt.Println(err)
	return nil, nil
}

//func ExecTx(txOrm orm.TxOrmer, sql string, args ...interface{}) (sql.Result, error) {
//	result, err := txOrm.Exec(sql, args)
//	return result, err
//}
