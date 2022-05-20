package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/prathameshbhalekar/trademarkia-task/pkg/controllers"
)

func RegisterUserRoutes(r *gin.Engine) {

	routes := r.Group("/users")

	routes.GET("/FindUsersInRange", controllers.FindInRange)
	routes.GET("/FindUsersByName", controllers.FindUsersByName)
}
