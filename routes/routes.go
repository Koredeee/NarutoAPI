package routes

import (
	"NarutoAPI/controllers"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	swaggerFiles "github.com/swaggo/files"     // swagger embed files
	ginSwagger "github.com/swaggo/gin-swagger" // gin-swagger middleware
)

func SetupRouter(db *gorm.DB) *gin.Engine {
	r := gin.Default()

	r.Use(func(ctx *gin.Context) {
		ctx.Set("db", db)
	})

	r.GET("/clans", controllers.GetAllClan)
	r.POST("/clans", controllers.CreateClan)
	r.GET("/clans/:id", controllers.GetClanById)
	r.GET("/clans/:id/shinobies", controllers.GetShinobiByClanId)
	r.PATCH("/clans/:id", controllers.UpdateClan)
	r.DELETE("/clans/:id", controllers.DeleteClan)

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	return r
}
