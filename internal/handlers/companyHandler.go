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
	traceId, ok := ctx.Value(middleware.TrackerIdKey).(string)
	if !ok {
		// If the traceId isn't found in the request, log an error and return
		log.Error().Msg("traceId missing from context")
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"msg": http.StatusText(http.StatusInternalServerError)})
		return
	}

	// Define a NewUser variable
	var newComp models.Company

	// Attempt to decode JSON from the request body into the NewUser variable
	err := json.NewDecoder(c.Request.Body).Decode(&newComp)
	if err != nil {
		// If there is an error in decoding, log the error and return
		log.Error().Err(err).Str("Trace Id", traceId)
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"msg": http.StatusText(http.StatusInternalServerError)})
		return
	}

	// Create a new validator and validate the NewUser variable
	validate := validator.New()
	err = validate.Struct(newComp)
	if err != nil {
		// If validation fails, log the error and return
		log.Error().Err(err).Str("Trace Id", traceId).Send()
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"msg": "please provide Name of company and City"})
		return
	}

	// Attempt to create the user
	Comp, err := h.s.CreateCompany(ctx, newComp)
	if err != nil {
		log.Error().Err(err).Str("Trace Id", traceId).Msg("company registration problem")
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"msg": "company Registeration failed"})
		return
	}

	// If everything goes right, respond with the created user
	c.JSON(http.StatusCreated, Comp)
}
