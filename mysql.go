package jotnar

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type Gorm struct {
	gormSet map[string]*gorm.DB
}

var gormInstance *Gorm

func initGorm() {
	gormInstance = &Gorm{make(map[string]*gorm.DB)}
	mainDb, err := gorm.Open("mysql", MysqlConfig["main"].Dsn)
	errExit(err)
	mainDb.DB().SetMaxIdleConns(MysqlConfig["main"].MaxIdle)
	mainDb.DB().SetMaxOpenConns(MysqlConfig["main"].MaxOpen)
	mainDb.LogMode(MysqlConfig["main"].EnableLog)

	gormInstance.gormSet["main"] = mainDb

	if MysqlConfig["salve"] != nil {
		slaveDb, err := gorm.Open("mysql", MysqlConfig["slave"].Dsn)
		errExit(err)

		slaveDb.DB().SetMaxIdleConns(MysqlConfig["slave"].MaxIdle)
		slaveDb.DB().SetMaxOpenConns(MysqlConfig["slave"].MaxOpen)
		mainDb.LogMode(MysqlConfig["slave"].EnableLog)

		gormInstance.gormSet["slave"] = slaveDb
	}
}

// use this to select
func ReadGorm() *gorm.DB {
	if sc, ok := gormInstance.gormSet["slave"]; ok {
		return sc
	}
	return gormInstance.gormSet["main"]
}

func WriteGorm() *gorm.DB {
	return gormInstance.gormSet["main"]
}
