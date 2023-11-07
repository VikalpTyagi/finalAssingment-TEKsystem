package database

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

func AutoMigrate(str *DbConnStruct) error {
	err := str.db.Migrator().AutoMigrate(&models.User{}, &models.Company{}, &models.Job{}, &models.Location{},&models.Qualification{},&models.Skill{})
	if err != nil {
		return err
	}
	return nil
}
