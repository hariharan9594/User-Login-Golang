package models

import "github.com/jinzhu/gorm"

type User struct {
	gorm.Model
	UserName string `json:"username" gorm:"type:varchar(100);unique_index"`
	Password string `json:"password"`
}
