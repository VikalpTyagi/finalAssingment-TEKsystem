package handlers

import (
	"finalAssing/internal/auth"
	"finalAssing/internal/cacheier"
	"finalAssing/internal/middleware"
	"finalAssing/internal/repository"
	"finalAssing/internal/services"
	"fmt"
	"net/http"

	"github.com/redis/go-redis/v9"
	"github.com/rs/zerolog/log"

	"github.com/gin-gonic/gin"
)

func API(a auth.Auth, c repository.RepoInterface, red *redis.Client) *gin.Engine {

	ginEngine := gin.New()
	mid, err := middleware.NewMid(a)
	if err != nil {
		panic(err) 
	}
	redConn,err := cacheier.NewRedConn(red)
	if err != nil {
		panic(err) 
	}
	ms := services.NewStore(c, redConn)
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
	ginEngine.POST("/api/companies",mid.Authenticate( h.RegisterCompany))
	ginEngine.GET("/api/companies",mid.Authenticate( h.fetchListOfCompany))
	ginEngine.GET("/api/companies/:ID",mid.Authenticate( h.companyById))
	ginEngine.POST("/api/companies/:ID/jobs",mid.Authenticate( h.addJobsById))

	ginEngine.GET("/api/jobs/:ID",mid.Authenticate( h.fetchJobById))
	ginEngine.GET("/api/companies/:ID/jobs",mid.Authenticate( h.jobsByCompanyById))
	ginEngine.GET("/api/jobs",mid.Authenticate( h.ViewAllJobs))

	ginEngine.POST("api/applicant",mid.Authenticate( h.AcceptApplicant))

	ginEngine.POST("api/forgetPassword",h.ForgetPassword)
	ginEngine.POST("api/updatePassword",h.SetNewPassword)

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
