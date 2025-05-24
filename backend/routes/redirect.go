package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/vi2hnu/devops-url_shortener/controllers"

)

func Redirect(r *gin.Engine){
	url:= r.Group("/")
	{
		url.GET("/:shortUrl",controllers.RedirectUrl)
	}
}