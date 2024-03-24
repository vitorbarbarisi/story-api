package routers

import (
	_ "story-api/docs"
	"story-api/routers/api"

	"github.com/gin-gonic/gin"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// InitRouter initialize routing information
func InitRouter() *gin.Engine {
	router := gin.New()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	apiGroup := router.Group("/api/")
	{
		apiGroup.GET("/stories", api.GetStories)
		apiGroup.POST("/stories", api.AddStory)
		apiGroup.PUT("/stories/:id", api.UpdateStory)
	}

	return router
}
