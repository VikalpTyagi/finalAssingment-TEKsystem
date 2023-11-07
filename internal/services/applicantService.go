package services

import (
	"context"
	"finalAssing/internal/models"
)

func (s *Store) FIlterApplication(ctx context.Context,applicantList *models.Applicant) (models.ApplicantRespo,error){
	selectedApplicants, err := s.ApplicantsFilter(ctx,applicantList)
	if err !=nil{
		return nil,err
	}
	return selectedApplicants,nil
}