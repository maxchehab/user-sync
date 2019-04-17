package main

import (
	"log"
	"models"
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

var database *gorm.DB

// IntializeDatabase creates tables for all of the models
func IntializeDatabase() (err error) {
	// set up DB connection
	// then attempt to connect 10 times over 10 seconds
	//connectionParams := "user=docker password=docker sslmode=disable host=db"
	//connectionWithDatabaseParams := "user=docker password=docker sslmode=disable host=db dbname=usersync"
	// connectionParams := fmt.Sprintf("user=%v password=%v sslmode=disable host=%v", Constants.DBUser, Constants.DBPassword, Constants.DBHost)
	// connectionWithDatabaseParams := fmt.Sprintf("user=%v password=%v sslmode=disable host=%v dbname=%v", Constants.DBUser, Constants.DBPassword, Constants.DBHost, Constants.DBName)

	// for i := 0; i < 30; i++ {
	// log.Printf("trying to connect to database, attempt: %v", i+1)
	log.Println(os.Getenv("DATABASE_URL"))
	database, err = gorm.Open("postgres", os.Getenv("DATABASE_URL"))
	// 	if err != nil {
	// 		time.Sleep(1 * time.Second)
	// 		continue
	// 	}

	// 	// log.Printf("creating database '%v'", "usersync")
	// 	// database.Exec("CREATE DATABASE " + "usersync")

	// 	// database, err = gorm.Open("postgres", connectionWithDatabaseParams)
	// 	// if err == nil {
	// 	// 	break
	// 	// }
	// }
	if err != nil {
		return
	}

	database.AutoMigrate(&models.SlackProfile{}, &models.User{})
	database.Model(&models.User{}).AddForeignKey("profile_id", "profile(id)", "CASCADE", "CASCADE")

	return
}

// InsertOrUpdateUser inserts or updates a user into the database
var InsertOrUpdateUser = func(user models.User) error {
	_, err := GetUserBySlackID(user.SlackID)
	if err != nil {
		return database.Save(&user).Error
	}
	return UpdateUser(user)
}

// UpdateUser updates a user informations
func UpdateUser(user models.User) error {
	err := database.Model(&models.User{}).Where("slack_id=?", user.SlackID).Omit("profile").Updates(user).Error
	if err != nil {
		return err
	}

	u, _ := GetUserBySlackID(user.SlackID)
	err = database.Model(&models.SlackProfile{}).Where("id=?", u.Profile.ID).Updates(user.Profile).Error
	if err != nil {
		return err
	}
	return nil
}

// GetUserBySlackID finds a user given a slackID
func GetUserBySlackID(slackID string) (*models.User, error) {
	user := &models.User{}
	err := database.Where("slack_id=?", slackID).Find(&user).Error
	if err != nil {
		return nil, err
	}
	err = database.Model(&user).Association("profile").Find(&user.Profile).Error
	return user, err
}
