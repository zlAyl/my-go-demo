package lesson02

import (
	"fmt"
	"time"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name     string
	Age      int
	Birthday time.Time
}

// BeforeCreate hook
func (u *User) BeforeCreate(tx *gorm.DB) (err error) {

	return
}

func Run(db *gorm.DB) {
	err := db.AutoMigrate(&User{})
	if err != nil {
		panic(err)
	}
	//user := User{Name: "wang", Age: 21, Birthday: time.Now()}
	//db.Create(&user)

	//db.Select("Name", "Age", "CreatedAt").Create(&user)

	//过滤掉某列用Omit
	//db.Omit("Age").Create(&user)

	//批量插入
	//var users = []User{{Name: "jinzhu1"}, {Name: "jinzhu2"}, {Name: "jinzhu3"}}
	//db.Create(&users)

	//for _, user := range users {
	//	fmt.Println(user.ID)
	//	//user.ID // 1,2,3
	//}

	//批量插入批次大小
	//db.Omit("Birthday").CreateInBatches(users, 2)

	//跳过hook
	//db.Session(&gorm.Session{SkipHooks: true}).Create(&users)

	//处理冲突
	//冲突时不报错 也不插入
	//db.Clauses(clause.OnConflict{DoNothing: true}).CreateInBatches(users, 2)
	//db.Debug().Clauses(clause.OnConflict{
	//	Columns:   []clause.Column{{Name: "name"}},
	//	DoNothing: true, // 指定冲突检测列
	//}).Create(users)

	var user User
	db.Last(&user)

	fmt.Println(user.Age)
}
