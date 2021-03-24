package app

import (
	"github.com/gin-gonic/gin"
	"github.com/quangdangfit/gosdk/utils/logger"
	"go.uber.org/dig"

	"github.com/quangdangfit/go-admin/app/api"
	"github.com/quangdangfit/go-admin/app/dbs"
	repoImpl "github.com/quangdangfit/go-admin/app/repositories/impl"
	"github.com/quangdangfit/go-admin/app/router"
	serviceImpl "github.com/quangdangfit/go-admin/app/services/impl"
	"github.com/quangdangfit/go-admin/pkg/jwt"
)

func BuildContainer() *dig.Container {
	container := dig.New()

	authen, err := InitAuth()
	_ = container.Provide(func() jwt.IJWTAuth {
		return authen
	})

	// Inject database
	err = dbs.Inject(container)
	if err != nil {
		logger.Error("Failed to inject database", err)
	}

	// Inject repositories
	err = repoImpl.Inject(container)
	if err != nil {
		logger.Error("Failed to inject repositories", err)
	}

	// Inject services
	err = serviceImpl.Inject(container)
	if err != nil {
		logger.Error("Failed to inject services", err)
	}

	// Inject APIs
	err = api.Inject(container)
	if err != nil {
		logger.Error("Failed to inject APIs", err)
	}

	return container
}

func InitGinEngine(container *dig.Container) *gin.Engine {
	app := gin.New()
	router.Docs(app)
	router.RegisterAPI(app, container)
	return app
}
