package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func initializeRoutes(router *gin.Engine) {
	v1 := router.Group("/api/v1")
	v1.GET("/opening", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"msg": "GET Opening",
		})
	})

	router.Run(":8080")
}
