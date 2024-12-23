package model

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"time"
)

// 是gorm的数据库连接对象
var db *gorm.DB

// Model 实现gorm的基本字段的要求
type Model struct {
	ID        uint64          `gorm:"primary_key" json:"id"`
	CreatedAt time.Time       `json:"created_at"`
	UpdatedAt time.Time       `json:"updated_at"`
	DeletedAt *gorm.DeletedAt `gorm:"index" json:"deleted_at"`
}

/*
	异常和错误是两回事
*/

func InitMysql() {
	dsn := "root:ALin0915=@tcp(127.0.0.1:3306)/polaris_backend_go?charset=utf8mb4&parseTime=True&loc=Local"
	var err error
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		// 如果这边有错误，后续的服务还要去执行么？
		panic(err)
	}
}

// UseDB 拿到model中的db
func UseDB() *gorm.DB {
	return db
}
