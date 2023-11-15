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

	wg := &sync.WaitGroup{}
	var ch = make(chan *models.ApplicantReq)
	for _, applicant := range applicantList {
		wg.Add(1)
		go func(appl *models.ApplicantReq) {
			defer wg.Done()

			err := s.Filter(ctx, appl) //@ Calls a function(Scroll down) for comparision and filteration
			if err != nil {
				log.Error().Err(err).Msg("Error in filter of application")

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

	return response, nil
}

func (s *Store) Filter(ctx context.Context, appl *models.ApplicantReq) error {

	// jobData, err := s.Repo.ApplicantsFilter(appl.JobId) //@ fetching job data
	jobData, err := s.fetchJobData(ctx, appl.JobId)
	if err != nil {
		log.Error().Err(err).Interface("Job ID", appl.JobId).Send()
		return errors.New("data not fetched from db")
	}
	if jobData.Budget < appl.Budget {
		log.Error().Err(errors.New("budget requirments not met")).Interface("applicant ID", appl.Name).Send()
		return errors.New("budget requirments not met")
	}
	if jobData.Experience < appl.Experience || appl.Experience < jobData.MinExp {
		log.Error().Err(errors.New("experience requirments not met")).Interface("applicant ID", appl.Name).Send()
		return errors.New("experience requirments not met")
	}
	if jobData.Max_NP < appl.Max_NP || appl.Max_NP < jobData.Min_NP {
		log.Error().Err(errors.New("notice periode requirments not met")).Interface("applicant ID", appl.Name).Send()
		return errors.New("notice periode requirments not met")
	}
	if jobData.Shift != appl.Shift {
		log.Error().Err(errors.New("working shift requirments not met")).Interface("applicant ID", appl.Name).Send()
		return errors.New("working shift requirments not met")
	}
	if jobData.WorkMode != appl.WorkMode {
		log.Error().Err(errors.New("work mode requirments not met")).Interface("applicant ID", appl.Name).Send()
		return errors.New("work mode requirments not met")
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
		return errors.New("qualification requirments not met")
	}
	available := false
	for _, j := range jobData.Locations {
		for _, a := range appl.Locations {
			if j.Model.ID == a {
				available = true
			}
		}
	}
	if !available {
		log.Error().Err(errors.New("location requirments not met")).Interface("applicant ID", appl.Name).Send()
		return errors.New("location requirments not met")
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
		return errors.New("skills requirments not met")
	}

	return nil
}

func (s *Store) fetchJobData(ctx context.Context, jobId uint) (*models.Job, error) {
	data, err := s.Cache.FetchJobData(ctx, jobId)
	if err != nil {
		data, err = s.Repo.ApplicantsFilter(jobId)
		if err != nil {
			return nil, err
		}
		err = s.Cache.AddJobData(ctx, jobId, data)
		if err != nil {
			return nil, err
		}
		// return data, nil
	}
	return data, nil
}
