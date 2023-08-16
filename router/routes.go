package router

import (
	docs "github.com/fowkeio/secjobs/docs"
	"github.com/fowkeio/secjobs/handler"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func initializeRoutes(router *gin.Engine) {
	// Initialize handler
	handler.InitializeHandler()
	basePath := "/api/v1"
	docs.SwaggerInfo.BasePath = basePath

	v1 := router.Group(basePath)

	v1.GET("/opening", handler.ShowOpeningHandler)

	v1.POST("/opening", handler.CreateOpeningHandler)

	v1.DELETE("/opening", handler.DeleteOpeningHandler)

	v1.PUT("/opening", handler.UpdateOpeningHandler)

	v1.GET("/openings", handler.ListOpeninsgHandler)

	router.GET("swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	router.Run(":8080")

}

// Initialize Swagger
