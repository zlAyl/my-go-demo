package main

import (
	"errors"
	"fmt"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name      string
	Age       int
	Gender    string
	Birth     time.Time `gorm:"type:date"`
	PostCount int       `gorm:"default:0"`
	Posts     []Post    `gorm:"foreignKey:UserID"`
	Comments  []Comment `gorm:"foreignKey:UserID"`
}

type Post struct {
	gorm.Model
	Title        string
	Author       string
	UserID       uint
	CommentCount uint      `gorm:"default:0"`
	CommentState uint      `gorm:"type:tinyint(1);default:0;comment:评论状态:1已评论0无评论"`
	Comments     []Comment `gorm:"foreignKey:PostID"`
}

func (post *Post) AfterSave(tx *gorm.DB) (err error) {
	var user User
	if err := tx.Find(&user, "id = ?", post.UserID).Error; err != nil {
		return err
	}
	user.PostCount += 1
	if err := tx.Save(&user).Error; err != nil {
		return err
	}
	return nil
}

type Comment struct {
	gorm.Model
	Message string
	PostID  uint
	UserID  uint
}

func (comment *Comment) AfterSave(tx *gorm.DB) (err error) {
	var post Post
	if err := tx.Where("id = ?", comment.PostID).Find(&post).Error; err != nil {
		return errors.New("没有找到评论对应的文章")
	}
	post.CommentCount += 1
	post.CommentState = 1
	if err := tx.Save(&post).Error; err != nil {
		return errors.New("更新文章评论数失败")
	}

	return nil
}

func (comment *Comment) AfterDelete(tx *gorm.DB) (err error) {
	fmt.Println("删除的数据", comment.DeletedAt.Valid)
	if comment.DeletedAt.Valid { //删除的时候
		var post Post
		if err := tx.Find(&post, "id = ?", comment.PostID).Error; err != nil {
			return err
		}
		post.CommentCount -= 1
		if post.CommentCount <= 0 {
			post.CommentState = 0
			post.CommentCount = 0
		}
		if err := tx.Save(&post).Error; err != nil {
			return err
		}
	}
	return nil
}

func main() {
	dsn := "root:12345677@tcp(127.0.0.1:3306)/grom?charset=utf8mb4&parseTime=True"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	db.AutoMigrate(&User{}, &Post{}, &Comment{})
	fmt.Println("完成")

	//插入数据
	//users := []User{
	//	{Name: "Alice", Age: 30, Gender: "male", Birth: time.Now()},
	//	{Name: "Bob", Age: 22, Gender: "female", Birth: time.Now()},
	//	{Name: "张三", Age: 20, Gender: "male", Birth: time.Now()},
	//	{Name: "李四", Age: 31, Gender: "female", Birth: time.Now()},
	//}
	//db.Debug().Create(&users)
	//fmt.Println("用户表的数据初始化完成")

	//posts := []Post{
	//	{Title: "我是一个小小鸟", Author: "大力", UserID: 1},
	//	{Title: "我是一个小小鸟2", Author: "大力", UserID: 1},
	//	{Title: "我是一个大大的树", Author: "小力", UserID: 2},
	//	{Title: "我是一个孤独的船", Author: "小力力", UserID: 3},
	//	{Title: "我是一个孤独的船2", Author: "小力力", UserID: 4},
	//}
	//db.Create(&posts)
	//fmt.Println("文章表的数据初始化完成")

	//comments := []Comment{
	//	{Message: "非常棒", PostID: 1, UserID: 3},
	//	{Message: "非常棒棒", PostID: 1, UserID: 4},
	//	{Message: "非常棒棒棒", PostID: 1, UserID: 2},
	//	{Message: "非常好", PostID: 2, UserID: 2},
	//	{Message: "非常好哦", PostID: 3, UserID: 2},
	//	{Message: "非常棒哦", PostID: 4, UserID: 1},
	//	{Message: "非常棒哦哦", PostID: 4, UserID: 2},
	//}
	//db.Debug().Create(&comments)
	//fmt.Println("评论表的数据初始化完成")

	//查询用户id=1发表的所有文章和以及对应的评论
	//var posts []Post
	//db.Debug().Preload("Comments").Where("user_id = ?", 1).Find(&posts)
	//for _, post := range posts {
	//	fmt.Printf("用户 %d 发表的文章 %s\n", post.UserID, post.Title)
	//	for _, comment := range post.Comments {
	//		fmt.Printf("文章 %s 的评论有 %s\n", post.Title, comment.Message)
	//	}
	//}
	//fmt.Println(posts)

	//查询评论数最多的文章信息
	//var post Post
	//db.Debug().Order("comment_count desc").First(&post)
	//
	//fmt.Println("评论数最多的文章信息:", post)

	//删除评论
	var comments []Comment
	db.Where("post_id = ?", 1).Find(&comments)
	db.Delete(&comments)
	fmt.Println("删除完成")

}
