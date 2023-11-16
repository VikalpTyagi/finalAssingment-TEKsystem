package services

import (
	"context"
	"finalAssing/internal/cacheier"
	"finalAssing/internal/models"
	"finalAssing/internal/repository"

	"github.com/golang-jwt/jwt/v5"
)

//go:generate mockgen -source service.go -destination service_mock.go -package services

type Service interface {
	GetAllJobs(ctx context.Context) ([]models.Job, error)
	GetJobById(ctx context.Context, jobId string) (models.Job, error)
	FetchJobByCompanyId(ctx context.Context, companyId string) ([]models.Job, error)
	JobByCompanyId(jobs []models.JobReq, compId string) ([]models.JobRespo, error)
	FetchCompanyByID(ctx context.Context, companyId string) (models.Company, error)
	ViewCompanies(ctx context.Context) ([]models.Company, error)
	CreateCompany(ctx context.Context, newComp models.Company) (models.Company, error)
	CreateUser(ctx context.Context, nu models.NewUser) (models.User, error)
	Authenticate(ctx context.Context, email, password string) (jwt.RegisteredClaims, error)

	FIlterApplication(ctx context.Context, applicantList []*models.ApplicantReq) ([]*models.ApplicantRespo, error)
}

type Store struct {
	Repo  repository.RepoInterface
	Cache cacheier.RedInterface
}

func NewStore(repo repository.RepoInterface, redis cacheier.RedInterface) Service {
	return &Store{Repo: repo, Cache: redis}
}
