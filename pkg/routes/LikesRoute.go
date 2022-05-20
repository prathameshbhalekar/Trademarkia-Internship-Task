package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/prathameshbhalekar/trademarkia-task/pkg/controllers"
)

func RegisterLikesRoutes(r *gin.Engine) {

	routes := r.Group("/likes")

	routes.GET("/GetAllMatches", controllers.GetAllMatches)
}
