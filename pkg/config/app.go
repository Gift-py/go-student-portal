/*
connecting to the database
return db instance
*/

package config

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var (
	db *gorm.DB
)

func Conn() *gorm.DB {
	db, err := gorm.Open("mysql", "root:AwesomeGod003#@!@/student-management-db?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic(err)
	}
	return db
}
