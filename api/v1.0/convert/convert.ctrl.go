package convert

import "github.com/gin-gonic/gin"

func pinyin(c *gin.Context) {
	c.JSON(200, "ok")
}