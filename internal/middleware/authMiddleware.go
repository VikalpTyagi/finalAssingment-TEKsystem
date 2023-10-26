package middleware

import (
	"context"
	"errors"
	"finalAssing/internal/auth"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
)

// * this for
func (mid *Mid) Authenticate(next gin.HandlerFunc) gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx := c.Request.Context()
		trackerId, ok := ctx.Value(TrackerIdKey).(string)
		if !ok {
			log.Error().Msg("tracker Id not Found!")
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": http.StatusText(http.StatusInternalServerError)})
			return
		}
		authHeader := c.Request.Header.Get("Authorization")
		tokenParts := strings.Split(authHeader, " ")

		if len(tokenParts) != 2 || strings.ToLower(tokenParts[0]) != "bearer" {
			err := errors.New("expected authorization header with: Bearer")
			log.Error().Err(err).Str("Traker Id", trackerId).Msg("token not with bearer")
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"Error": http.StatusUnauthorized})
			return
		}
		claims, err := mid.auth.ValidateToken(tokenParts[1])
		if err != nil {
			log.Error().Err(err).Str("Tracker Id", trackerId).Send()
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": http.StatusText(http.StatusUnauthorized)})
			return
		}
		ctx = context.WithValue(ctx, auth.AuthKey, claims)
		req := c.Request.WithContext(ctx)
		c.Request = req

		next(c)
	}
}