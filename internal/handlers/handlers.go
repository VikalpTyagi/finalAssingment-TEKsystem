package handlers

import (
	"finalAssing/internal/auth"
	"finalAssing/internal/middleware"
	"finalAssing/internal/repository"
	"finalAssing/internal/services"
	"fmt"
	"net/http"

	"github.com/rs/zerolog/log"

	"github.com/gin-gonic/gin"
)

func API(a *auth.Auth, c *repository.ReposStruct) *gin.Engine {

	ginEngine := gin.New()
	mid, err := middleware.NewMid(a)
	ms := services.NewStore(c)
	h := handler{
		s: ms,
		a: a,
	}

	if err != nil {
		log.Panic().Msg("middlewares not set up")
	}
	ginEngine.Use(middleware.Logger(), gin.Recovery())

	ginEngine.GET("/check", mid.Authenticate(check))
	ginEngine.POST("/api/register", h.Signup)
	ginEngine.POST("/api/login", h.Login)
	ginEngine.POST("/api/companies", h.RegisterCompany)
	ginEngine.GET("/api/companies", h.fetchListOfCompany)
	ginEngine.GET("/api/companies/:ID", h.companyById)
	ginEngine.POST("/api/companies/:ID/jobs", h.addJobsById)

	ginEngine.GET("/api/jobs/:ID", h.fetchJobById)
	ginEngine.GET("/api/companies/:ID/jobs", h.jobsByCompanyById)
	ginEngine.GET("/api/jobs", h.ViewAllJobs)

	ginEngine.POST("api/applicant",h.AcceptApplicant)

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
