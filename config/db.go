package config

import (
	"NarutoAPI/models"
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func ConnectDatabase() *gorm.DB {
	username := "root"
	password := "root123"
	host := "tcp(127.0.0.1:3306)"
	database := "database_naruto"

	dsn := fmt.Sprintf("%v:%v@%v/%v?charset=utf8mb4&parseTime=True&loc=Local", username, password, host, database)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic(err.Error())
	}

	db.AutoMigrate(&models.Clan{}, &models.Jutsu{}, &models.NatureType{}, &models.Shinobi{})

	return db
}
