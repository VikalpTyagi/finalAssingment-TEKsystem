package services

import (
	"errors"
	"finalAssing/internal/models"

	"gorm.io/gorm"
)

type DbConnStruct struct {
	db *gorm.DB
}

func NewConn(dbInstance *gorm.DB) (*DbConnStruct, error) {
	if dbInstance == nil {
		return nil, errors.New("provide the databse instance")
	}
	return &DbConnStruct{db: dbInstance}, nil
}

func (s *DbConnStruct) AutoMigrate() error {

	// AutoMigrate function will ONLY create tables, missing columns and missing indexes, and WON'T change existing column's type or delete unused columns
	err := s.db.Migrator().AutoMigrate(&models.User{}, &models.Company{}, &models.Job{})
	if err != nil {
		// If there is an error while migrating, log the error message and stop the program
		return err
	}
	return nil
}
