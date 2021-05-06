package controller

import (
	"Gin/common"
	"Gin/dto"
	"Gin/model"
	"Gin/response"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"strconv"
)

// AddRecord 新增选实验记录
func AddRecord(c *gin.Context) {
	DB := common.GetDB()

	EID := c.Query("EID")
	UID := c.Query("UID")
	date := c.Query("date")
	time, _ := strconv.Atoi(c.Query("time"))
	site, _ := strconv.Atoi(c.Query("site"))

	var experiment model.Experiment

	DB.Where("eid = ?", EID).First(&experiment)
	if experiment.ID == 0 {
		response.Success(c, gin.H{"data": -1}, "不存在该实验")
		return
	}

	// 判断该用户是否在当天当时段选择了实验
	var testRecords []model.Record
	DB.Table("records").
		Where("user_id = ? AND date = ? AND time = ?", UID, date, time).
		Find(&testRecords)

	if len(testRecords) != 0 {
		response.Success(c, gin.H{"data": -3, "records": testRecords}, "当前时段您存在其他实验")
		return
	}

	if IsSelected(EID, UID, date, time, site) == 1 {
		response.Success(c, gin.H{"data": -1}, "已经选过")
		return
	} else if IsSelected(EID, UID, date, time, site) == 2 {
		response.Success(c, gin.H{"data": -2}, "已经被占")
		return
	}

	var theExperiment model.Experiment
	DB.Table("experiments").Where("eid = ?", EID).First(&theExperiment)

	newRecord := model.Record{
		Model:          gorm.Model{},
		Date:           date,
		Site:           site,
		Time:           time,
		ExperimentID:   EID,
		UserID:         UID,
		ExperimentName: theExperiment.Name,
		Lab:            theExperiment.Lab,
	}

	DB.Create(&newRecord)

	// 返回数据
	response.Success(c, gin.H{"data": 1, "record": dto.ToRecordDto(newRecord)}, "新增记录成功")
}

// IsSelected 已经被选过
func IsSelected(eid string, uid string, date string, time int, site int) int {
	DB := common.GetDB()

	var record model.Record
	DB.Where("experiment_id = ? AND date = ? AND time = ? AND site = ?", eid, date, time, site).First(&record)

	if record.ID != 0 {
		if record.UserID == uid {
			return 1
		} else {
			return 2
		}
	} else {
		return 0
	}
}

// DeleteRecord 取消选实验记录
func DeleteRecord(c *gin.Context) {
	DB := common.GetDB()

	EID := c.Query("EID")
	UID := c.Query("UID")
	date := c.Query("date")
	time, _ := strconv.Atoi(c.Query("time"))
	site, _ := strconv.Atoi(c.Query("site"))

	// 存在该记录
	if IsSelected(EID, UID, date, time, site) == 0 {
		response.Success(c, gin.H{"data": -1}, "不存在该记录")
		return
	}

	var record model.Record
	DB.Where("experiment_id = ? AND user_id = ? AND date = ? AND time = ? AND site = ?", EID, UID, date, time, site).First(&record)
	DB.Delete(record)

	response.Success(c, gin.H{"data": 1}, "取消成功")
}

// GetAllRecordByEID 获取当前实验的所有记录
func GetAllRecordByEID(c *gin.Context) {
	DB := common.GetDB()

	EID := c.Query("EID")

	var records []model.Record
	DB.Where("experiment_id = ?", EID).Find(&records)

	if len(records) == 0 {
		response.Success(c, gin.H{"data": -1}, "未搜索到相关记录")
		return
	}

	response.Success(c, gin.H{"data": 1, "records": records}, "搜索成功")
}

// GetAllRecordsByStudentID 按StudentID查找所有的记录
func GetAllRecordsByStudentID(c *gin.Context) {
	DB := common.GetDB()

	StudentId := c.Query("UID")
	var records []model.Record
	DB.Table("records").Where("user_id = ?", StudentId).Find(&records)

	response.Success(c, gin.H{"data": records}, "")
}

// GetAllSiteSelected 获取当前实验所有已选的位置
func GetAllSiteSelected(c *gin.Context) {
	DB := common.GetDB()

	Eid := c.Query("EID")
	Date := c.Query("Date")
	Time, _ := strconv.Atoi(c.Query("Time"))

	var records []model.Record
	DB.Table("records").
		Where("experiment_id = ? AND date = ? AND time = ?", Eid, Date, Time).
		Find(&records)

	if len(records) == 0 {
		response.Success(c, gin.H{"data": -1}, "当前实验无选择记录")
		return
	}

	sitesSelected := make([]int, len(records))
	for i := 0; i < len(records); i++ {
		sitesSelected[i] = records[i].Site
	}

	response.Success(c, gin.H{"data": records, "array": sitesSelected}, "")
}
