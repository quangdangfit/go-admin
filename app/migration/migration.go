package migration

import (
	"go.uber.org/dig"

	"github.com/quangdangfit/go-admin/app/interfaces"
	"github.com/quangdangfit/go-admin/app/models"
)

// CreateAdmin create new user role admin
func CreateAdmin(container *dig.Container) error {
	return container.Invoke(func(
		userRepo interfaces.IUserRepository,
		roleRepo interfaces.IRoleRepository,
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

// Migrate migrate to database
func Migrate(container *dig.Container) error {
	return container.Invoke(func(
		db interfaces.IDatabase,
	) error {
		User := models.User{}
		Role := models.Role{}

		db.GetInstance().AutoMigrate(&User, &Role)
		db.GetInstance().Model(&User).AddForeignKey("role_id", "roles(id)", "RESTRICT", "RESTRICT")

		return nil
	})
}
