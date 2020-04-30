package main

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type User5 struct {
	gorm.Model
	Name   string
	Age    int64
	Active bool
}

func (u *User5) TableName() string {
	return "user_update"
}

func main5() {
	db, err := gorm.Open("mysql", "root:12345678@(127.0.0.1:3306)/gorm_test?charset=utf8mb4&parseTime=True&loc=Local")
	if err != nil {
		panic(err)
	}
	defer db.Close()
	db.LogMode(true)

	// 迁移数据库.
	//db.AutoMigrate(&User5{})

	// 创建
	//user := User5{Name: "aaa", Age: 18, Active: false}
	//user2 := User5{Name: "bbb", Age: 28, Active: true}
	//db.Create(&user)
	//db.Create(&user2)

	var user3 User5
	db.First(&user3)

	user3.Name = "hjx"
	user3.Age = 100
	db.Save(&user3)
	// 只有第一个值"name"会被更新，"age"没更新
	db.Model(&user3).Update("name", "hjx", "age", 10)
	// 会全部更新"name"，"age"
	db.Model(&user3).Updates(User5{Name: "hjx", Age: 10})
	// 只会更新非零值
	db.Model(&user3).Updates(User5{Name: "", Age: 0, Active: false})

	m1 := map[string]interface{}{
		"name":   "hello",
		"age":    0,
		"active": false,
	}
	// map零值也会更新
	db.Model(&user3).Updates(m1)
	// 只更新某个列
	db.Model(&user3).Select("name").Updates(m1)
	// 忽略某些列的更新
	db.Model(&user3).Omit("age", "name").Updates(m1)

	// 更新指定列，不更新钩子
	db.Model(&user3).UpdateColumn("name", "hello")

	// 不执行钩子
	db.Table("user_update").Where("id IN (?)", []int{10, 11}).Updates(map[string]interface{}{"name": "hello", "age": 18})
	// 执行钩子
	db.Model(User5{}).Updates(User5{Name: "hello", Age: 18})
	// 执行钩子
	rawAffected := db.Model(User5{}).Updates(User5{Name: "hello", Age: 18}).RowsAffected
	fmt.Println("rawAffected: ", rawAffected)

	// sql表达式
	db.Model(User5{}).Update("age", gorm.Expr("age + ?", 2))

}
