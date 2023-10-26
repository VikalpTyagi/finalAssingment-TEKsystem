// * User model
package models

import "gorm.io/gorm"

type User struct{
	gorm.Model
	Name string	`Json:"Name"`
	Email string	`Json:"Email"`
	PassHash string	`Json:"-"`
}
type NewUser struct{
	gorm.Model
	Name string `json:"Name" validate:"required"`
	Email string `json:"Email" validate:"required,email"`
	Password string `Json:"Password" validate:"required`
}