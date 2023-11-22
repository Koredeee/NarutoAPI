package main

import (
	"NarutoAPI/config"
	"NarutoAPI/routes"
)

func main()  {
	db := config.ConnectDatabase()


	sqlDB, _ := db.DB()
	defer sqlDB.Close()

	r := routes.Set
}