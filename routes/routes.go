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
	r.GET("/clans/:id/shinobies", controllers.GetShinobiesByClanId)
	r.PATCH("/clans/:id", controllers.UpdateClan)
	r.DELETE("/clans/:id", controllers.DeleteClan)

	r.GET("/jutsus", controllers.GetAllJutsu)
	r.POST("/jutsus", controllers.CreateJutsu)
	r.GET("/jutsus/:id", controllers.GetJutsuById)
	r.PATCH("/jutsus/:id", controllers.UpdateJutsu)
	r.DELETE("/jutsus/:id", controllers.DeleteJutsu)

	r.GET("/nature-types", controllers.GetAllNatureType)
	r.POST("/nature-types", controllers.CreateNatureType)
	r.GET("/nature-types/:id", controllers.GetNatureTypeById)
	r.GET("/nature-types/:id/jutsus", controllers.GetJutsuByNaturetypeId)
	r.PATCH("/nature-types/:id", controllers.UpdateNatureType)
	r.DELETE("/nature-types/:id", controllers.DeleteNatureType)

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	return r
}
