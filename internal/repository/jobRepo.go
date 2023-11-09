package repository

import (
	"context"
	"finalAssing/internal/models"
	"strconv"

	"gorm.io/gorm"
)
//@ modified code according new models
func (r *ReposStruct) SaveJobsByCompanyId(jobs []models.JobReq, compId string) ([]models.JobRespo, error) {
	companyId, err := strconv.ParseUint(compId, 10, 64)
	if err != nil {
		return nil, err
	}
	var listOfJobs []models.JobRespo
	var gormJobs []models.Job
	for i := 0; i < len(jobs); i++ {
		gormData := models.Job{
			Name:           jobs[i].Name,
			Field:          jobs[i].Field,
			Experience:     jobs[i].Experience,
			Min_NP:         jobs[i].Min_NP,
			Max_NP:         jobs[i].Max_NP,
			Budget:         jobs[i].Budget,
			Locations:      intToLocation(jobs[i].Locations),
			Stack:          intToSkill(jobs[i].Stack),
			WorkMode:       jobs[i].WorkMode,
			Description:    jobs[i].Description,
			MinExp:         jobs[i].MinExp,
			Qualifications: intToQuali(jobs[i].Qualifications),
			Shift:          jobs[i].Shift,
			CompanyId:      companyId,
		}
		
		err := r.db.Create(&gormData).Error
		if err != nil {
			return nil, err
		}
		gormJobs = append(gormJobs, gormData)
		respoData := models.JobRespo{
			Id: gormJobs[i].ID,
		}
		listOfJobs = append(listOfJobs, respoData)
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

func (r *ReposStruct) FetchAllJobs(ctx context.Context) ([]models.Job, error) {
	var listJobs []models.Job
	tx := r.db.WithContext(ctx)
	err := tx.Find(&listJobs).Error
	if err != nil {
		return nil, err
	}

	return listJobs, nil
}

func intToLocation(slice []uint) []models.Location {
	var locRespo []models.Location
	for _, s := range slice {
		loc := models.Location{
			Model: gorm.Model{ID: s},
		}
		locRespo = append(locRespo, loc)
	}
	return locRespo
}

func intToSkill(slice []uint) []models.Skill {
	var locRespo []models.Skill
	for _, s := range slice {
		loc := models.Skill{
			Model: gorm.Model{ID: s},
		}
		locRespo = append(locRespo, loc)
	}
	return locRespo
}
func intToQuali(slice []uint) []models.Qualification {
	var skillRespo []models.Qualification
	for _, s := range slice {
		skil := models.Qualification{
			Model: gorm.Model{ID: s},
		}
		skillRespo = append(skillRespo, skil)
	}
	return skillRespo
}
