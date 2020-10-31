package auth

import "github.com/gin-gonic/gin"

func index(c *gin.Context)  {
	c.JSON(200, "auth index page")
}


func ApplyRoutes (r *gin.RouterGroup) {
	auth := r.Group("/auth")
	{
		auth.GET("/", index)
	}
}