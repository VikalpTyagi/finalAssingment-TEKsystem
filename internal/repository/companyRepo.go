package repository

import (
	"context"
	"finalAssing/internal/models"
)

func (r *ReposStruct) SaveCompany(newComp models.Company) (models.Company, error) {
	comp := models.Company{
		Name: newComp.Name,
		City: newComp.City,
		Jobs: newComp.Jobs,
	}
	err := r.db.Create(&comp).Error
	if err != nil {
		return models.Company{}, err
	}
	return comp, nil
}

func (r *ReposStruct) FetchAllCompanies(ctx context.Context) ([]models.Company, error) {
	var listComp []models.Company
	tx := r.db.WithContext(ctx)
	err := tx.Find(&listComp).Error
	if err != nil {
		return nil, err
	}
	return listComp, nil
}

func (r *ReposStruct) GetCompaniesById(ctx context.Context, companyId string) (models.Company, error) {
	var comp models.Company
	tx := r.db.WithContext(ctx).Where("ID = ?", companyId)
	err := tx.Find(&comp).Error
	if err != nil {
		return models.Company{}, err
	}
	return comp, nil
}
