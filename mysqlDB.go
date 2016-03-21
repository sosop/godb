package godb

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
	"github.com/sosop/libconfig"
)

var MysqlDB *sql.DB

func init() {
	iniConf := libconfig.NewIniConfig("db.conf")
	driver := iniConf.GetString("mysql::driver")
	dataSource := iniConf.GetString("mysql::dataSource")
	var err error
	MysqlDB, err = sql.Open(driver, dataSource)
	if err != nil {
		panic(err)
	}
	// defer MysqlDB.Close()
	err = MysqlDB.Ping()
	if err != nil {
		MysqlDB.Close()
		panic(err)
	}
}
