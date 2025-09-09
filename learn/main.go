package main

import (
	"fmt"

	"github.com/zlAyl/my-go-demo/learn/lesson02"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	dsn := "root:12345677@tcp(127.0.0.1:3306)/grom?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	//lesson01.Run(db)
	
	lesson02.Run(db)

	fmt.Println("完成")

}
