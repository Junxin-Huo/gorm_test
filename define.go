package main

import (
	"database/sql"
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"time"
)

type User struct {
	gorm.Model
	Name         string
	Age          sql.NullInt64
	Birthday     *time.Time
	Email        string  `gorm:"type:varchar(100);unique_index"`
	Role         string  `gorm:"size:255"`        // 设置字段大小为255
	MemberNumber *string `gorm:"unique;not null"` // 设置会员号（member number）唯一并且不为空
	Num          int     `gorm:"AUTO_INCREMENT"`  // 设置 num 为自增类型
	Address      string  `gorm:"index:addr"`      // 给address字段创建名为addr的索引
	IgnoreMe     int     `gorm:"-"`               // 忽略本字段
}

type User2 struct {
	ID        uint      // column name is `id`
	Name      string    // column name is `name`
	Birthday  time.Time // column name is `birthday`
	CreatedAt time.Time // column name is `created_at`
}

func (u *User2) TableName() string {
	return "user_table"
}

type Animal struct {
	AnimalId int64     `gorm:"column:beast_id"`         // set column name to `beast_id`
	Birthday time.Time `gorm:"column:day_of_the_beast"` // set column name to `day_of_the_beast`
	Age      int64     `gorm:"column:age_of_the_beast"` // set column name to `age_of_the_beast`
}

func main2() {
	db, err := gorm.Open("mysql", "root:12345678@(127.0.0.1:3306)/gorm_test?charset=utf8mb4&parseTime=True&loc=Local")
	if err != nil {
		panic(err)
	}
	defer db.Close()
	db.LogMode(true)

	// 创建表、更新表结构
	db.AutoMigrate(&User{}, &User2{}, &Animal{})

	// 插入
	//t, _ := time.ParseInLocation("2006-01-02 15:04:05", "2019-06-18 08:47:50", time.Local)
	//num := "887799669"
	//u1 := User{
	//	Name:         "aaa",
	//	Age:          sql.NullInt64{25, true},
	//	Birthday:     &t,
	//	Email:        "1234@qq.com",
	//	Role:         "Student",
	//	MemberNumber: &num,
	//	Address:      "qqwwaass",
	//	IgnoreMe:     666,
	//}
	//db.Create(&u1)

	// 查询
	var u2 []User
	db.Find(&u2, "role=?", "Student")
	fmt.Printf("%#v\n", u2)

	var u3 User
	db.Where("email=?", "123@qq.com").First(&u3)
	fmt.Printf("%#v\n", u3)

	//u3.Age = sql.NullInt64{26, true}
	//// model为实例，更新该条记录
	//db.Model(u3).Updates(User{
	//	Name: "Steven",
	//})

	// model为类型，更新全表
	//db.Model(User{}).Updates(User{
	//	Name: "Bob",
	//})

	// 删除（delete_at字段设置当前时间，并未删除，软删除）
	db.Delete(u3)
}
