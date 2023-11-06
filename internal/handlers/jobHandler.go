// * job related handlers endpoint
package handlers

import (
	"encoding/json"
	"finalAssing/internal/middleware"
	"finalAssing/internal/models"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
)

func (h *handler) addJobsById(c *gin.Context) {
	ctx := c.Request.Context()
	trackerId, ok := ctx.Value(middleware.TrackerIdKey).(string)
	if !ok {
		log.Error().Msg("trackerId missing from context")
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"msg": http.StatusText(http.StatusInternalServerError)})
		return
	}
	compId := c.Param("ID")
	var jobs []models.Job
	err := json.NewDecoder(c.Request.Body).Decode(&jobs)
	if err != nil {
		log.Error().Err(err).Str("tracker Id", trackerId)
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"Error": http.StatusText(http.StatusInternalServerError)})
		fmt.Println("============Body oarsing failed",c.Request.Body)
		return
	}
	jobData, err := h.s.JobByCompanyId(jobs, compId)
	if err != nil {
		log.Error().Err(err).Str("tracker Id", trackerId).Msg("Add Job by companyId problem")
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"msg": "Job creation failed"})
		return
	}
	c.JSON(http.StatusCreated, jobData)
}

func (h *handler) jobsByCompanyById(c *gin.Context) {
	ctx := c.Request.Context()
	trackerId, ok := ctx.Value(middleware.TrackerIdKey).(string)
	if !ok {
		log.Error().Msg("trackerId missing from context")
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"msg": http.StatusText(http.StatusInternalServerError)})
		return
	}
	companyId := c.Param("companyId")
	listOfJobs, err := h.s.FetchJobByCompanyId(ctx, companyId)
	if err != nil {
		log.Error().Err(err).Str("tracker Id", trackerId)
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"msg": "problem in viewing list of company by ID"})
		return
	}
	c.JSON(http.StatusOK, listOfJobs)
}

func (h *handler) fetchJobById(c *gin.Context) {
	ctx := c.Request.Context()
	trackerId, ok := ctx.Value(middleware.TrackerIdKey).(string)
	if !ok {
		log.Error().Msg("TrackerId missing from context")
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"msg": http.StatusText(http.StatusInternalServerError)})
		return
	}
	jobId := c.Param("ID")
	job, err := h.s.GetJobById(ctx, jobId)
	if err != nil {
		log.Error().Err(err).Str("Tracker Id", trackerId)
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"msg": http.StatusText(http.StatusBadRequest)})
		return
	}
	c.JSON(http.StatusOK, job)
}

func (h *handler) GetAllJobs(c *gin.Context) {
	ctx := c.Request.Context()
	trackerId, ok := ctx.Value(middleware.TrackerIdKey).(string)
	if !ok {
		log.Error().Msg("TrackerId missing from context")
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"msg": http.StatusText(http.StatusInternalServerError)})
		return
	}
	job, err := h.s.GetAllJobs(ctx)
	if err != nil {
		log.Error().Err(err).Str("Tracker Id", trackerId)
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"msg": "problem in viewing list of company by ID"})
		return
	}
	c.JSON(http.StatusOK, job)
}
