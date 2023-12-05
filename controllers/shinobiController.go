package controllers

import (
	"NarutoAPI/models"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type ShinobiInput struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	ClanID      int    `json:"clan_id"`
}

func GetAllShinobi(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)

	var shinobies []models.Shinobi

	db.Find(&shinobies)

	c.JSON(http.StatusOK, gin.H{"data": shinobies})
}

func CreateShinobi(c *gin.Context) {
	var input ShinobiInput

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	shinobi := models.Shinobi{
		Name:        input.Name,
		Description: input.Description,
		ClanID:      input.ClanID,
	}

	db := c.MustGet("db").(*gorm.DB)

	db.Create(&shinobi)

	c.JSON(http.StatusOK, gin.H{"data": shinobi})
}

func GetShinobiById(c *gin.Context) {
	var shinobi models.Shinobi

	db := c.MustGet("db").(*gorm.DB)

	if err := db.Where("id = ?", c.Param("id")).First(&shinobi).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": shinobi})
}

func GetNatureTypeByShinobiId(c *gin.Context) {

	var natureTypes []models.NatureType

	db := c.MustGet("db").(*gorm.DB)

	if err := db.Where("shinobi_id = ?", c.Param("id")).Find(&natureTypes).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": natureTypes})
}

func UpdateShinobi(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)

	var shinobi models.Shinobi
	if err := db.Where("id = ?", c.Param("id")).First(&shinobi).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found"})
		return
	}

	// input validation
	var input ShinobiInput

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var updatedInputShinobi models.Shinobi

	updatedInputShinobi.Name = input.Name
	updatedInputShinobi.Description = input.Description
	updatedInputShinobi.ClanID = input.ClanID
	updatedInputShinobi.UpdatedAt = time.Now()

	db.Model(&shinobi).Updates(updatedInputShinobi)

	c.JSON(http.StatusOK, gin.H{"data": shinobi})
}

func DeleteShinobi(c *gin.Context) {
	var shinobi models.Shinobi

	db := c.MustGet("db").(*gorm.DB)

	if err := db.Where("id = ?", c.Param("id")).First(&shinobi).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found"})
		return
	}

	db.Delete(&shinobi)

	c.JSON(http.StatusOK, gin.H{"data": true})
}
