package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/vi2hnu/devops-url_shortener/controllers"

)

func Newurl(r *gin.Engine){
	url:= r.Group("/create")
	{
		url.POST("/",controllers.CreateNewURL)
	}
}