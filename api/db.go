package api

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"log"
)

var db *gorm.DB
func InitDb() *gorm.DB {
	db, err := gorm.Open("mysql", "root:gpstracker@/gps_tracker?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
        log.Panic(err)
    }
	return db
}