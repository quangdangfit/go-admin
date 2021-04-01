package api

import (
	"go.uber.org/dig"
)

// Inject apis
func Inject(container *dig.Container) error {
	_ = container.Provide(NewAuthAPI)
	_ = container.Provide(NewUserAPI)
	_ = container.Provide(NewRoleAPI)
	return nil
}
