package repository

import (
	"context"
	"errors"
	"finalAssing/internal/models"

	"gorm.io/gorm"
)

//go:generate mockgen -source repo.go -destination mockRepo/repository_mock.go -package repository

type RepoInterface interface {
	SaveUser(userData models.NewUser, hashedPass []byte) (models.User, error)
	CheckEmail(email, password string) (models.User, error)

	SaveCompany(newComp models.Company) (models.Company, error)
	FetchAllCompanies(ctx context.Context) ([]models.Company, error)
	GetCompaniesById(ctx context.Context, companyId string) (models.Company , error)

	SaveJobsByCompanyId(jobs []models.Job, compId string) ([]models.Job, error)
	GetJobsByCId(ctx context.Context, companyId string) ([]models.Job, error)
	FetchByJobId(ctx context.Context, jobId string) (models.Job, error)
	FetchAllJobs(ctx context.Context) ([]models.Job,error)
}

type ReposStruct struct {
	db *gorm.DB
}

func NewRepo(dataB *gorm.DB) (*ReposStruct, error) {
	if dataB == nil {
		return nil, errors.New("db is empty")
	}
	return &ReposStruct{db: dataB}, nil
}
