package services

import (
	"go.uber.org/dig"
)

// Inject services
func Inject(container *dig.Container) error {
	_ = container.Provide(NewAuthService)
	_ = container.Provide(NewUserService)
	_ = container.Provide(NewRoleService)
	return nil
}
