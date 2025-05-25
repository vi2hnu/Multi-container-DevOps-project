package main

import (
	"os"
	"log"
	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
	"github.com/vi2hnu/devops-url_shortener/controllers"
	"github.com/vi2hnu/devops-url_shortener/database"
	"github.com/vi2hnu/devops-url_shortener/routes"
	"github.com/joho/godotenv"
	"github.com/gin-contrib/cors"
)

func main() {
	// init
	router := gin.Default()

	// CORS middleware
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:3000","http://192.168.29.117:3000"},
		AllowMethods:     []string{"GET", "POST", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	}))

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	database.DB = database.ConnectDB()
	database.CreateIndexes(database.DB)

	rdb := redis.NewClient(&redis.Options{
		Addr:     os.Getenv("REDIS_PORT"),
		Password: os.Getenv("REDIS_PASSWORD"),
		DB:       0,
	})
	controllers.InitRedisClient(rdb)

	// routes
	routes.Newurl(router)
	routes.Redirect(router)

	router.Run(os.Getenv("PORT"))
}
