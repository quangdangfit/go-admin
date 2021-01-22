package migration

import (
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"go.uber.org/dig"

	"github.com/quangdangfit/go-admin/app/dbs"
	"github.com/quangdangfit/go-admin/app/models"
	"github.com/quangdangfit/go-admin/app/repositories"
	"github.com/quangdangfit/go-admin/app/schema"
)

func createAdmin(container *dig.Container) error {
	return container.Invoke(func(
		userRepo repositories.IUserRepository,
		roleRepo repositories.IRoleRepository,
	) error {
		admin, err := roleRepo.CreateRole(&schema.RoleBodyParam{Name: "admin", Description: "Admin"})
		_, err = roleRepo.CreateRole(&schema.RoleBodyParam{Name: "user", Description: "User"})
		if err != nil {
			return err
		}

		_, err = userRepo.Register(&schema.RegisterBodyParam{
			Username: "admin",
			Password: "admin",
			Email:    "admin@admin.com",
			RoleID:   admin.ID,
		})
		if err != nil {
			return err
		}

		return nil
	})
}

func Migrate(container *dig.Container) error {
	return container.Invoke(func(
		db dbs.IDatabase,
	) error {
		User := models.User{}
		Role := models.Role{}

		db.GetInstance().AutoMigrate(&User, &Role)
		db.GetInstance().Model(&User).AddForeignKey("role_id", "roles(id)", "RESTRICT", "RESTRICT")

		err := createAdmin(container)
		if err != nil {
			return err
		}

		return nil
	})
}
