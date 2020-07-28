package main

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"time"
)

type User7 struct {
	//gorm.Model
	Id        int
	Name      string
	Age       int64
	Active    bool
	DeletedAt *time.Time `gorm:"index:index_user_on_delete"`
}

func (u *User7) TableName() string {
	return "user_migrate"
}

func main() {
	db, err := gorm.Open("mysql", "root:12345678@(127.0.0.1:3306)/gorm_test?charset=utf8mb4&parseTime=True&loc=Local")
	if err != nil {
		panic(err)
	}
	defer db.Close()
	db.LogMode(true)

	// 迁移数据库.
	//db.AutoMigrate(&User7{})

	// 创建
	//user := User7{Name: "ccc", Age: 18, Active: false}
	//user2 := User7{Name: "ddd", Age: 28, Active: true}
	//db.Create(&user)
	//db.Create(&user2)

	// 删除
	db.Where("name=?", "ccc").Delete(User7{})

	var list []User7
	db.Find(&list)
}
