// // * Job Service
package services

import (
	"context"
	"finalAssing/internal/models"
	"strconv"
)

func (s *DbConnStruct) JobByCompanyId(jobs []models.Job, compId string) ([]models.Job, error) {
	companyId, err := strconv.ParseUint(compId, 10, 64)
	if err != nil {
		return nil, err
	}
	for _, j := range jobs {
		job := models.Job{
			Name:       j.Name,
			Field:      j.Field,
			Experience: j.Experience,
			CompanyId:  companyId,
		}
		err := s.db.Create(&job).Error
		if err != nil {
			return nil, err
		}
	}
	return jobs, nil
}

func (s *DbConnStruct) FetchJobByCompanyId(ctx context.Context, companyId string) ([]models.Job, error) {
	var listOfJobs []models.Job
	tx := s.db.WithContext(ctx).Where("company_id = ?", companyId)
	err := tx.Find(&listOfJobs).Error
	if err != nil {
		return nil, err
	}

	return listOfJobs, nil
}

func (s *DbConnStruct) GetJobById(ctx context.Context, jobId string) (models.Job,error){
	var jobData models.Job
	tx := s.db.WithContext(ctx).Where("ID = ?", jobId)
	err := tx.Find(&jobData).Error
	if err != nil {
		return models.Job{}, err
	}

	return jobData, nil
}

func (s *DbConnStruct) GetAllJobs(ctx context.Context) ([]models.Job,error){
	var listJobs []models.Job
	tx := s.db.WithContext(ctx)
	err := tx.Find(&listJobs).Error
	if err != nil {
		return nil, err
	}

	return listJobs, nil
}