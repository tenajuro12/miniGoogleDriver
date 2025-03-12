package model

import "gorm.io/gorm"

type File struct {
	gorm.Model
	Name string `gorm:"uniqueIndex;not null"`
	Size int64  `gorm:"not null"`
}
