package main

import (
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"learn-chinese-api/api"
	"os"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}

	port := os.Getenv("PORT")
	host := os.Getenv("HOST")
	app := gin.Default() // create gin app
	api.ApplyRoutes(app) // apply api router
	app.Run(host + ":" + port)  // listen to given port
}