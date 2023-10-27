package handlers

import (
	"finalAssing/internal/auth"
	"finalAssing/internal/middleware"
	"finalAssing/internal/services"
	"fmt"
	"net/http"

	"github.com/rs/zerolog/log"

	"github.com/gin-gonic/gin"
)

func API(a *auth.Auth, c *services.DbConnStruct) *gin.Engine {

	// Create a new Gin engine; Gin is a HTTP web framework written in Go
	ginEngine := gin.New()

	// Attempt to create new middleware with authentication
	// Here, *auth.Auth passed as a parameter will be used to set up the middleware
	mid, err := middleware.NewMid(a)
	ms := services.NewStore(c)
	h := handler{
		s: ms,
		a: a,
	}

	// If there is an error in setting up the middleware, panic and stop the application
	// then log the error message
	if err != nil {
		log.Panic().Msg("middlewares not set up")
	}

	// Attach middleware's Log function and Gin's Recovery middleware to our application
	// The Recovery middleware recovers from any panics and writes a 500 HTTP response if there was one.
	ginEngine.Use(middleware.Logger(), gin.Recovery())

	// Define a route at path "/check"
	// If it receives a GET request, it will use the m.Authenticate(check) function.
	ginEngine.GET("/check", mid.Authenticate(check))
	ginEngine.POST("/signup", h.Signup)
	ginEngine.POST("/login", h.Login)
	ginEngine.POST("/registerCompany", h.RegisterCompany)
	ginEngine.GET("/listCompanies", h.fetchListOfCompany)
	ginEngine.GET("/company/:ID", h.companyById)

	// Return the prepared Gin engine
	return ginEngine
}

func check(c *gin.Context) {
	select {
	case <-c.Request.Context().Done():
		fmt.Println("user not there")
		return
	default:
		c.JSON(http.StatusOK, gin.H{"msg": "statusOk"})
	}
}
