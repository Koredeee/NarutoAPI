package controllers

import (
	"NarutoAPI/models"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type ClanInput struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}

// Get All Clan godoc
// @Summary Get All Clan
// @Description Get List of Clan
// @Tags Clan
// @Produce json
// @Success 200 {object} []models.Clan
// @Router /clans [get]
func GetAllClan(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)

	var clans []models.Clan

	db.Find(&clans)

	c.JSON(http.StatusOK, gin.H{"data": clans})
}

// Create a Clan godoc
// @Summary Create Clan
// @Description Create new Clan
// @Tags Clan
// @Param Body body ClanInput true "the body to create new clan"
// @Produce json
// @Success 200 {object} models.Clan
// @Router /clans [post]
func CreateClan(c *gin.Context) {
	var input ClanInput

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	clan := models.Clan{Name: input.Name, Description: input.Description}

	db := c.MustGet("db").(*gorm.DB)

	db.Create(&clan)

	c.JSON(http.StatusOK, gin.H{"data": clan})
}

// Get a Clan godoc
// @Summary Get a Clan by Id
// @Description Get one Clan by Id
// @Tags Clan
// @Produce json
// @Param id path string true "Clan Id"
// @Success 200 {object} models.Clan
// @Router /clans/{id} [get]
func GetClanById(c *gin.Context) {
	var clan models.Clan

	db := c.MustGet("db").(*gorm.DB)

	if err := db.Where("id = ?", c.Param("id")).First(&clan).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": clan})
}

// Get Shinobi from one Clan godoc
// @Summary Get Shinobies by Clan by Id
// @Description Get all Shinobies by Clan by Id
// @Tags Clan
// @Produce json
// @Param id path string true "Clan Id"
// @Success 200 {object} []models.Shinobi
// @Router /clans/{id}/shinobies [get]
func GetShinobiesByClanId(c *gin.Context) {

	var shinobies []models.Shinobi

	db := c.MustGet("db").(*gorm.DB)

	if err := db.Where("clan_id = ?", c.Param("id")).Find(&shinobies).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": shinobies})
}

// Update a Clan godoc
// @Summary Update Clan
// @Description Update Clan by Id
// @Tags Clan
// @Produce json
// @Param id path string true "Clan Id"
// @Param Body body ClanInput true "the body to Update new clan"
// @Success 200 {object} models.Clan
// @Router /clans/{id} [patch]
func UpdateClan(c *gin.Context) {

	db := c.MustGet("db").(*gorm.DB)

	var clan models.Clan
	if err := db.Where("id = ?", c.Param("id")).First(&clan).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found"})
		return
	}

	// input validation
	var input ClanInput

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	// update validation
	var updatedInputClan models.Clan

	updatedInputClan.Name = input.Name
	updatedInputClan.Description = input.Description
	updatedInputClan.UpdatedAt = time.Now()

	db.Model(&clan).Updates(updatedInputClan)

	c.JSON(http.StatusOK, gin.H{"data": clan})
}

// Delete a Clan godoc
// @Summary Delete a Clan by Id
// @Description Delete one Clan by Id
// @Tags Clan
// @Produce json
// @Param id path string true "Clan Id"
// @Success 200 {object} map[string]boolean
// @Router /clans/{id} [delete]
func DeleteClan(c *gin.Context) {
	var clan models.Clan

	db := c.MustGet("db").(*gorm.DB)

	if err := db.Where("id = ?", c.Param("id")).First(&clan).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found"})
		return
	}

	// Delete
	db.Delete(&clan)

	c.JSON(http.StatusOK, gin.H{"data": true})
}
