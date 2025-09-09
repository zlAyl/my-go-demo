package main

import (
	"errors"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Account struct {
	ID      uint `gorm:"primary_key"`
	Balance float64
}

type Transaction struct {
	ID            uint `gorm:"primary_key"`
	FromAccountId uint
	ToAccountId   uint
	Amount        float64
}

func main() {
	dsn := "root:12345677@tcp(127.0.0.1:3306)/grom?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	//db.AutoMigrate(&Account{})
	//db.AutoMigrate(&Transaction{})

	//accounts := []Account{
	//	{Balance: 200},
	//	{Balance: 300},
	//}
	//db.Create(&accounts)

	err = db.Transaction(func(tx *gorm.DB) error {
		aAccount := Account{}
		db.Where("id", 1).Find(&aAccount)

		bAccount := Account{}
		db.Where("id", 2).Find(&bAccount)
		amount := float64(100)
		if aAccount.Balance >= amount {
			aAccount.Balance -= amount
			bAccount.Balance += amount
			if err := db.Debug().Save(&aAccount).Error; err != nil {
				return err
			}
			if err := db.Debug().Save(&bAccount).Error; err != nil {
				return err
			}
			transaction := &Transaction{FromAccountId: aAccount.ID, ToAccountId: bAccount.ID, Amount: amount}
			if err := db.Debug().Create(&transaction).Error; err != nil {
				return err
			}
			return nil
		}
		return errors.New("账户余额不足")
	})
	if err != nil {
		println(err.Error())
	}
	println("转账完成")
}
