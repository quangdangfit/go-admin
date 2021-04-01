package main

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	_ "github.com/jinzhu/gorm/dialects/postgres"
	_ "github.com/lib/pq"
	"github.com/quangdangfit/gosdk/utils/logger"

	"github.com/quangdangfit/go-admin/app"
	"github.com/quangdangfit/go-admin/app/migration"
	"github.com/quangdangfit/go-admin/config"
	_ "github.com/quangdangfit/go-admin/docs"
)

// @title Go Admin API Documents
// @version 1.0
// @description Swagger API for Golang Admin API.

// @securityDefinitions.basic BasicAuth
// @in header
// @name Authorization

// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization

const (
	ProductionEnv = "production"
)

func main() {
	logger.Initialize(config.Config.Env == ProductionEnv)
	container := app.BuildContainer()
	engine := app.InitGinEngine(container)

	err := migration.Migrate(container)
	if err != nil {
		logger.Warn("Failed to migrate data: ", err)
	}

	server := &http.Server{
		Addr:    ":8888",
		Handler: engine,
	}

	go func() {
		// service connections
		logger.Info("Listen at:", server.Addr)
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			logger.Fatalf("Error: %s\n", err)
		}
	}()

	// Wait for interrupt signal to gracefully shutdown the server with
	// a timeout of 1 seconds.
	quit := make(chan os.Signal)
	// kill (no param) default send syscanll.SIGTERM
	// kill -2 is syscall.SIGINT
	// kill -9 is syscall. SIGKILL but can"t be catch, so don't need add it
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	logger.Info("Shutdown Server ...")

	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()
	if err := server.Shutdown(ctx); err != nil {
		logger.Fatal("Server Shutdown: ", err)
	}
	// catching ctx.Done(). timeout of 1 seconds.
	select {
	case <-ctx.Done():
		logger.Info("Timeout of 1 seconds.")
	}
	logger.Info("Server exiting")
}
