package lesson01

import (
	"database/sql"
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID           uint           // Standard field for the primary key
	Name         string         // A regular string field
	Email        *string        // A pointer to a string, allowing for null values
	Age          uint8          // An unsigned 8-bit integer
	Birthday     *time.Time     // A pointer to time.Time, can be null
	MemberNumber sql.NullString // Uses sql.NullString to handle nullable strings
	ActivatedAt  sql.NullTime   // Uses sql.NullTime for nullable time fields
	CreatedAt    time.Time      // Automatically managed by GORM for creation time
	UpdatedAt    time.Time      // Automatically managed by GORM for update time
	ignored      string         // fields that aren't exported are ignored
}

type Author struct {
	Name  string
	Email string
}
type Blog2 struct {
	ID     int64
	Author Author `gorm:"embedded;embeddedPrefix:author_"` // embedded 声明是嵌入模式 并且增加前缀
	// Author  Author
	Upvotes int32
}

func Run(db *gorm.DB) {
	//err := db.AutoMigrate(&User{}) //自动创建表
	//if err != nil {
	//	panic(err)
	//}
	//user := &User{}
	//user.MemberNumber.Valid = true //开启这个可以让默认null变成空字符串
	//db.Create(user)                //创建一条记录

}
