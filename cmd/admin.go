package main

import (
	"github.com/quangdangfit/gosdk/utils/logger"

	"github.com/quangdangfit/go-admin/app"
	"github.com/quangdangfit/go-admin/app/migration"
)

func main() {
	container := app.BuildContainer()
	err := migration.CreateAdmin(container)
	if err != nil {
		logger.Error("Failed to create admin: ", err)
	}
}
