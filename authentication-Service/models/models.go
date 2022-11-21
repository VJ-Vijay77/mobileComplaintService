package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
	Password  string `json:"password"`
	Phone     int    `json:"phone"`
}
