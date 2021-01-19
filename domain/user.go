package domain

import "gorm.io/gorm"

type User struct {
	ID   int `json:"id" gorm:"primary_key"`
	Name string
	gorm.Model
}
