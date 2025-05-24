package main

import (
	"context"
	"os"
	"log"
	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
	"github.com/vi2hnu/devops-url_shortener/controllers"
	"github.com/vi2hnu/devops-url_shortener/database"
	"github.com/vi2hnu/devops-url_shortener/routes"
	"github.com/joho/godotenv"
)

var ctx = context.Background()


func main(){
	//init
	router:= gin.Default()
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

	
	//routes
	routes.Newurl(router)
	routes.Redirect(router)
	router.Run(os.Getenv("PORT"))
}