package app

import (
	"github.com/gin-gonic/gin"
	"github.com/quangdangfit/gocommon/logger"
	"go.uber.org/dig"

	"github.com/quangdangfit/go-admin/app/api"
	"github.com/quangdangfit/go-admin/app/dbs"
	"github.com/quangdangfit/go-admin/app/repositories"
	"github.com/quangdangfit/go-admin/app/router"
	"github.com/quangdangfit/go-admin/app/services"
	"github.com/quangdangfit/go-admin/pkg/jwt"
)

// BuildContainer build dig container
func BuildContainer() *dig.Container {
	container := dig.New()

	auth, err := InitAuth()
	_ = container.Provide(func() jwt.IJWTAuth {
		return auth
	})

	// Inject database
	err = dbs.Inject(container)
	if err != nil {
		logger.Error("Failed to inject database", err)
	}

	// Inject repositories
	err = repositories.Inject(container)
	if err != nil {
		logger.Error("Failed to inject repositories", err)
	}

	// Inject services
	err = services.Inject(container)
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

// InitGinEngine initial new gin engine
func InitGinEngine(container *dig.Container) *gin.Engine {
	app := gin.New()
	router.Docs(app)
	err := router.RegisterAPI(app, container)
	if err != nil {
		return nil
	}

	return app
}
