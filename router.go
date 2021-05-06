package main

import (
	"Gin/controller"
	"Gin/middleware"
	"github.com/gin-gonic/gin"
)

func CollectRouter(r *gin.Engine) *gin.Engine {
	// User
	r.POST("/api/user/register", controller.Register)
	r.POST("/api/user/login", controller.Login)
	r.GET("/api/user/info", middleware.AuthMiddleware(), controller.Info)
	r.POST("/api/user/changeTelephone", controller.ChangeTelephone)
	r.POST("/api/user/telephoneIsExist", controller.TelephoneIsExisted)

	// Experiment
	r.POST("/api/experiment/add", controller.AddExperiment)
	r.POST("/api/experiment/delete", controller.DeleteExperiment)
	r.POST("/api/experiment/get", controller.GetExperimentByLabel)
	r.POST("/api/experiment/all", controller.GetAllExperiments)

	// Record
	r.POST("/api/record/add", controller.AddRecord)
	r.POST("/api/record/delete", controller.DeleteRecord)
	r.POST("/api/record/getAll", controller.GetAllRecordByEID)
	r.POST("/api/record/getRecordsSites", controller.GetAllSiteSelected)
	r.POST("/api/record/getUserRecords", controller.GetAllRecordsByStudentID)

	return r
}
