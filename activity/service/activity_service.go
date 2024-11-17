package service

import (
	"awesomeProject/db"
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
	"time"
)

var connection *gorm.DB = db.GetConnection()

type ActivityRequest struct {
	Name          string    `gorm:"column:name" json:"name"`
	Cover         string    `gorm:"column:cover" json:"cover"`
	StartTime     time.Time `gorm:"column:start_time" json:"start_time"`
	EndTime       time.Time `gorm:"column:end_time" json:"end_time"`
	HoldingDate   time.Time `gorm:"column:holding_date" json:"holding_date"`
	Location      string    `gorm:"column:location" json:"location"`
	Status        int       `gorm:"column:status" json:"status"`
	Color         string    `gorm:"column:color" json:"color"`
	Opposing      string    `gorm:"column:opposing" json:"opposing"`
	OpposingColor string    `gorm:"column:opposing_color" json:"opposing_color"`
}

func newActivity(id, name, location string, startTime, endTime time.Time) *Activity {
	return &Activity{
		ID:        id,
		Name:      name,
		Location:  location,
		StartTime: startTime,
		EndTime:   endTime,
	}
}

func newActivityInfo(id, color, opposing, opposingColor string) *ActivityInfo {
	return &ActivityInfo{
		ActivityID:    id,
		Color:         color,
		Opposing:      opposing,
		OpposingColor: opposingColor,
	}
}

type Activity struct {
	ID          string    `gorm:"column:id;primary_key" json:"id"`
	Name        string    `gorm:"column:name" json:"name"`
	Cover       string    `gorm:"column:cover" json:"cover"`
	StartTime   time.Time `gorm:"column:start_time" json:"start_time"`
	EndTime     time.Time `gorm:"column:end_time" json:"end_time"`
	HoldingDate time.Time `gorm:"column:holding_date" json:"holding_date"`
	Location    string    `gorm:"column:location" json:"location"`
	Status      int       `gorm:"column:status" json:"status"`

	ActivityInfo ActivityInfo `gorm:"foreignKey:ActivityID;references:ID" json:"activity_info"`
}

func (Activity) TableName() string {
	return "rs_activity"
}

type ActivityInfo struct {
	ActivityID    string `gorm:"column:activity_id;primary_key" json:"activity_id"`
	Color         string `gorm:"column:color" json:"color"`
	Opposing      string `gorm:"column:opposing" json:"opposing"`
	OpposingColor string `gorm:"column:opposing_color" json:"opposing_color"`
}

func (ActivityInfo) TableName() string {
	return "rs_activity_info"
}

func CreateActivity(c *gin.Context) {
	var activityRequest ActivityRequest
	if err := c.ShouldBindJSON(&activityRequest); err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "无法转化为activityRequest"})
		return
	}
	currentTime := time.Now()
	id := currentTime.Format("2006-01-02 15:04:05")
	activity := newActivity(id,
		activityRequest.Name,
		activityRequest.Location,
		activityRequest.StartTime,
		activityRequest.EndTime,
	)
	result := connection.Create(activity)
	if result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "无法创建activity"})
		return
	}

	activityInfo := newActivityInfo(id, "", "", "")
	infoResult := connection.Create(activityInfo)
	if infoResult.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "无法创建activityInfo"})
		return
	}
	c.JSON(http.StatusOK, id)
}

func QueryAllActivity(c *gin.Context) {
	var activity []Activity
	result := connection.Preload("ActivityInfo").Order("holding_date desc").Find(&activity)
	if result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "无法查询所有的activity"})
		return
	}
	c.JSON(http.StatusOK, activity)

}

func QueryActivityById(c *gin.Context) {
	id := c.Param("id")
	var activity Activity
	result := connection.Preload("ActivityInfo").Where("id=?", id).Find(&activity)
	if result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": fmt.Sprintf("无法查询指定id为 %s 的activity", id)})
		return
	}
	c.JSON(http.StatusOK, activity)

}
