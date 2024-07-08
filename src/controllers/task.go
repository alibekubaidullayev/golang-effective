package controllers

import (
	"log/slog"
	"net/http"
	"rest/db"
	"time"

	"github.com/gin-gonic/gin"
)

func AssignUserToTask(c *gin.Context) {
	var taskUser db.TaskUser
	if err := c.ShouldBindJSON(&taskUser); err != nil {
		slog.Error("Error binding JSON", "error", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	taskUser.StartDate = time.Now()
	taskUser.EndDate = nil

	if err := db.DB.Create(&taskUser).Error; err != nil {
		slog.Error("Error creating TaskUser", "error", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error assigning user to task"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"result": taskUser})
}

func EndTask(c *gin.Context) {
	id := c.Param("id")
	var task db.Task
	if err := db.DB.First(&task, id).Error; err != nil {
		slog.Error("Error finding task", "error", err)
		c.JSON(http.StatusNotFound, gin.H{"error": "Task not found"})
		return
	}

	now := time.Now()
	task.CompletedAt = &now
	task.Status = db.Completed

	if err := db.DB.Save(&task).Error; err != nil {
		slog.Error("Error ending task", "error", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error ending task"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"result": task})
}

func CalculatePayment(c *gin.Context) {
	userId := c.Query("id")

	slog.Debug(userId)

	var taskUsers []db.TaskUser

	if err := db.DB.Where("user_id = ?", userId).Find(&taskUsers).Error; err != nil {
		slog.Error("Error finding Task-User relation", "error", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error finding Task-User relation"})
		return
	}

	totalPayment := 0.0
	currentTime := time.Now()

	for _, taskUser := range taskUsers {
		endDate := taskUser.EndDate
		if endDate == nil || currentTime.Before(*endDate) {
			endDate = &currentTime
		}
		duration := endDate.Sub(taskUser.StartDate).Hours()
		totalPayment += duration * taskUser.PaymentRate
	}

	c.JSON(http.StatusOK, gin.H{"total_payment": totalPayment})
}
