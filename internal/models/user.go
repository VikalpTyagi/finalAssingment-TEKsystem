// * User model
package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name     string `json:"name"`
	Email    string `json:"email"`
	PassHash string `json:"-"`
}
type NewUser struct {
	gorm.Model
	Name     string `json:"name" validate:"required"`
	Email    string `json:"email" gorm:"unique;not null" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

// * =============================== Applicants struck ==============================================
type Applicant struct {
	gorm.Model
	Name           string 			`json:"name" validate:"required"`
	JobId            uint			`json:"job" validate:"required"`
	Experience     uint            `json:"experience" validate:"required"`
	Min_NP         uint            `json:"min-NP" validate:"required"`
	Max_NP         uint            `json:"max-NP" validate:"required"`
	Budget         uint            `json:"salary" validate:"required"`
	Locations      Location      `json:"locations" validate:"required" gorm:"many2many:job_location;"`
	Stack          []Skill         `json:"skills" validate:"required" gorm:"many2many:job_skill;"`
	WorkMode       string          `json:"workMode" validate:"required"`
	Qualifications Qualification `json:"qualification" validate:"required" gorm:"many2many:job_qualifications;"`
	Shift          string          `json:"shift" validate:"required"`
}

type ApplicantRespo struct {
	Id				uint
	Name           string
	JobId			uint
}
