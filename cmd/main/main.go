package main

import (
	"os"

	"github.com/gin-gonic/gin"
	"github.com/prathameshbhalekar/trademarkia-task/pkg/db"
	"github.com/prathameshbhalekar/trademarkia-task/pkg/routes"
)

func main() {

	// port := os.Getenv("PORT")
	dbUrl := os.Getenv("DB_URL")

	r := gin.Default()
	db.ConnectDatabase(dbUrl)

	routes.RegisterUserRoutes(r)
	routes.RegisterLikesRoutes(r)

	r.Run()
}
