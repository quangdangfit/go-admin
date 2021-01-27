package migration

import (
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"go.uber.org/dig"

	"github.com/quangdangfit/go-admin/app/dbs"
	"github.com/quangdangfit/go-admin/app/models"
	"github.com/quangdangfit/go-admin/app/repositories"
)

func CreateAdmin(container *dig.Container) error {
	return container.Invoke(func(
		userRepo repositories.IUserRepository,
		roleRepo repositories.IRoleRepository,
	) error {
		adminRole := &models.Role{Name: "admin", Description: "Admin"}
		userRole := &models.Role{Name: "user", Description: "User"}
		err := roleRepo.Create(adminRole)
		err = roleRepo.Create(userRole)
		if err != nil {
			return err
		}

		err = userRepo.Create(&models.User{
			Username: "admin",
			Password: "admin",
			Email:    "admin@admin.com",
			RoleID:   adminRole.ID,
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

		return nil
	})
}
