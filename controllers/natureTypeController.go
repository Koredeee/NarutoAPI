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

// Get All NatureTypes godoc
// @Summary Get All NatureType
// @Description Get List of NatureType
// @Tags NatureType
// @Produce json
// @Success 200 {object} []models.NatureType
// @Router /nature-types [get]
func GetAllNatureType(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)

	var natureTypes []models.NatureType

	db.Find(&natureTypes)

	c.JSON(http.StatusOK, gin.H{"data": natureTypes})
}

// Create a NatureType godoc
// @Summary Create NatureType
// @Description Create new NatureType
// @Tags NatureType
// @Param Body body NatureTypeInput true "the body to create new NatureType"
// @Produce json
// @Success 200 {object} models.NatureType
// @Router /nature-types [post]
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

// Get a NatureType godoc
// @Summary Get a NatureType by Id
// @Description Get one NatureType by Id
// @Tags NatureType
// @Producde json
// @Param id path string true "NatureType Id"
// @Success 2 {object} models.NatureType
// @Router /nature-types/{id} [get]
func GetNatureTypeById(c *gin.Context) {
	var natureType models.NatureType

	db := c.MustGet("db").(*gorm.DB)

	if err := db.Where("id = ?", c.Param("id")).First(&natureType).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": natureType})
}

// Get Jutsu from one NatureType godoc
// @Summary Get Jutsus by NatureType by id
// @Description Get all Jutsus by NatureType by Id
// @Tags NatureType
// @Produce json
// @Param id path string true "NatureType Id"
// @Success 200 {object} []models.Jutsu
// @Router /nature-types/{id}/jutsus [get]
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

// Update a NatureType godoc
// @Summary Update NatureType
// @Description Update NatureType by Id
// @Tags NatureType
// @Produce json
// @Param id path string true "NatureType Id"
// @Param Body body NatureTypeInput true "the body to Update new NatureType"
// @Success  200 {object} models.NatureType
// @Router /nature-types/{id} [patch]
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

// Delete a NatureType godoc
// @Summary Delete a NatureType by Id
// @Description Delete one NatureType by Id
// @Tags NatureType
// @Produce json
// @Param  id path string true "NatureType Id"
// @Success 200 {object} map[string]boolean
// @Router /nature-types/{id} [delete]
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
