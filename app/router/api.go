package router

import (
	"github.com/gin-gonic/gin"
	"gitlab.com/quangdangfit/gocommon/utils/logger"
	"go.uber.org/dig"

	"go-admin/app/api"
)

func RegisterAPI(r *gin.Engine, container *dig.Container) error {
	err := container.Invoke(func(
		user *api.User,
		role *api.Role,
	) error {
		auth := r.Group("/auth")
		{
			auth.POST("auth/register", user.Register)
			auth.POST("auth/login", user.Login)
		}

		admin := r.Group("admin")
		{
			admin.POST("/roles", role.CreateRole)
		}

		//--------------------------------API-----------------------------------
		apiV1 := r.Group("api/v1")
		{
			apiV1.GET("/users/:uuid", user.GetUserByID)
		}
		return nil
	})

	if err != nil {
		logger.Error(err)
	}

	return err
}
