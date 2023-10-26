package models

import "gorm.io/gorm"

type Company struct {
	gorm.Model
	CompanyId uint
	Name string
	City string
	Jobs []Job `gorm:"foreignKey:JobId"`
}
type Job struct {
	gorm.Model
	JobId uint
	Name  string
	Field string
	Experience uint
}
