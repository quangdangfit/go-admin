package repositories

import (
	"go.uber.org/dig"
)

func Inject(container *dig.Container) error {
	_ = container.Provide(NewUserRepository)
	_ = container.Provide(NewRoleRepository)
	return nil
}
