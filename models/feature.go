package models

import "github.com/jinzhu/gorm"

type Feature struct {
	gorm.Model
	Name string `json:"name"`
}

