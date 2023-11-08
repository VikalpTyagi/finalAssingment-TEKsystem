package repository

import (
	"finalAssing/internal/models"

	"github.com/rs/zerolog/log"
)

func (r *ReposStruct) ApplicantsFilter(jobId uint) (*models.Job, error) {
	var jobData models.Job
	err := r.db.Preload("Location").Preload("Skill").Preload("Qualification").Where("ID = ?", jobId).Find(&jobData).Error
	if err != nil {
		log.Error().Err(err).Msg("Problem in fetching joba data")
		return nil, err
	}
	return &jobData, nil
}

// func (r *ReposStruct) ApplicantsFilter(ctx context.Context, applicantList []*models.ApplicantReq) ([]*models.ApplicantRespo, error) {
// 	var job models.Job
// 	var sApplicant []*models.ApplicantRespo
// 	ch := make(chan *models.ApplicantReq)
// 	wg := &sync.WaitGroup{}
// 	for _, data := range applicantList {
// 		tx := r.db.WithContext(ctx).Where("ID = ?", data.JobId)
// 		err := tx.First(&job).Error
// 		if err != nil {
// 			log.Error().Err(err).Str("Error", "Job Id not found in db").Send()
// 			return nil,err
// 		}
// 		wg.Add(1)
// 		go func(job models.Job, data *models.ApplicantReq) {
// 			defer wg.Done()
// 			// for _, q := range job.Qualifications {
// 			// 	if q == data.Qualifications {
// 			// 		break
// 			// 	}
// 			// }
// 			if !(job.MinExp <= data.Experience && data.Experience <= job.Experience) {
// 				log.Error().Err(errors.New("experience requirments not met")).Send()
// 				return
// 			}else if data.Budget > job.Budget {
// 				log.Error().Err(errors.New("salary requirments not met")).Send()
// 			return
// 			}else if !(job.Max_NP >= data.Max_NP && job.Min_NP >= data.Min_NP) {
// 				log.Error().Err(errors.New("notice periode requirments not met")).Send()
// 				return
// 			}else if job.WorkMode != data.WorkMode {
// 				log.Error().Err(errors.New("work mode requirments not met")).Send()
// 				return
// 			}else if job.Shift != data.Shift {
// 				log.Error().Err(errors.New("shift requirments not met")).Send()
// 				return
// 			}
// 			// for _,s := range job.Stack {
// 			// 	for _,as := range data.Stack{

// 			// 	}
// 			// }
// 			ch <- data
// 		}(job, data)

// 		go func() {
// 			wg.Wait()
// 			close(ch)
// 		}()

// 		for appl := range ch {
// 			sApplicant = append(sApplicant, &models.ApplicantRespo{
// 				Name:  appl.Name,
// 				JobId: appl.JobId,
// 			})
// 		}
// 	}

// 	return sApplicant, nil
// }
