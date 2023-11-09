// // * Job Service
package services

import (
	"context"
	"finalAssing/internal/models"
)

//@ modified according to new models
func (s *Store) JobByCompanyId(jobs []models.JobReq, compId string) ([]models.JobRespo, error) {
	listOfjobs, err := s.Repo.SaveJobsByCompanyId(jobs, compId)
	if err != nil {
		return nil, err
	}
	return listOfjobs, nil
}

func (s *Store) FetchJobByCompanyId(ctx context.Context, companyId string) ([]models.Job, error) {
	listOfJob, err := s.Repo.GetJobsByCId(ctx, companyId)
	if err != nil {
		return nil, err
	}

	return listOfJob, nil
}

func (s *Store) GetJobById(ctx context.Context, jobId string) (models.Job, error) {
	jobData, err := s.Repo.FetchByJobId(ctx, jobId)
	if err != nil {
		return models.Job{}, err
	}

	return jobData, nil
}

func (s *Store) GetAllJobs(ctx context.Context) ([]models.Job, error) {
	listJobs, err := s.Repo.FetchAllJobs(ctx)
	if err != nil {
		return []models.Job{}, err
	}
	return listJobs, nil
}
