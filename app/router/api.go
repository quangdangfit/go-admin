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
		authAPI *api.AuthAPI,
		userAPI *api.UserAPI,
		roleAPI *api.RoleAPI,
	) error {
		jwtMiddle := middleware.UserAuthMiddleware(jwt)

		{
			r.POST("/register", wrapper.Wrap(authAPI.Register))
			r.POST("/login", wrapper.Wrap(authAPI.Login))
			r.POST("/refresh", wrapper.Wrap(authAPI.Refresh))
			r.POST("/logout", jwtMiddle, wrapper.Wrap(authAPI.Logout))
		}

		adminPath := r.Group("/admin")
		{
			adminPath.POST("/roles", roleAPI.CreateRole)
		}

		//--------------------------------API-----------------------------------
		apiPath := r.Group("/api/v1", jwtMiddle)
		{
			apiPath.GET("/users/:id", userAPI.GetByID)
			apiPath.GET("/users", wrapper.Wrap(userAPI.List))
		}
		return nil
	})

	if err != nil {
		logger.Error(err)
	}

	return err
}
