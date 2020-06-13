package config

import (
	"github.com/jinzhu/gorm"
	_ "github.com/go-sql-driver/mysql"
)

func GetDB() (db *gorm.DB, err error)	{
	dbDriver:="mysql"
	dbUser:="root"
	dbPass:=""
	dbName:="event_db"
	db, err=gorm.Open(dbDriver,dbUser+":"+dbPass+"@(localhost)/"+dbName+"?charset=utf8&parseTime=True&loc=Local")	
	return
}