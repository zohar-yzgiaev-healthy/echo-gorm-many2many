package models

import "github.com/jinzhu/gorm"

type Place struct {
	gorm.Model
	Name string `json:"name"`
	Features []Feature `gorm:"many2many:place_features;"`
	Photos []Photo `gorm:"many2many:place_photos;"`
}
