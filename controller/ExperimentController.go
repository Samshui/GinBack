package controller

import (
	"Gin/common"
	"Gin/dto"
	"Gin/model"
	"Gin/response"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"log"
	"strconv"
)

// AddExperiment 新增实验
func AddExperiment(c *gin.Context) {
	DB := common.GetDB()

	// 获取参数
	eID := c.Query("EID")
	eLabel, _ := strconv.Atoi(c.Query("ELabel"))
	eName := c.Query("EName")
	lab := c.Query("lab")
	EM, _ := strconv.Atoi(c.Query("EM"))
	EN, _ := strconv.Atoi(c.Query("EN"))
	EE, _ := strconv.Atoi(c.Query("EE"))
	Site, _ := strconv.Atoi(c.Query("site"))

	log.Printf("EM:%d, EN:%d, EE:%d\n", EM, EN, EE)

	// 检查是否已存在
	if IsExperimentExisted(eID) {
		response.Success(c, gin.H{"data": -1}, "该实验已存在")
		return
	}

	// 新建实验
	newExperiment := model.Experiment{
		Model:       gorm.Model{},
		Eid:         eID,
		ELabel:      eLabel,
		Lab:         lab,
		Name:        eName,
		TimeMorning: EM,
		TimeNoon:    EN,
		TimeEvening: EE,
		SiteSize:    Site,
	}

	DB.Create(&newExperiment)

	// 返回数据
	response.Success(c, gin.H{"data": 1, "experiment": dto.ToExperimentDto(newExperiment)}, "新增实验成功")
}

// IsExperimentExisted 查看实验是否存在
func IsExperimentExisted(eID string) bool {
	DB := common.GetDB()

	var experiment model.Experiment
	DB.Where("EID = ?", eID).First(&experiment)

	if experiment.ID != 0 {
		return true
	}
	return false
}

// DeleteExperiment 删除实验
func DeleteExperiment(c *gin.Context) {
	DB := common.GetDB()

	eID := c.Query("EID")

	var experiment model.Experiment
	DB.Where("EID = ?", eID).First(&experiment)

	if experiment.ID != 0 {
		DB.Delete(&experiment)
		response.Success(c, gin.H{"data": 1}, "删除成功")
		return
	}

	response.Success(c, gin.H{"data": -1}, "无法删除不存在的数据")
}

// GetExperimentByLabel 获取专栏实验
func GetExperimentByLabel(c *gin.Context) {
	DB := common.GetDB()

	eLabel, _ := strconv.Atoi(c.Query("ELabel"))
	log.Printf("elabel:%d", eLabel)

	var experiments []model.Experiment
	DB.Where("e_label = ?", eLabel).Find(&experiments)

	if len(experiments) == 0 {
		response.Success(c, gin.H{"data": -1}, "未搜索到相关实验")
		return
	}

	response.Success(c, gin.H{"data": 1, "experiments": experiments}, "获取到实验")
}

// GetAllExperiments 获取所有实验
func GetAllExperiments(c *gin.Context) {
	DB := common.GetDB()

	var allExperiments []model.Experiment
	DB.Order("eid").Find(&allExperiments)

	if len(allExperiments) == 0 {
		response.Success(c, gin.H{"data": -1}, "没有实验存在")
		return
	}

	response.Success(c, gin.H{"data": 1, "all": allExperiments}, "")
}

// GetExperimentByEID 获取EID的实验
func GetExperimentByEID(c *gin.Context) {
	DB := common.GetDB()

	EID := c.Query("EID")

	var experiment model.Experiment
	DB.Table("experiments").Where("eid = ?", EID).Find(&experiment)

	if experiment.ID == 0 {
		response.Success(c, gin.H{"data": -1}, "不存在该实验")
		return
	}

	response.Success(c, gin.H{"data": experiment}, "获取成功")
}
