package models

import "github.com/jinzhu/gorm"

type Favorite struct {
	gorm.Model
	UserID    string
	AnnonceID uint
}
