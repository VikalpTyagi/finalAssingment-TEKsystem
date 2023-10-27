// * job related handlers endpoint
package handlers

import (
	"encoding/json"
	"finalAssing/internal/middleware"
	"finalAssing/internal/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
)

func (h *handler) addJobsById(c *gin.Context) {
	ctx := c.Request.Context()
	trackerId, ok := ctx.Value(middleware.TrackerIdKey).(string)
	if !ok {
		log.Error().Msg("TrackerId missing from context")
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"msg": http.StatusText(http.StatusInternalServerError)})
		return
	}
	compId:=c.Param("ID")
	var jobs []models.Job
	err := json.NewDecoder(c.Request.Body).Decode(&jobs)
	if err != nil {
		// If there is an error in decoding, log the error and return
		log.Error().Err(err).Str("Tracker Id", trackerId)
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"msg": http.StatusText(http.StatusInternalServerError)})
		return
	}
	// Create a new validator and validate the NewUser variable
	// validate := validator.New()
	// err = validate.Struct(jobs)
	// if err != nil {
	// 	// If validation fails, log the error and return
	// 	log.Error().Err(err).Str("Tracker Id", trackerId).Send()
	// 	c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"msg": "please provide Name of company and City"})
	// 	return
	// }
	jobData,err:=h.s.JobByCompanyId(ctx,jobs,compId)
	if err != nil {
		log.Error().Err(err).Str("Tracker Id", trackerId).Msg("Add Job by companyId problem")
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"msg": "Job creation failed"})
		return
	}
	c.JSON(http.StatusCreated, jobData)
}
