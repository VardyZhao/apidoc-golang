package middleware

import (
	"github.com/gin-gonic/gin"
)

func Load(r *gin.Engine) {

	r.Use(gin.Logger())

	r.Use(gin.Recovery())

	// 可以在这里添加自定义中间件
}
