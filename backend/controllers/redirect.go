package controllers

import (
    "context"
    "net/http"
	"time"
    "github.com/gin-gonic/gin"
    "github.com/vi2hnu/devops-url_shortener/models"
    "github.com/vi2hnu/devops-url_shortener/database"
	"go.mongodb.org/mongo-driver/bson"
    "github.com/redis/go-redis/v9"
	
)

var rdb *redis.Client

func InitRedisClient(client *redis.Client) {
	rdb = client
}


func RedirectUrl(ctx *gin.Context) {
    shortUrl := ctx.Param("shortUrl")
    mongoCtx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
    defer cancel()
    cachedOriginalUrl, err := rdb.Get(ctx, shortUrl).Result()
    if err == nil {
        // fmt.Println("Cache hit, redirecting from Redis")
        ctx.Redirect(http.StatusMovedPermanently, cachedOriginalUrl)
        return
    }

    collection := database.GetCollection(database.DB, "urls")

    var result models.Url
    err = collection.FindOne(mongoCtx, bson.M{"shortened_url": shortUrl}).Decode(&result)
    if err != nil {
        ctx.JSON(http.StatusNotFound, gin.H{"error": "Shortened URL not found"})
        return
    }
    rdb.Set(ctx, shortUrl, result.OriginalUrl, 24*time.Hour)

    ctx.Redirect(http.StatusMovedPermanently, result.OriginalUrl)
}
