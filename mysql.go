package jotnar

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type Gorm struct {
	gormSet map[string]*gorm.DB
}

var gormInstance *Gorm

func InitGorm() {
	gormInstance = &Gorm{make(map[string]*gorm.DB)}
	mainDb, err := gorm.Open("mysql", MConfig.MysqlSet["main"].Dsn)
	errExit(err)

	gormInstance.gormSet["main"] = mainDb

	if MConfig.MysqlSet["salve"] != nil {
		slaveDb, err := gorm.Open("mysql", MConfig.MysqlSet["slave"].Dsn)
		errExit(err)

		gormInstance.gormSet["slave"] = slaveDb
	}
}

// use this to select
func ReadGorm() *gorm.DB {
	if gormInstance.gormSet["slave"] != nil {
		return gormInstance.gormSet["slave"]
	}
	return gormInstance.gormSet["main"]
}

func WriteGorm() *gorm.DB {
	return gormInstance.gormSet["main"]
}
