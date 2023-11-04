// * user related handlers endpoints
package handlers

import (
	"encoding/json"
	"finalAssing/internal/auth"
	"finalAssing/internal/middleware"
	"finalAssing/internal/models"
	"finalAssing/internal/services"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator"
	"github.com/rs/zerolog/log"
)

type handler struct {
	s services.Service
	a *auth.Auth
}

// Signup is a method for the handler struct which handles user registration
func (h *handler) Signup(c *gin.Context) {
	ctx := c.Request.Context()
	traceId, ok := ctx.Value(middleware.TrackerIdKey).(string)
	if !ok {
		// If the traceId isn't found in the request, log an error and return
		log.Error().Msg("traceId missing from context")
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"msg": http.StatusText(http.StatusInternalServerError)})
		return
	}

	// Define a NewUser variable
	var nu models.NewUser

	// Attempt to decode JSON from the request body into the NewUser variable
	err := json.NewDecoder(c.Request.Body).Decode(&nu)
	if err != nil {
		// If there is an error in decoding, log the error and return
		log.Error().Err(err).Str("tracker Id", traceId)
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"msg in decoder": http.StatusText(http.StatusInternalServerError)})
		return
	}

	// Create a new validator and validate the NewUser variable
	validate := validator.New()
	err = validate.Struct(nu)
	if err != nil {
		// If validation fails, log the error and return
		log.Error().Err(err).Str("tracker Id", traceId).Send()
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"msg": "please provide Name, Email and Password"})
		return
	}
	usr, err := h.s.CreateUser(ctx, nu)
	if err != nil {
		log.Error().Err(err).Str("tracker Id", traceId).Msg("user signup problem")
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"msg": "user signup failed"})
		return
	}
	c.JSON(http.StatusOK, usr)
}

func (h *handler) Login(c *gin.Context) {
	ctx := c.Request.Context()
	traceId, ok := ctx.Value(middleware.TrackerIdKey).(string)
	if !ok {
		log.Error().Msg("traceId missing from context")
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"msg": http.StatusText(http.StatusInternalServerError)})
		return
	}
	var login struct {
		Email    string `json:"email" validate:"required,email"`
		Password string `json:"password" validate:"required"`
	}

	err := json.NewDecoder(c.Request.Body).Decode(&login)
	if err != nil {
		log.Error().Err(err).Str("tracker Id", traceId)
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"msg": http.StatusText(http.StatusInternalServerError)})
		return
	}
	validate := validator.New()
	err = validate.Struct(login)
	if err != nil {
		log.Error().Err(err).Str("tracker Id", traceId).Send()
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"msg": "please provide Email and Password"})
		return
	}
	claims, err := h.s.Authenticate(ctx, login.Email, login.Password)
	if err != nil {
		log.Error().Err(err).Str("Tracker Id", traceId).Send()
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"msg": "login failed"})
		return
	}
	var tkn struct {
		Token string `json:"token"`
	}
	tkn.Token, err = h.a.GenerateToken(claims)
	if err != nil {
		log.Error().Err(err).Msg("generating token")
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"msg": http.StatusText(http.StatusInternalServerError)})
		return
	}
	c.JSON(http.StatusOK, tkn)

}
