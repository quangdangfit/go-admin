package router

import (
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	_ "github.com/quangdangfit/go-admin/docs"
)

func Docs(e *gin.Engine) {
	e.GET("swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
}
