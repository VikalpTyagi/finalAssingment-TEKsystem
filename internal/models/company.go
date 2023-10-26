package models

import "gorm.io/gorm"

type Company struct {
	gorm.Model
	Name string	`Json:"Company Name" validate:"required"`
	City string	`Json:"City" validate:"required"`
	Jobs []Job `Json:"Jobs" gorm:"foreignKey:ID"`
}
type Job struct {
	gorm.Model
	Name  string
	Field string
	Experience uint
}
