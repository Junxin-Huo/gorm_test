package main

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type User4 struct {
	gorm.Model
	Name string
	Age  int64
}

func (u *User4) TableName() string {
	return "user_find"
}

func AgeGreaterThan20(db *gorm.DB) *gorm.DB {
	return db.Where("age > ?", 20)
}

func NameIsNotPig(db *gorm.DB) *gorm.DB {
	return db.Where("name <> ?", "pig")
}

func main4() {
	db, err := gorm.Open("mysql", "root:12345678@(127.0.0.1:3306)/gorm_test?charset=utf8mb4&parseTime=True&loc=Local")
	if err != nil {
		panic(err)
	}
	defer db.Close()
	db.LogMode(true)

	// 迁移数据库
	//db.AutoMigrate(&User4{})

	// 创建
	//user := User4{Name: "aaa", Age: 18}
	//user2 := User4{Name: "bbb", Age: 28}
	//db.Debug().Create(&user)
	//db.Debug().Create(&user2)

	// 查询
	//var user4 User4
	user4 := new(User4)
	db.First(user4)
	fmt.Printf("user: %#v\n", user4)

	users := make([]User4, 2, 4)
	//fmt.Printf("users: %#v\n", users)
	db.Find(&users)
	fmt.Printf("users: %#v\n", users)

	ints := make(chan int, 3)
	m := make(map[string]int)
	i := make([]int, 5)
	ints <- 1
	m["aa"] = 11
	i = append(i, 10)
	fmt.Printf("ints: %#v\nm: %#v\ni: %#v\n", ints, m, i)
	
	var user5 User4
	// FirstOrInit 不会创建
	// Attrs 未找到赋值
	// Assign 找到未找到都赋值
	//db.Attrs(User4{Age: 99}).FirstOrInit(&user5, User4{Name: "aaa"})
	db.Assign(User4{Age: 99}).FirstOrInit(&user5, User4{Name: "aaa"})
	// FirstOrCreate 会创建
	//db.FirstOrCreate(&user5, User4{Name: "hjx"})
	fmt.Printf("user5: %#v\n", user5)

	var user6 User4
	db.Scopes(AgeGreaterThan20, NameIsNotPig).Find(&user6)
	fmt.Printf("user6: %#v\n", user6)
}
