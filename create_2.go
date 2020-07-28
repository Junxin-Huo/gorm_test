package main

import (
	_ "database/sql"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type User8 struct {
	gorm.Model
	//Name string `gorm:"default:'小王子'"`
	Name *string `gorm:"default:'小王子'"`
	//Name sql.NullString `gorm:"default:'小王子'"`
	//Age  int64 `gorm:"default:'20'"`
	Age   *int64 `gorm:"default:'20'"`
	IsX64 *int32
}

func (u *User8) TableName() string {
	return "user_create_2"
}

func main() {
	db, err := gorm.Open("mysql", "root:12345678@(127.0.0.1:3306)/gorm_test?charset=utf8mb4&parseTime=True&loc=Local")
	if err != nil {
		panic(err)
	}
	defer db.Close()
	//db.LogMode(true)

	// 迁移数据库
	db.AutoMigrate(&User8{})

	// 插入空
	//user := User8{Name: "", Age: 0}
	//user := User8{Name: new(string), Age: new(int64)}

	//fmt.Println(db.NewRecord(user)) // 主键为空返回`true`
	//db.Debug().Create(&user)                // 创建user
	//fmt.Println(db.NewRecord(user)) // 创建`user`后返回`false`

	//db.Debug().Model(User8{}).Where("id = ?", "3").Updates(User8{Name: "", Age: 0})
	//name := ""
	//var age int64
	//age = 0
	//db.Debug().Model(User8{}).Where("id = ?", "3").Updates(User8{Name: &name, Age: &age})

	db.Debug().Where("name IN (?)", []string{"b", "a", "z", "y"}).Delete(User8{})
}
