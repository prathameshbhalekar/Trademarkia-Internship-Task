package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/prathameshbhalekar/trademarkia-task/pkg/db"
	"github.com/prathameshbhalekar/trademarkia-task/pkg/routes"
)

func main() {

	err := godotenv.Load("./../../pkg/env/.env")

	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	port := os.Getenv("PORT")
	dbUrl := os.Getenv("DB_URL")

	r := gin.Default()
	db.ConnectDatabase(dbUrl)

	routes.RegisterUserRoutes(r)
	routes.RegisterLikesRoutes(r)

	r.Run(port)
}
