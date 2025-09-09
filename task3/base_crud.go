package main

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Student struct {
	Id    uint `gorm:"primary_key"`
	Name  string
	Age   int
	Grade string
}

// 插入数据
func (student *Student) createStudent(db *gorm.DB) error {
	db.Debug().Create(student)
	return nil
}

func createStudents(db *gorm.DB, students *[]Student) error {
	db.Debug().Create(students)
	return nil
}

func main() {
	dsn := "root:12345677@tcp(127.0.0.1:3306)/grom?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	err = db.AutoMigrate(&Student{})
	if err != nil {
		panic(err)
	}

	//插入数据
	//student := Student{Name: "李磊", Age: 34, Grade: "五年级"}
	//err = student.createStudent(db)
	//if err != nil {
	//	return
	//}

	//

	//查询年龄大于18的所有学生
	//var students []Student
	//db.Debug().Where("Age >= ?", 18).Find(&students)
	//fmt.Println(students)
	//fmt.Println("完成")

	//更新张三为四年纪
	//db.Debug().Where("name = ?", "张三").Update("Grade", "四年级")
	//var student Student
	//db.Where("name = ?", "张三").Find(&student)
	//db.Debug().Model(&student).Updates(Student{Grade: "四年级"})

	//删除年龄小于17岁的学生记录
	db.Debug().Where("age < ?", 17).Delete(&Student{})

}
