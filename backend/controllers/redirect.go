package controllers

import (
    "context"
    "net/http"
	"time"
    "github.com/gin-gonic/gin"
    "github.com/vi2hnu/devops-url_shortener/models"
    "github.com/vi2hnu/devops-url_shortener/database"
	"go.mongodb.org/mongo-driver/bson"
	
)



func RedirectUrl(ctx *gin.Context) {
    shortUrl := ctx.Param("shortUrl")
    mongoCtx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
    defer cancel()

    collection := database.GetCollection(database.DB, "urls")

    var result models.Url
    err := collection.FindOne(mongoCtx, bson.M{"shortened_url": shortUrl}).Decode(&result)
    if err != nil {
        ctx.JSON(http.StatusNotFound, gin.H{"error": "Shortened URL not found"})
        return
    }

    ctx.Redirect(http.StatusMovedPermanently, result.OriginalUrl)
}
