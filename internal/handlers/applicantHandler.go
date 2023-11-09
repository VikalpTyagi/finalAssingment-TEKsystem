package handlers

import (
	"encoding/json"
	"finalAssing/internal/middleware"
	"finalAssing/internal/models"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator"
	"github.com/rs/zerolog/log"
)

func (h *handler) AcceptApplicant(c *gin.Context) {
	ctx := c.Request.Context()
	trackerId, ok := ctx.Value(middleware.TrackerIdKey).(string)
	if !ok {
		log.Error().Msg("TrackerId missing from context in Filter Apllicant")
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"msg": http.StatusText(http.StatusInternalServerError)})
		return
	}
	var newApplicant []*models.ApplicantReq

	err := json.NewDecoder(c.Request.Body).Decode(&newApplicant)
	fmt.Println("body after saving data", newApplicant)
	if err != nil {
		
		log.Error().Err(err).Str("tracker Id", trackerId).Send()
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"msg": http.StatusText(http.StatusInternalServerError)})
		return
	}
	validate := validator.New()
	for _, data := range newApplicant {
		err = validate.Struct(data)
		if err != nil {
			log.Error().Err(err).Str("tracker Id", trackerId).Interface("body", newApplicant).Send()
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"Error": "All fields are mandatory"})
			return
		}
	}

	filteredData, err := h.s.FIlterApplication(ctx, newApplicant)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"msg": http.StatusText(http.StatusInternalServerError)})
		return
	}
	c.JSON(http.StatusOK, filteredData)
}
