package main

import (
	"errors"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

type Book struct {
	Id     int     `db:"id,primary_key,autoincrement"`
	Title  string  `db:"title"`
	Author string  `db:"author"`
	Price  float64 `db:"price"`
}

func main() {
	dsn := "root:12345677@tcp(127.0.0.1:3306)/grom?charset=utf8mb4&parseTime=True"
	db, err := sqlx.Connect("mysql", dsn)
	if err != nil {
		panic(err)
	}
	//if err := createTableBooks(db); err != nil {
	//	panic(err)
	//}
	//fmt.Println("创建 books 表成功")

	//books := []Book{
	//	{Title: "Go语言编程", Author: "许式伟", Price: 65.00},
	//	{Title: "Go语言实战", Author: "William Kennedy", Price: 79.00},
	//	{Title: "Go Web编程", Author: "Sau Sheong Chang", Price: 59.00},
	//	{Title: "Go语言学习笔记", Author: "雨痕", Price: 45.00},
	//}
	//if err := batchInsertBooks(db, books); err != nil {
	//	panic(err)
	//}
	//fmt.Println("插入数据成功")

	books, err := getBooksByMinPrice(db, 50)
	if err != nil {
		panic(err)
	}
	fmt.Println("价格大于50元的书籍有:", books)

}

func createTableBooks(db *sqlx.DB) error {
	createTableSQL := `
    CREATE TABLE IF NOT EXISTS books (
        id INT AUTO_INCREMENT PRIMARY KEY,
        title VARCHAR(100) NOT NULL,
        author VARCHAR(50) NOT NULL,
        price DECIMAL(10, 2) NOT NULL
    ) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;
    `
	_, err := db.Exec(createTableSQL)
	return err
}

func batchInsertBooks(db *sqlx.DB, books []Book) error {
	tx, err := db.Beginx()
	if err != nil {
		return err
	}
	for _, book := range books {
		_, err := tx.NamedExec("insert into books (title, author, price) values(:title, :author, :price)", book)
		if err != nil {
			err := tx.Rollback()
			if err != nil {
				return errors.New("事务回滚失败")
			}
			return err
		}
	}
	if err := tx.Commit(); err != nil {
		return errors.New("事务提交失败")
	}
	return nil
}

func getBooksByMinPrice(db *sqlx.DB, minPrice float64) ([]Book, error) {
	var books []Book
	err := db.Select(&books, "select * from books where price >= ?", minPrice)
	if err != nil {
		return nil, err
	}
	return books, nil

}
