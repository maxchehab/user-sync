package main

import (
	"models"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

var database *gorm.DB

// IntializeDatabase creates tables for all of the models
func IntializeDatabase() (err error) {
	database, err = gorm.Open("postgres", Constants.DBUrl)
	if err != nil {
		return err
	}
	database.AutoMigrate(&models.SlackProfile{}, &models.User{})

	return nil
}

// InsertOrUpdateUser inserts or updates a user into the database
var InsertOrUpdateUser = func(user models.User) error {
	_, err := GetUserBySlackID(user.SlackID)
	if err != nil {
		return database.Save(&user).Error
	}
	return UpdateUser(user)
}

// GetAllUsers gets all the users and returns them
var GetAllUsers = func() ([]*models.User, error) {
	var shallowUsers = []models.User{}
	err := database.Find(&shallowUsers).Error
	if err != nil {
		return nil, err
	}

	var users = []*models.User{}
	for _, u := range shallowUsers {
		user, err := GetUserBySlackID(u.SlackID)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}

	return users, nil
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
