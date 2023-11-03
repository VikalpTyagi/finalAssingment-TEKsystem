package services

import (
	"context"
	"finalAssing/internal/models"
	"finalAssing/internal/repository"

	"github.com/golang-jwt/jwt/v5"
)

//go:generate mockgen -source service.go -destination mockmodels/service_mock.go -package mockmodels

type Service interface {
	GetAllJobs(ctx context.Context) ([]models.Job, error)
	GetJobById(ctx context.Context, jobId string) (models.Job, error)
	FetchJobByCompanyId(ctx context.Context, companyId string) ([]models.Job, error)
	JobByCompanyId(jobs []models.Job, compId string) ([]models.Job, error)
	FetchCompanyByID(ctx context.Context, companyId string) (models.Company, error)
	ViewCompanies(ctx context.Context) ([]models.Company, error)
	CreateCompany(ctx context.Context, newComp models.Company) (models.Company, error)
	CreateUser(ctx context.Context, nu models.NewUser) (models.User, error)
	Authenticate(ctx context.Context, email, password string) (jwt.RegisteredClaims,error)
}

type Store struct {
	Repo *repository.ReposStruct
}

func NewStore(repo *repository.ReposStruct) *Store {
	return &Store{Repo: repo}
}
