package controllers

import (
	"NarutoAPI/models"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type NatureTypeInput struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	ShinobiID   int    `json:"shinobi_id"`
}

func GetAllNatureType(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)

	var natureTypes []models.NatureType

	db.Find(&natureTypes)

	c.JSON(http.StatusOK, gin.H{"data": natureTypes})
}

func CreateNatureType(c *gin.Context) {
	var input NatureTypeInput

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	natureType := models.NatureType{
		Name:        input.Name,
		Description: input.Description,
		ShinobiID:   input.ShinobiID,
	}

	db := c.MustGet("db").(*gorm.DB)

	db.Create(&natureType)

	c.JSON(http.StatusOK, gin.H{"data": natureType})
}

func GetNatureTypeById(c *gin.Context) {
	var natureType models.NatureType

	db := c.MustGet("db").(*gorm.DB)

	if err := db.Where("id = ?", c.Param("id")).First(&natureType).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": natureType})
}

func GetJutsuByNaturetypeId(c *gin.Context) {
	var jutsus []models.Jutsu

	db := c.MustGet("db").(*gorm.DB)

	if err := db.Where("nature_type_id", c.Param("id")).Find(&jutsus).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found"})
		return
	}

	// return with http response and json records
	c.JSON(http.StatusOK, gin.H{"data": jutsus})
}

func UpdateNatureType(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)

	var natureType models.NatureType
	if err := db.Where("id = ?", c.Param("id")).First(&natureType).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found"})
		return
	}

	var input NatureTypeInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// update validation
	var updatedInputNatureType models.NatureType

	updatedInputNatureType.Name = input.Name
	updatedInputNatureType.Description = input.Description
	updatedInputNatureType.ShinobiID = input.ShinobiID
	updatedInputNatureType.UpdatedAt = time.Now()

	db.Model(&natureType).Updates(updatedInputNatureType)

	c.JSON(http.StatusOK, gin.H{"data": natureType})
}

func DeleteNatureType(c *gin.Context) {
	var natureType models.NatureType

	db := c.MustGet("db").(*gorm.DB)

	if err := db.Where("id = ?", c.Param("id")).First(&natureType).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	db.Delete(&natureType)

	c.JSON(http.StatusOK, gin.H{"data": true})
}
