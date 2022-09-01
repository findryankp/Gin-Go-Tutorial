package routes

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	"go-rest-api-gin/controllers"

	swaggerFiles "github.com/swaggo/files"     // swagger embed files
	ginSwagger "github.com/swaggo/gin-swagger" // gin-swagger middleware
)

func SetupRouter(db *gorm.DB) *gin.Engine {
	r := gin.Default()

	// set db to gin context
	r.Use(func(c *gin.Context) {
		c.Set("db", db)
	})
	r.GET("/movies", controllers.GetAllMovie)
	r.POST("/movies", controllers.CreateMovie)
	r.GET("/movies/:id", controllers.GetMovieById)
	r.PATCH("/movies/:id", controllers.UpdateMovie)
	r.DELETE("movies/:id", controllers.DeleteMovie)

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	return r
}
