// * User model
package models

import "gorm.io/gorm"

type User struct{
	gorm.Model
	UserID uint `Json: "ID"`
	Name string	`Json:"Name"`
	Email string	`Json:"Email"`
	PassHash string	`Json:"-"`
}
type NewUser struct{
	gorm.Model
	UserID uint `json:"ID" validate:"required"`
	Name string `json:"name" validate:"required"`
	Email string `json:"Email" validate:"required"`
}