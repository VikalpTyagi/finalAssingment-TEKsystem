package services

import (
	"context"
	"finalAssing/internal/models"
)

func (s *Store) CreateCompany(ctx context.Context, newComp models.Company) (models.Company, error) {
comp,err := s.Repo.SaveCompany(newComp)
	if err != nil {
		return models.Company{}, err
	}
	return comp, nil
}

func (s *Store) ViewCompanies(ctx context.Context)([]models.Company,error){
	listComp,err := s.Repo.FetchAllCompanies(ctx)
	if err != nil {
		return nil, err
	}
	return listComp, nil
}

func (s *Store) FetchCompanyByID(ctx context.Context, companyId string) (models.Company , error) {
	comp,err := s.Repo.GetCompaniesById(ctx,companyId)
	if err != nil {
		return models.Company{}, err
	}
	return comp, nil

}