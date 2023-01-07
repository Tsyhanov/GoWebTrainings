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

func DbManager() *gorm.DB {
	return db
}

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

//Create tables if it does not exist
func migrate() {
	fmt.Println("db:migrate")

	err := db.AutoMigrate(&models.User{}, &models.Post{}, &models.Comment{})
	if err != nil {
		panic("db.migrate:Could not create table")
	}
	//create user for test if it doesn't exist
	addTestUser()
}

//add test user if it does not exist
func addTestUser() {
	fmt.Println("db:addTestUser")
	var u *models.User
	db.Where("email = ?", "test@test.com").First(&u)
	if u.ID != 0 {
		fmt.Printf("email already exist")
	} else {
		u = config.LoadTestUser()

		result := db.Create(u)
		if result.Error != nil {
			panic("Could not create test user")
		}
	}
}

//check if email exist in db
func CheckIfEmailExist(email string) bool {
	var u *models.User
	err := db.Where("email = ?", email).First(&u).Error
	return err != gorm.ErrRecordNotFound
}

//Create new user
func CreateUser(u *models.User) bool {
	fmt.Println("db:CreateUser")
	result := db.Create(&u)
	return result.Error == nil
}

//get user from db
func GetUser(u *models.User, email string) error {
	fmt.Println("db:GetUser")
	err := db.First(&u, "email =?", email).Error
	return err
}

//get all comments
func GetComments() ([]models.Comment, error) {
	var cmt []models.Comment

	fmt.Println("getComments")
	result := db.Find(&cmt)

	if result.Error != nil {
		fmt.Println("getComments error:", result.Error)
		return cmt, result.Error
	}
	return cmt, nil
}

//get all posts
func GetPosts() ([]models.Post, error) {
	fmt.Println("getPosts")
	var p []models.Post
	result := db.Find(&p)
	if result.Error != nil {
		fmt.Println("select from posts error")
		return p, result.Error
	}
	return p, nil
}

//get post by id
func GetPostById(id int) (models.Post, error) {
	fmt.Println("getPostById")
	var p models.Post
	result := db.First(&p, id)

	if result.Error != nil {
		return p, result.Error
	}
	return p, nil
}

//add post
func AddPost(p models.Post) error {
	result := db.Create(&p)
	if result.Error != nil {
		fmt.Println("insert into posts error")
	}

	if result.Error != nil {
		fmt.Println("AddPost error:", result.Error)
		return result.Error
	}
	return nil
}

//add comment
func AddComment(c models.Comment) error {
	result := db.Create(&c)
	if result.Error != nil {
		fmt.Println("AddPost error:", result.Error)
		return result.Error
	}
	return nil
}
