package dao

import (
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var Db *gorm.DB

func UpdateDb(isOpen bool) {

	if isOpen {
		var err error
		Db, err = gorm.Open("mysql", "root:123456@tcp(127.0.0.1:3306)/gin")

		if err != nil {
			panic(err)

		} else {
			log.Println("数据库 mysql 连接成功")
		}
	} else {
		err := Db.Close()
		if err == nil {
			log.Println("数据库 mysql 连接关闭")
		} else {
			panic(err)
		}
	}
}

func GetDb() *gorm.DB {
	if Db == nil {
		var err error
		Db, err = gorm.Open("mysql", "root:123456@tcp(127.0.0.1:3306)/gin")

		if err != nil {
			panic(err)

		} else {
			log.Println("数据库 mysql 连接成功")
		}
	}
	return Db
}
