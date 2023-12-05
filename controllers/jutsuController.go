package controllers

import (
	"NarutoAPI/models"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type JutsuInput struct {
	Name         string `json:"jutsu_name"`
	Description  string `json:"description"`
	NatureTypeID int    `json:"nature_type_id"`
}

func GetAllJutsu(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)

	var jutsus []models.Jutsu

	db.Find(&jutsus)

	// return statusOK and retrun with map
	c.JSON(http.StatusOK, gin.H{"data": jutsus})
}

func CreateJutsu(c *gin.Context) {
	// input with created struct above
	var input JutsuInput

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	jutsu := models.Jutsu{
		Name:         input.Name,
		Description:  input.Description,
		NatureTypeID: input.NatureTypeID,
	}

	db := c.MustGet("db").(*gorm.DB)

	db.Create(&jutsu)

	c.JSON(http.StatusOK, gin.H{"data": jutsu})
}

func GetJutsuById(c *gin.Context) {
	var jutsu models.Jutsu

	// get db from ginv context
	db := c.MustGet("db").(*gorm.DB)

	// checking db result with put the param id
	if err := db.Where("id = ?", c.Param("id")).First(&jutsu).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found"})
		return
	}
}

func Updatejutsu(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)

	// get jutsu if exist
	var jutsu models.Jutsu
	if err := db.Where("id = ?", c.Param("id")).First(&jutsu).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found"})
		return
	}

	var input JutsuInput

	// for checking if input is available or not
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var updatedInputJutsu models.Jutsu

	updatedInputJutsu.Name = input.Name
	updatedInputJutsu.Description = input.Description
	updatedInputJutsu.NatureTypeID = input.NatureTypeID
	updatedInputJutsu.UpdatedAt = time.Now()

	// input to db
	db.Model(&jutsu).Updates(updatedInputJutsu)

	// return status and the map
	c.JSON(
		http.StatusOK,
		gin.H{
			"message": "Record successfully updated",
			"data":    jutsu,
		},
	)
}

func DeleteJutsu(c *gin.Context) {
	var jutsu models.Jutsu

	db := c.MustGet("db").(*gorm.DB)

	if err := db.Where("id = ?", c.Param("id")).First(&jutsu).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found"})
		return
	}

	db.Delete(&jutsu)

	c.JSON(
		http.StatusOK,
		gin.H{
			"message": "Record successfully deleted",
			"data":    true,
		},
	)
}
