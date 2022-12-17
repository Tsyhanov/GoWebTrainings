package db

import (
	"fmt"
	"test-registration-form/config"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db *gorm.DB
var err error

func Init() {
	fmt.Println("db init")
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		config.Config.DBUser,
		config.Config.DBPassword,
		config.Config.DBIp,
		config.Config.DBPort,
		config.Config.DBName)

	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	// defer db.Close()
	if err != nil {
		panic("DB Connection Error")
	}
	//	db.AutoMigrate(&model.User{})
}

func DbManager() *gorm.DB {
	return db
}
