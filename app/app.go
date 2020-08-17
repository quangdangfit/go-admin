package app

import (
	"gitlab.com/quangdangfit/gocommon/utils/logger"
	"go.uber.org/dig"

	"go-admin/app/api"
	repoImpl "go-admin/app/repositories/impl"
	serviceImpl "go-admin/app/services/impl"
)

func BuildContainer() *dig.Container {
	container := dig.New()

	// Inject repositories
	err := repoImpl.Inject(container)
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
