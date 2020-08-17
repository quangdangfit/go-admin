package migrations

import (
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"go.uber.org/dig"

	"go-admin/app/models"
	"go-admin/app/repositories"
	"go-admin/app/schema"
	"go-admin/dbs"
)

func createAdmin(container *dig.Container) {
	container.Invoke(func(
		userRepo repositories.IUserRepository,
		roleRepo repositories.IRoleRepository,
	) error {
		role, _ := roleRepo.CreateRole(&schema.RoleBodyParam{Name: "admin", Description: "Admin"})
		userRepo.Register(&schema.Register{
			Username: "admin",
			Password: "admin",
			Email:    "admin@admin.com",
			RoleUUID: role.UUID,
		})

		return nil
	})
}

func Migrate(container *dig.Container) {
	User := models.User{}
	Role := models.Role{}

	dbs.Database.AutoMigrate(&User, &Role)
	dbs.Database.Model(&User).AddForeignKey("role_uuid", "roles(uuid)", "RESTRICT", "RESTRICT")

	createAdmin(container)
}
