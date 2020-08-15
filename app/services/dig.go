package services

import (
	"go.uber.org/dig"
)

func Inject(container *dig.Container) error {
	_ = container.Provide(NewUserService)
	_ = container.Provide(NewRoleService)
	return nil
}
