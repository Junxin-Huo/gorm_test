package main

import (
	"database/sql"
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type User3 struct {
	gorm.Model
	//Name string `gorm:"default:'小王子'"`
	//Name *string `gorm:"default:'小王子'"`
	Name sql.NullString `gorm:"default:'小王子'"`
	Age  int64
}

func (u *User3) TableName() string {
	return "user_create"
}

func main() {
	db, err := gorm.Open("mysql", "root:12345678@(127.0.0.1:3306)/gorm_test?charset=utf8mb4&parseTime=True&loc=Local")
	if err != nil {
		panic(err)
	}
	defer db.Close()
	//db.LogMode(true)

	// 迁移数据库
	//db.AutoMigrate(&User3{})

	//user := User3{Name: "q1mi", Age: 18}

	// 插入空
	//user := User3{Name: "", Age: 18}
	//user := User3{Name: new(string), Age: 28}
	user := User3{Name: sql.NullString{
		String: "",
		Valid:  true,
	}, Age: 18}

	fmt.Println(db.NewRecord(user)) // 主键为空返回`true`
	db.Debug().Create(&user)                // 创建user
	fmt.Println(db.NewRecord(user)) // 创建`user`后返回`false`
}
