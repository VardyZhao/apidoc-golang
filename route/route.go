package route

import (
	"apidoc/controller/api"
	"github.com/gin-gonic/gin"
)

func Load(r *gin.Engine) {

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	apiDoc := r.Group("api/doc")
	{
		apiDoc.POST("/generate", api.Doc{}.Generate)
	}
}
