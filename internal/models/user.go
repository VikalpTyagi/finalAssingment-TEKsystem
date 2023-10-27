// * User model
package models

import "gorm.io/gorm"

type User struct{
	gorm.Model
	Name string	`json:"name"`
	Email string	`json:"email"`
	PassHash string	`json:"-"`
}
type NewUser struct{
	gorm.Model
	Name string `json:"name" validate:"required"`
	Email string `json:"email" validate:"required,email"`
	Password string `Json:"password" validate:"required`
}