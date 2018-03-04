package models

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	//"time"
)

var (
	mysqlname   = "root"
	mysqlpwd    = "xchus2a9"
	mysqlhost   = "127.0.0.1"
	mysqlport   = 3306
	mysqldbname = "excelgo"
	maxIdle     = 50
	maxOpen     = 100
)

var db *gorm.DB

func init() {
	db = CreateDB(mysqlname, mysqlpwd, mysqlhost, mysqlport, mysqldbname, maxIdle, maxOpen)
	db.AutoMigrate(&WscnWork{})
}

func Clean() {
	db.Close()
}

func CreateDB(mysqlname, mysqlpwd, mysqlhost string, mysqlport int, mysqldbname string, maxIdle, maxOpen int) *gorm.DB {
	connStr := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local",
		mysqlname, mysqlpwd, mysqlhost, mysqlport, mysqldbname)
	database, err := gorm.Open("mysql", connStr)
	if err != nil {
		return database
	}

	database.DB().SetMaxIdleConns(maxIdle)
	database.DB().SetMaxOpenConns(maxOpen)

	return database
}

func Session() *gorm.DB {
	return db.Begin()
}

func DB() *gorm.DB {
	return db
}

//type TimeRec struct {
//	CreateAt		time.Time		`gorm:"columnn:created_at;type:TIMESTAMP;default:CURRENT_TIMESTAMP"`
//	UpdateAt		time.Time		`gorm:"columnn:created_at;type:TIMESTAMP;default:CURRENT_TIMESTAMP"`
//}
