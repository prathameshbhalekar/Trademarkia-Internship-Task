package main

import (
	"github.com/gin-gonic/gin"
	"github.com/prathameshbhalekar/trademarkia-task/pkg/db"
	"github.com/prathameshbhalekar/trademarkia-task/pkg/routes"
)

func main() {

	// port := os.Getenv("PORT")
	// dbUrl := os.Getenv("DB_URL")

	r := gin.Default()
	db.ConnectDatabase("postgres://sdproibdhkcqbn:b97f7c1b216de296f76929bbd627c78505e0ce334306a51a84e71b2fcb58a35b@ec2-3-228-235-79.compute-1.amazonaws.com:5432/d5nhe9naad517d")

	routes.RegisterUserRoutes(r)
	routes.RegisterLikesRoutes(r)

	r.Run()
}
