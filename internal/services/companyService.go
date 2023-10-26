package services

import (
	"context"
	"finalAssing/internal/models"
)

func (s *DbConnStruct) CreateCompany(ctx context.Context, newComp models.Company) (models.Company, error) {

	comp := models.Company{
		Name: newComp.Name,
		City: newComp.City,
		Jobs: newComp.Jobs,
	}

	// We attempt to create the new User record in the database.
	err := s.db.Create(&comp).Error
	if err != nil {
		return models.Company{}, err
	}

	// Successfully created the record, return the user.
	return comp, nil
}
