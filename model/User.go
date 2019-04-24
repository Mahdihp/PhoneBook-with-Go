package model

import "github.com/jinzhu/gorm"

type User struct {
	gorm.Model

	Username    string `gorm:"type:varchar(255);unique;not null`
	Pasword     string `gorm:"type:varchar(255)`
	Email       string `gorm:"type:varchar(255)`
	Regfirebase string `gorm:"type:varchar(255)`
}

func NewUser(username string, pasword string, email string, regfirebase string) *User {
	return &User{Username: username, Pasword: pasword, Email: email, Regfirebase: regfirebase}
}

func (User) TableName() string {
	return "users"
}
func DBMigrate(db *gorm.DB) *gorm.DB {
	db.AutoMigrate(&User{})
	return db
}
