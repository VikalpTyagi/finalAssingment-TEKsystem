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
	a auth.Auth
}

func (h *handler) Signup(c *gin.Context) {
	ctx := c.Request.Context()
	trackerId, ok := ctx.Value(middleware.TrackerIdKey).(string)
	if !ok {
		log.Error().Msg("trakerId missing from context")
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"msg": http.StatusText(http.StatusInternalServerError)})
		return
	}

	var nu models.NewUser

	err := json.NewDecoder(c.Request.Body).Decode(&nu)
	if err != nil {
		log.Error().Err(err).Str("tracker Id", trackerId)
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"msg in decoder": http.StatusText(http.StatusInternalServerError)})
		return
	}

	validate := validator.New()
	err = validate.Struct(nu)
	if err != nil {
		log.Error().Err(err).Str("tracker Id", trackerId).Send()
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"msg": "please provide Name, Email and Password"})
		return
	}
	usr, err := h.s.CreateUser(ctx, nu)
	if err != nil {
		log.Error().Err(err).Str("tracker Id", trackerId).Msg("user signup problem")
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"msg": "user signup failed"})
		return
	}
	c.JSON(http.StatusOK, usr)
}

func (h *handler) Login(c *gin.Context) {
	ctx := c.Request.Context()
	trackerId, ok := ctx.Value(middleware.TrackerIdKey).(string)
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
		log.Error().Err(err).Str("tracker Id", trackerId)
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"msg": http.StatusText(http.StatusInternalServerError)})
		return
	}
	validate := validator.New()

	err = validate.Struct(login)
	if err != nil {
		log.Error().Err(err).Str("tracker Id", trackerId).Send()
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"msg": "please provide Email and Password"})
		return
	}
	claims, err := h.s.Authenticate(ctx, login.Email, login.Password)
	if err != nil {
		log.Error().Err(err).Str("Tracker Id", trackerId).Send()
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

func (h *handler) ForgetPassword(c *gin.Context) {
	ctx := c.Request.Context()
	trackerId, ok := ctx.Value(middleware.TrackerIdKey).(string)
	if !ok {
		log.Error().Msg("traceId missing from context")
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"msg": http.StatusText(http.StatusInternalServerError)})
		return
	}
	var forgetPass models.ForgetPass
	err := json.NewDecoder(c.Request.Body).Decode(&forgetPass)
	if err != nil {
		log.Error().Err(err).Str("tracker Id", trackerId)
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"msg": http.StatusText(http.StatusInternalServerError)})
		return
	}

	validate := validator.New()
	err = validate.Struct(forgetPass)
	if err != nil {
		log.Error().Err(err).Str("Tracker Id", trackerId).Msg("validation failed for forget password")
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"Error": "Provided Invalid data"})
		return
	}

	err = h.s.VerifyEmailnDob(ctx, &forgetPass)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"Error": "Provided Invalid data"})
		return
	}
	c.JSON(http.StatusAccepted, gin.H{"Msg": "OTP sent"})
}

func (h *handler) SetNewPassword(c *gin.Context) {
	ctx := c.Request.Context()
	trackerId, ok := ctx.Value(middleware.TrackerIdKey).(string)
	if !ok {
		log.Error().Msg("traceId missing from context")
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"msg": http.StatusText(http.StatusInternalServerError)})
		return
	}
	var data models.OTPcont
	err := json.NewDecoder(c.Request.Body).Decode(&data)
	if err != nil {
		log.Error().Err(err).Str("tracker Id", trackerId)
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"msg": http.StatusText(http.StatusInternalServerError)})
		return
	}
	validate := validator.New()
	err = validate.Struct(data)
	if err != nil {
		log.Error().Err(err).Msg("validation failed for forget password")
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"Error": "Invalid data provided"})
		return
	}
	if data.Password != data.ConfirmPassword {
		log.Error().Interface("Email", data.Email).Msg("Password and confirm password not same ")
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"Error": "Password not changed"})
		return
	}

	err = h.s.VerifyOtp(ctx, &data)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"Error": "Password not changed"})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"Msg": "Password Changed"})
}
