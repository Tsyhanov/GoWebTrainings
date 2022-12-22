package db

import (
	"fmt"
	"test-registration-form/config"
	"test-registration-form/models"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db *gorm.DB
var err error

//init db connection, migration
func Init() {
	fmt.Println("db:init")
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

	migrate()
}

func DbManager() *gorm.DB {
	return db
}

//Create tables if it does not exist
func migrate() {
	fmt.Println("db:migrate")

	err := db.AutoMigrate(&models.User{})
	if err != nil {
		panic("Could not create User table")
	}
	//create first and one user for test
	addTestUser()
}

//add test user if it does not exist
func addTestUser() {
	fmt.Println("db:addTestUser")
	var u *models.User
	db.Where("email = ?", "test@test.com").First(&u)
	if u.ID != 0 {
		fmt.Printf("email already exist")
	}

	u = config.LoadTestUser()
}
