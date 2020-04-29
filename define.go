package main

import (
	"database/sql"
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

func main() {
	db, err := gorm.Open("mysql", "root:12345678@(127.0.0.1:3306)/gorm_test?charset=utf8mb4&parseTime=True&loc=Local")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	db.AutoMigrate(&User{}, &User2{}, &Animal{})
}
