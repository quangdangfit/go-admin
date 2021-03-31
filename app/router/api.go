package router

import (
	"github.com/gin-gonic/gin"
	"github.com/quangdangfit/gosdk/utils/logger"
	"go.uber.org/dig"

	"github.com/quangdangfit/go-admin/app/api"
	"github.com/quangdangfit/go-admin/app/middleware"
	"github.com/quangdangfit/go-admin/pkg/http/wrapper"
	"github.com/quangdangfit/go-admin/pkg/jwt"
)

func RegisterAPI(r *gin.Engine, container *dig.Container) error {
	err := container.Invoke(func(
		jwt jwt.IJWTAuth,
		auth *api.Auth,
		user *api.UserAPI,
		role *api.RoleAPI,
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
		api := r.Group("/api/v1", middleware.UserAuthMiddleware(jwt))
		{
			api.GET("/users/:id", user.GetByID)
			api.GET("/users", wrapper.Wrap(user.List))
		}
		return nil
	})

	if err != nil {
		logger.Error(err)
	}

	return err
}
