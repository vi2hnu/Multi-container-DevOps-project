package main

import (
	"context"

	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
	"github.com/vi2hnu/devops-url_shortener/controllers"
	"github.com/vi2hnu/devops-url_shortener/database"
	"github.com/vi2hnu/devops-url_shortener/routes"
)

var ctx = context.Background()


func main(){
	//init
	router:= gin.Default()
	database.DB = database.ConnectDB()
	database.CreateIndexes(database.DB)
	rdb := redis.NewClient(&redis.Options{
        Addr:     "localhost:6379",
        Password: "", 
        DB:       0,
    })
	controllers.InitRedisClient(rdb)

	
	//routes
	routes.Newurl(router)
	routes.Redirect(router)
	router.Run(":5000")
}