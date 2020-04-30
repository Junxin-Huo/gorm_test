package main

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type User6 struct {
	gorm.Model
	Name   string
	Age    int64
	Active bool
}

func (u *User6) TableName() string {
	return "user_delete"
}

func main() {
	db, err := gorm.Open("mysql", "root:12345678@(127.0.0.1:3306)/gorm_test?charset=utf8mb4&parseTime=True&loc=Local")
	if err != nil {
		panic(err)
	}
	defer db.Close()
	db.LogMode(true)

	// 迁移数据库.
	//db.AutoMigrate(&User6{})

	// 创建
	//user := User6{Name: "ccc", Age: 18, Active: false}
	//user2 := User6{Name: "ddd", Age: 28, Active: true}
	//db.Create(&user)
	//db.Create(&user2)

	// 删除
	//var user3 User6
	//// 删除执行id
	////user3.ID = 1
	//// 全删除（从删库到跑路）
	//user3.Name = "ccc"
	//db.Delete(user3)

	// 正确的删除
	db.Where("name=?", "ccc").Delete(User6{})
	db.Delete(User6{}, "name=?", "ccc")

	// 有deleted_at，默认软删除
	// 真正删除，或者要查找软删除了的数据，用Unscoped()
	var user4 User6
	fmt.Println(db.Where("name = ?", "ccc").Find(&user4).RowsAffected)
	fmt.Printf("%v\n", user4)
	fmt.Println(db.Unscoped().Where("name = ?", "ccc").Find(&user4).RowsAffected)
	fmt.Printf("%v\n", user4)

	// 物理删除
	db.Unscoped().Where("name = ?", "bbb").Delete(User6{})
}
