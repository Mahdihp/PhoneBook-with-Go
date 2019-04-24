package model

import "github.com/jinzhu/gorm"

type Video struct {
	gorm.Model

	Title       string
	Description string
	Url         string
}

func NewVideo(title string, description string, url string) *Video {
	return &Video{Title: title, Description: description, Url: url}
}
func (Video) TableName() string {
	return "videos"
}
