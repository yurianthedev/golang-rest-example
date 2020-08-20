package routing

import (
	"github.com/gin-gonic/gin"

	"github.com/yurianxdev/rest-example/api/controller"
)

func UserRouting(r *gin.RouterGroup) {
	r.GET("/users", controller.GETUsers)
	r.GET("/users/:id")
	r.POST("/users", controller.POSTUser)
	r.PUT("/users/:id")
	r.DELETE("/users/:id")
}
