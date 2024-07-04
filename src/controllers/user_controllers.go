package controllers

import (
	"log/slog"
	"net/http"
	"rest/db"

	"github.com/gin-gonic/gin"
)

func readJSON(c *gin.Context, person *db.Person) bool {
	if err := c.ShouldBindJSON(person); err != nil {
		slog.Error("Error unmarshalling JSON", "error", err)
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Error unmarshalling JSON",
		})
		return false
	}
	return true
}

func ListUsers(c *gin.Context) {
	params := make(map[string]interface{})

	for key, values := range c.Request.URL.Query() {
		params[key] = values
	}

	var persons []db.Person
	if err := db.DB.Where(params).Find(&persons).Error; err != nil {
		slog.Error("Error listing users", "error", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Error listing users",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"persons": persons,
	})
}

func GetUser(c *gin.Context) {
	id := c.Param("id")
	var person db.Person

	if err := db.DB.First(&person, id).Error; err != nil {
		slog.Error("Error getting user", "id", id, "error", err)
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"person": person,
	})
}

func CreateUser(c *gin.Context) {
	slog.Debug("Creating new user")
	var person db.Person

	if !readJSON(c, &person) {
		return
	}

	if err := db.DB.Create(&person).Error; err != nil {
		slog.Error("Error creating user", "error", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error creating user"})
		return
	}
	slog.Debug("Successfully created user", "user", person)
	c.JSON(http.StatusCreated, gin.H{
		"Created User": person,
	})
}

func DeleteUser(c *gin.Context) {
	id := c.Param("id")
	var person db.Person

	if err := db.DB.First(&person, id).Error; err != nil {
		slog.Error("Error finding user for deletion", "id", id, "error", err)
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	if err := db.DB.Delete(&person).Error; err != nil {
		slog.Error("Error deleting user", "id", id, "error", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error deleting user"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"Deleted User": person,
	})
}

func UpdateUser(c *gin.Context) {
	id := c.Param("id")
	var person db.Person

	if err := db.DB.First(&person, id).Error; err != nil {
		slog.Error("Error finding user for update", "id", id, "error", err)
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	var personNew db.Person
	if !readJSON(c, &personNew) {
		return
	}

	if err := db.DB.Model(&person).Updates(&personNew).Error; err != nil {
		slog.Error("Error updating user", "id", id, "error", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error updating user"})
		return
	}

	slog.Debug("Successfully updated user", "id", id)
	c.JSON(http.StatusOK, gin.H{
		"Updated User": person,
	})
}
