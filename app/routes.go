package app

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func mapRoutes() {
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	api := router.Group("/api")
	{
		api.GET("/users", func(c *gin.Context) {
			c.JSON(http.StatusOK, struct {
				Message string
				error   string
			}{
				Message: "some",
				error:   "some error",
			})
		})
	}
}
