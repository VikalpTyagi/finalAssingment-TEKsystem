package models

import "gorm.io/gorm"

type Company struct {
	gorm.Model
	Name string `json:"companyName" validate:"required"`
	City string `json:"city" validate:"required"`
	Jobs []Job  `json:"jobs,omitempty" gorm:"foreignKey:CompanyId"`
}

