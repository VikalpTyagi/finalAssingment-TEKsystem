package repository

import (
	"context"
	"errors"
	"finalAssing/internal/models"
	"sync"

	"github.com/rs/zerolog/log"
)

func (r *ReposStruct) ApplicantsFilter(ctx context.Context, applicantList []*models.Applicant) ([]*models.ApplicantRespo, error) {
	var job models.Job
	var sApplicant []*models.ApplicantRespo
	ch := make(chan *models.Applicant)
	wg := &sync.WaitGroup{}
	for _, data := range applicantList {
		tx := r.db.WithContext(ctx).Where("ID = ?", data.JobId)
		err := tx.First(&job).Error
		wg.Add(1)
		go func(job models.Job, data *models.Applicant) {
			defer wg.Done()
			if err != nil {
				log.Error().Err(err).Str("Error", "Job Id not found in db").Send()
				ch <- nil
			}
			// for _, q := range job.Qualifications {
			// 	if q == data.Qualifications {
			// 		break
			// 	}
			// }
			if !(job.MinExp <= data.Experience && data.Experience <= job.Experience) {
				log.Error().Err(errors.New("experience requirments not met")).Send()
				ch <- nil
			}
			if !(job.Budget >= data.Budget) {
				log.Error().Err(errors.New("salary requirments not met")).Send()
				ch <- nil
			}
			if !(job.Max_NP >= data.Max_NP && job.Min_NP >= data.Min_NP) {
				log.Error().Err(errors.New("notice periode requirments not met")).Send()
				ch <- nil
			}
			if job.WorkMode != data.WorkMode {
				log.Error().Err(errors.New("work mode requirments not met")).Send()
				ch <- nil
			}
			if job.Shift != data.Shift {
				log.Error().Err(errors.New("shift requirments not met")).Send()
				ch <- nil
			}
			// for _,s := range job.Stack {
			// 	for _,as := range data.Stack{

			// 	}
			// }
			ch <- data
		}(job, data)

		go func ()  {
			wg.Wait()
			close(ch)
		}

		for appl := range ch {
			sApplicant = append(sApplicant, &models.ApplicantRespo{
				Id:    appl.ID,
				Name:  appl.Name,
				JobId: appl.JobId,
			})
		}
	}

	return sApplicant, nil
}
