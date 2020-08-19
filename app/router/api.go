package router

import (
	"github.com/gin-gonic/gin"
	"gitlab.com/quangdangfit/gocommon/utils/logger"
	"go.uber.org/dig"

	"go-admin/app/api"
	"go-admin/app/middleware"
	"go-admin/pkg/http/wrapper"
	"go-admin/pkg/jwt"
)

func RegisterAPI(r *gin.Engine, container *dig.Container) error {
	err := container.Invoke(func(
		jwt jwt.IJWTAuth,
		auth *api.Auth,
		user *api.User,
		role *api.Role,
	) error {
		{
			r.POST("/register", wrapper.Wrap(auth.Register))
			r.POST("/login", wrapper.Wrap(auth.Login))
			r.POST("/refresh", wrapper.Wrap(auth.Refresh))
			r.POST("/logout", middleware.UserAuthMiddleware(jwt), wrapper.Wrap(auth.Logout))
		}

		admin := r.Group("/admin")
		{
			admin.POST("/roles", role.CreateRole)
		}

		//--------------------------------API-----------------------------------
		api := r.Group("/api/v1")
		{
			api.GET("/users/:id", user.GetUserByID)
			api.GET("/users", user.List)
		}
		return nil
	})

	if err != nil {
		logger.Error(err)
	}

	return err
}
