package apiv1

import (
	"github.com/gin-gonic/gin"
	"learn-chinese-api/api/v1.0/auth"
)

func ApplyRoutes(r *gin.RouterGroup) {
	apiv1 := r.Group("/v1.0")
	{
		auth.ApplyRoutes(apiv1)
	}
}
