package controllers

import (
	"fmt"
	"log/slog"
	"net/http"
	"reflect"
	"rest/db"
	"strconv"

	"github.com/gin-gonic/gin"
)

func List(c *gin.Context, t reflect.Type) {
	params := make(map[string]interface{})

	for key, values := range c.Request.URL.Query() {
		if key == "page" || key == "page_size" {
			continue
		}
		params[key] = values
	}

	objects := reflect.MakeSlice(reflect.SliceOf(t), 0, 0).Interface()

	page, err := strconv.Atoi(c.DefaultQuery("page", "1"))
	if err != nil || page < 1 {
		page = 1
	}
	pageSize, err := strconv.Atoi(c.DefaultQuery("page_size", "10"))
	if err != nil || pageSize < 1 {
		pageSize = 10
	}
	offset := (page - 1) * pageSize

	query := db.DB.Where(params).Offset(offset).Limit(pageSize).Find(&objects)

	if err := query.Error; err != nil {
		slog.Error("Error listing objects", "type", t.Name(), "error", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": fmt.Sprintf("Error listing %s", t.Name()),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"result": objects,
	})
}

func Get(c *gin.Context, t reflect.Type) {
	id := c.Param("id")
	object := reflect.New(t).Interface()

	if err := db.DB.First(object, id).Error; err != nil {
		slog.Error("Error getting object", "type", t.Name(), "id", id, "error", err)
		c.JSON(http.StatusNotFound, gin.H{"error": fmt.Sprintf("%s not found", t.Name())})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"result": object,
	})
}

func Create(c *gin.Context, t reflect.Type) {
	slog.Debug(fmt.Sprintf("Creating new %s", t.Name()))
	object := reflect.New(t).Interface()

	if !readJSON(c, object) {
		return
	}

	if err := db.DB.Create(object).Error; err != nil {
		slog.Error("Error creating object", "type", t.Name(), "error", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Error creating %s", t.Name())})
		return
	}
	slog.Debug(fmt.Sprintf("Successfully created %s", t.Name()), "object", object)
	c.JSON(http.StatusCreated, gin.H{
		"result": object,
	})
}

func Delete(c *gin.Context, t reflect.Type) {
	id := c.Param("id")
	object := reflect.New(t).Interface()

	if err := db.DB.First(object, id).Error; err != nil {
		slog.Error("Error finding object for deletion", "type", t.Name(), "id", id, "error", err)
		c.JSON(http.StatusNotFound, gin.H{"error": fmt.Sprintf("%s not found", t.Name())})
		return
	}

	if err := db.DB.Delete(object).Error; err != nil {
		slog.Error("Error deleting object", "type", t.Name(), "id", id, "error", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Error deleting %s", t.Name())})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"result": object,
	})
}

func Update(c *gin.Context, t reflect.Type) {
	id := c.Param("id")
	object := reflect.New(t).Interface()

	if err := db.DB.First(object, id).Error; err != nil {
		slog.Error("Error finding object for update", "type", t.Name(), "id", id, "error", err)
		c.JSON(http.StatusNotFound, gin.H{"error": fmt.Sprintf("%s not found", t.Name())})
		return
	}

	updatedObject := reflect.New(t).Interface()
	if !readJSON(c, updatedObject) {
		return
	}

	if err := db.DB.Model(object).Updates(updatedObject).Error; err != nil {
		slog.Error("Error updating object", "type", t.Name(), "id", id, "error", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Error updating %s", t.Name())})
		return
	}

	slog.Debug(fmt.Sprintf("Successfully updated %s", t.Name()), "id", id)
	c.JSON(http.StatusOK, gin.H{
		"result": object,
	})
}
