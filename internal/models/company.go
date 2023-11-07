package models

import "gorm.io/gorm"

type Company struct {
	gorm.Model
	Name string `json:"companyName" validate:"required"`
	City string `json:"city" validate:"required"`
	Jobs []Job  `json:"jobs,omitempty" gorm:"foreignKey:CompanyId"`
}
type Job struct {
	gorm.Model
	Name           string          `json:"title" validate:"required"`
	Field          string          `json:"field" validate:"required"`
	Experience     uint            `json:"experience" validate:"required"`
	Min_NP         uint            `json:"min-NP" validate:"required"`
	Max_NP         uint            `json:"max-NP" validate:"required"`
	Budget         uint            `json:"salary" validate:"required"`
	Locations      []Location      `json:"locations" validate:"required" gorm:"many2many:job_location;"`
	Stack          []Skill         `json:"skills" validate:"required" gorm:"many2many:job_skill;"`
	WorkMode       string          `json:"workMode" validate:"required"`
	Description    string          `json:"desc" validate:"required"`
	MinExp         uint            `json:"minExp" validate:"required"`
	Qualifications []Qualification `json:"qualification" validate:"required" gorm:"many2many:job_qualifications;"`
	Shift          string          `json:"shift" validate:"required"`
	CompanyId      uint64          `json:"companyId"`
}

type Location struct {
	gorm.Model
	City    string `json:"city" `
	Country string `json:"country"`
}
type Skill struct {
	gorm.Model
	Sname       string `json:"sname"`
	Proficiency int    `json:"proficiency"`
}
type Qualification struct {
	gorm.Model
	Degree     string  `json:"degree"`
	Percentage float64 `json:"percentage"`
}

