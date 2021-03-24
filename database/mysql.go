package database

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

var Eloquent *gorm.DB

func init() {
	var err error
	Eloquent, err = gorm.Open("mysql", "root:password@tcp(ip:3306)/gin_api?charset=utf8&parseTime=True&loc=Local")

	if err != nil {
		fmt.Printf("mysql connect error %V", err)
	}

	fmt.Printf("mysql connect success!")

	Eloquent.SingularTable(true)
	Eloquent.DB().SetMaxIdleConns(10)
	Eloquent.DB().SetMaxOpenConns(100)
}

func CloseDB() {
	defer Eloquent.Close()
}
