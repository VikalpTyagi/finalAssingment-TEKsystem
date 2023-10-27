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
	err := s.db.Create(&comp).Error
	if err != nil {
		return models.Company{}, err
	}
	return comp, nil
}

func (s *DbConnStruct) ViewCompanies(ctx context.Context)([]models.Company,error){
	var listComp []models.Company
	tx := s.db.WithContext(ctx)
	err := tx.Find(&listComp).Error
	if err != nil {
		return nil, err
	}

	return listComp, nil
}

func (s *DbConnStruct) FetchCompanyByID(ctx context.Context, companyId string) (models.Company , error) {
	var comp models.Company
	tx := s.db.WithContext(ctx).Where("ID = ?", companyId)
	err := tx.Find(&comp).Error
	if err != nil {
		return models.Company{}, err
	}

	return comp, nil

}