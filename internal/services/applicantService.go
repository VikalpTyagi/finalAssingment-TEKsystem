package services

import (
	"context"
	"errors"
	"finalAssing/internal/models"
	"sync"

	"github.com/rs/zerolog/log"
)

func (s *Store) FIlterApplication(ctx context.Context, applicantList []*models.ApplicantReq) ([]*models.ApplicantRespo, error) {
	var response []*models.ApplicantRespo
	ch := make(chan *models.ApplicantReq)
	wg := &sync.WaitGroup{}

	for _, applicant := range applicantList {
		wg.Add(1)
		go func(appl *models.ApplicantReq) {
			defer wg.Done()
			jobData, err := s.Repo.ApplicantsFilter(appl.JobId) //@ fetching job data
			if err != nil {
				log.Error().Err(err).Interface("Job ID", appl.JobId).Send()
				return
			}
			if jobData.Budget < appl.Budget {
				log.Error().Err(errors.New("budget requirments not met")).Interface("applicant ID", appl.Name).Send()
				return
			}
			if jobData.Experience < appl.Experience || appl.Experience < jobData.MinExp {
				log.Error().Err(errors.New("experience requirments not met")).Interface("applicant ID", appl.Name).Send()
				return
			}
			if jobData.Max_NP < appl.Max_NP ||  appl.Max_NP < jobData.Min_NP {
				log.Error().Err(errors.New("notice periode requirments not met")).Interface("applicant ID", appl.Name).Send()
				return
			}
			if jobData.Shift != appl.Shift {
				log.Error().Err(errors.New("working shift requirments not met")).Interface("applicant ID", appl.Name).Send()
				return
			}
			if jobData.WorkMode != appl.WorkMode {
				log.Error().Err(errors.New("work mode requirments not met")).Interface("applicant ID", appl.Name).Send()
				return
			}
			if jobData.WorkMode != appl.WorkMode {
				log.Error().Err(errors.New("work mode requirments not met")).Interface("applicant ID", appl.Name).Send()
				return
			}
			var passed bool
			for _, j := range jobData.Qualifications {
				for _, a := range appl.Qualifications {
					if j.Model.ID == a {
						passed = true
					}
				}
			}
			if !passed {
				log.Error().Err(errors.New("qualification requirments not met")).Interface("applicant ID", appl.Name).Send()
				return
			}
			var available bool
			for _, j := range jobData.Locations {
				for _, a := range appl.Locations {
					if j.Model.ID == a {
						available = true
					}
				}
			}
			if !available {
				log.Error().Err(errors.New("location requirments not met")).Interface("applicant ID", appl.Name).Send()
				return
			}
			count := 0
			for _, j := range jobData.Stack {
				for _, a := range appl.Stack {
					if j.Model.ID == a {
						count++
					}
				}
			}
			if count < (len(jobData.Stack) / 2) {
				log.Error().Err(errors.New("skills requirments not met")).Interface("applicant ID", appl.Name).Send()
				return
			}

			ch <- appl
		}(applicant)
	}

	go func() {
		wg.Wait()
		close(ch)
	}()

	for data := range ch {
		respo := models.ApplicantRespo{
			Name:  data.Name,
			JobId: data.JobId,
		}
		response = append(response, &respo)
	}
	if response == nil {
		log.Info().Err(errors.New("no candidates passed the requirments")).Send()
	}
	return response, nil
}
