// // * Job Service
package services

import (
	"finalAssing/internal/models"
	"strconv"

	"golang.org/x/net/context"
)

func (s *DbConnStruct) jobByCompanyId(ctx context.Context,jobs []models.Job,compId string,)([]models.Job,error){
	companyId,err := strconv.ParseUint(compId,10,64)
	if err!= nil{
		return nil,err
	}
	for _,j := range jobs{

		job := models.Job{
			Name: j.Name,
			Field:j.Field,
			Experience:j.Experience,
			CompanyId: companyId,

		}
		err := s.db.Create(&job).Error
	if err != nil {
		return nil, err
	}
	}
	return jobs,nil
}
