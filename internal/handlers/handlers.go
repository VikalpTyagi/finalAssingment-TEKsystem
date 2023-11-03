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
	ginEngine.POST("/signup", h.Signup)
	ginEngine.POST("/login", h.Login)
	ginEngine.POST("/registerCompany", h.RegisterCompany)
	ginEngine.GET("/listCompanies", h.fetchListOfCompany)
	ginEngine.GET("/company/:ID", h.companyById)
	ginEngine.POST("/addJobs/:ID", h.addJobsById)

	ginEngine.GET("/fetchJob/:ID", h.fetchJobById)
	ginEngine.GET("/jobBycompany/:companyId", h.jobsByCompanyById)
	ginEngine.GET("/getAllJob", h.GetAllJobs)

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
