package main

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

// UserInfo 用户信息
type UserInfo struct {
	ID     uint
	Name   string
	Gender string
	Hobby  string
}

func main2() {
	db, err := gorm.Open("mysql", "root:12345678@(127.0.0.1:3306)/gorm_test?charset=utf8mb4&parseTime=True&loc=Local")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	// 自动迁移
	//db.AutoMigrate(&UserInfo{})

	//u1 := UserInfo{1, "七米", "男", "篮球"}
	//u2 := UserInfo{2, "沙河娜扎", "女", "足球"}
	//// 创建记录
	//db.Create(&u1)
	//db.Create(&u2)
	//// 查询
	var u = new(UserInfo)
	db.First(u)
	fmt.Printf("%#v\n", u)

	var u2 UserInfo
	db.First(&u2)
	fmt.Printf("%#v\n", u2)

	var uu UserInfo
	db.Find(&uu, "hobby=?", "双色球")
	fmt.Printf("%#v\n", uu)

	// 更新
	//db.Model(&uu).Update("hobby", "双色球")
	//db.Model(&uu).Updates(UserInfo{Name: "new name"})
	// 删除
	db.Delete(&uu)
}
