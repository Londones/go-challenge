package models

import (
	"github.com/jinzhu/gorm"
)

type Annonce struct {
	gorm.Model
	Description *string `gorm:"type:varchar(250)"`
	UserID      uint
	Cats        []Cats
	Favorite    []Favorite
	Rating      []Rating
}