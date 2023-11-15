package repository

import (
	"finalAssing/internal/models"
	"fmt"

	"github.com/rs/zerolog/log"
)

func (r *ReposStruct) ApplicantsFilter(jobId uint) (*models.Job, error) {
	var jobData models.Job
	fmt.Println("job data fetched from db")
	err := r.db.Preload("Locations").Preload("Stack").Preload("Qualifications").Where("ID = ?", jobId).Find(&jobData).Error
	if err != nil {
		log.Error().Err(err).Msg("Problem in fetching joba data")
		return nil, err
	}
	return &jobData, nil
}
