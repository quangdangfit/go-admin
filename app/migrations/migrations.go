package migrations

import (
	_ "github.com/jinzhu/gorm/dialects/postgres"

	"go-admin/app/models"
	"go-admin/app/repositories"
	"go-admin/app/schema"
	"go-admin/dbs"
)

func createAdmin() {
	roleRepo := repositories.NewRoleRepository()
	role, _ := roleRepo.CreateRole(&schema.RoleBodyParam{Name: "admin", Description: "Admin"})

	userRepo := repositories.NewUserRepository()
	userRepo.Register(&schema.Register{
		Username: "admin",
		Password: "admin",
		Email:    "admin@admin.com",
		RoleUUID: role.UUID,
	})
}

func Migrate() {
	User := models.User{}
	Role := models.Role{}

	dbs.Database.AutoMigrate(&User, &Role)
	dbs.Database.Model(&User).AddForeignKey("role_uuid", "roles(uuid)", "RESTRICT", "RESTRICT")

	//createAdmin()
}
