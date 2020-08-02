package models

import "github.com/jinzhu/gorm"

//Product @createdAt 01/08/2020 @createdBy allanfvc
type Product struct {
	gorm.Model
	Name string `json:"name"`
}
