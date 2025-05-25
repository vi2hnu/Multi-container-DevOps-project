package controllers

import (
    "context"
    "net/http"
    "fmt"
    "time"
    "math/rand"
    "github.com/gin-gonic/gin"
    "github.com/vi2hnu/devops-url_shortener/models"
    "github.com/vi2hnu/devops-url_shortener/database"
	"go.mongodb.org/mongo-driver/mongo"
    "github.com/redis/go-redis/v9"

)
var rdb *redis.Client

func InitRedisClient(client *redis.Client) {
	rdb = client
}

var seededRand *rand.Rand = rand.New(rand.NewSource(time.Now().UnixNano()))
const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

func StringWithCharset(length int, charset string) string {
    b := make([]byte, length)
    for i := range b {
        b[i] = charset[seededRand.Intn(len(charset))]
    }
    return string(b)
}

func CreateNewURL(ctx *gin.Context) {
    fmt.Print("got data");
    var newUrl models.Url
    if err := ctx.ShouldBindJSON(&newUrl); err != nil {
        ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
        return
    }

    cachedOriginalUrl, err := rdb.Get(ctx, newUrl.OriginalUrl).Result()
    if err == nil {
        ctx.JSON(http.StatusCreated, gin.H{
            "original_url":  newUrl.OriginalUrl,
            "shortened_url": cachedOriginalUrl,
        })
        return
    }

    mongoCtx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
    defer cancel()

    collection := database.GetCollection(database.DB, "urls")

    const maxRetries = 5
    for i := 0; i < maxRetries; i++ {
        newUrl.ShortenedUrl = StringWithCharset(7, charset)

        res, err := collection.InsertOne(mongoCtx, newUrl)
        if err != nil {
            if mongo.IsDuplicateKeyError(err) {
                continue 
            }

            ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save URL"})
            return
        }
        rdb.Set(ctx, newUrl.OriginalUrl, newUrl.ShortenedUrl, 24*time.Hour)
        ctx.JSON(http.StatusCreated, gin.H{
            "id":            res.InsertedID,
            "original_url":  newUrl.OriginalUrl,
            "shortened_url": newUrl.ShortenedUrl,
        })
        return
    }

    ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Could not generate a unique shortened URL"})
}

