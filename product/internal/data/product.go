package data

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	Id     int32  `gorm:"primaryKey"`
	Name   string `gorm:"type:varchar(255);not null;default:''"`
	Desc   string `gorm:"type:varchar(255);not null;default:''"`
	Stock  int32  `gorm:"not null;default:0"`
	Amount int32  `gorm:"not null;default:0"`
	Status int32  `gorm:"not null;default:0"`
}

var DbConn *gorm.DB

func Connect() {
	var err error
	dsn := "root:example@tcp(127.0.0.1:33306)/xshop?charset=utf8mb4&parseTime=true"
	DbConn, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	DbConn.AutoMigrate(&Product{})
}
