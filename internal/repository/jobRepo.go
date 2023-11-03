package repository

import (
	"context"
	"finalAssing/internal/models"
	"strconv"
)

func (r *ReposStruct) SaveJobsByCompanyId(jobs []models.Job, compId string) ([]models.Job, error) {
	companyId, err := strconv.ParseUint(compId, 10, 64)
	if err != nil {
		return nil, err
	}
	var listOfJobs []models.Job
	for _, j := range jobs {
		j.CompanyId=companyId
		listOfJobs = append(listOfJobs, j)
		err := r.db.Create(&j).Error
		if err != nil {
			return nil, err
		}
	}
	return listOfJobs, nil
}

func (r *ReposStruct) GetJobsByCId(ctx context.Context, companyId string) ([]models.Job, error) {
	var listOfJobs []models.Job
	tx := r.db.WithContext(ctx).Where("company_id = ?", companyId)
	err := tx.Find(&listOfJobs).Error
	if err != nil {
		return nil, err
	}

	return listOfJobs, nil
}

func (r *ReposStruct) FetchByJobId(ctx context.Context, jobId string) (models.Job, error) {
	var jobData models.Job
	tx := r.db.WithContext(ctx).Where("ID = ?", jobId)
	err := tx.Find(&jobData).Error
	if err != nil {
		return models.Job{}, err
	}

	return jobData, nil
}

func (r *ReposStruct) FetchAllJobs(ctx context.Context) ([]models.Job,error){
	var listJobs []models.Job
	tx := r.db.WithContext(ctx)
	err := tx.Find(&listJobs).Error
	if err != nil {
		return nil, err
	}

	return listJobs, nil
}