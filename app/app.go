package app

import (
	"gitlab.com/quangdangfit/gocommon/utils/logger"
	"go.uber.org/dig"

	"github.com/quangdangfit/go-admin/app/api"
	repoImpl "github.com/quangdangfit/go-admin/app/repositories/impl"
	serviceImpl "github.com/quangdangfit/go-admin/app/services/impl"
	"github.com/quangdangfit/go-admin/pkg/jwt"
)

func BuildContainer() *dig.Container {
	container := dig.New()

	authen, err := InitAuth()
	_ = container.Provide(func() jwt.IJWTAuth {
		return authen
	})

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
