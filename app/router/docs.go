package router

import (
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// Docs routes
func Docs(e *gin.Engine) {
	e.GET("swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
}
