package handlers

import (
	"encoding/json"
	"finalAssing/internal/middleware"
	"finalAssing/internal/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator"
	"github.com/rs/zerolog/log"
)

func (h *handler) RegisterCompany(c *gin.Context) {
	ctx := c.Request.Context()
	trackerId, ok := ctx.Value(middleware.TrackerIdKey).(string)
	if !ok {
		// If the traceId isn't found in the request, log an error and return
		log.Error().Msg("TrackerId missing from context")
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"msg": http.StatusText(http.StatusInternalServerError)})
		return
	}

	// Define a NewUser variable
	var newComp models.Company

	// Attempt to decode JSON from the request body into the NewUser variable
	err := json.NewDecoder(c.Request.Body).Decode(&newComp)
	if err != nil {
		// If there is an error in decoding, log the error and return
		log.Error().Err(err).Str("tracker Id", trackerId)
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"msg": http.StatusText(http.StatusInternalServerError)})
		return
	}

	// Create a new validator and validate the NewUser variable
	validate := validator.New()
	err = validate.Struct(newComp)
	if err != nil {
		// If validation fails, log the error and return
		log.Error().Err(err).Str("Tracker Id", trackerId).Send()
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"msg": "please provide Name of company and City"})
		return
	}


	Comp, err := h.s.CreateCompany(ctx, newComp)
	if err != nil {
		log.Error().Err(err).Str("Tracker Id", trackerId).Msg("company registration problem")
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"msg": "company Registeration failed"})
		return
	}
	c.JSON(http.StatusCreated, Comp)
}

func (h *handler) fetchListOfCompany(c *gin.Context){
	ctx := c.Request.Context()
	trackerId, ok := ctx.Value(middleware.TrackerIdKey).(string)
	if !ok {
		log.Error().Msg("traceId missing from context")
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"msg": http.StatusText(http.StatusInternalServerError)})
		return
	}
	listComp, err := h.s.ViewCompanies(ctx)
	if err != nil {
		log.Error().Err(err).Str("Tracker Id", trackerId)
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"msg": "problem in viewing list of companies"})
		return
	}
	c.JSON(http.StatusOK, listComp)
}

func (h *handler) companyById(c *gin.Context){
	ctx:= c.Request.Context()
	trackerId, ok := ctx.Value(middleware.TrackerIdKey).(string)
	if !ok {
		// If the traceId isn't found in the request, log an error and return
		log.Error().Msg("TrackerId missing from context")
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"msg": http.StatusText(http.StatusInternalServerError)})
		return
	}
	companyId:= c.Param("ID")
	compData,err :=h.s.FetchCompanyByID(ctx,companyId)
	if err != nil {
		log.Error().Err(err).Str("Tracker Id", trackerId)
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"msg": "problem in viewing list of company by ID"})
		return
	}
	c.JSON(http.StatusOK, compData)
}