package middleware

import (
	"github.com/gin-gonic/gin"

	"go-admin/app/icontext"
	"go-admin/pkg/app"
	"go-admin/pkg/jwt"
	"go-admin/pkg/logger"
)

func wrapUserAuthContext(c *gin.Context, userID string) {
	app.SetUserID(c, userID)
	ctx := icontext.NewUserID(c.Request.Context(), userID)
	ctx = logger.NewUserIDContext(ctx, userID)
	c.Request = c.Request.WithContext(ctx)
}

// User Auth Middleware
func UserAuthMiddleware(a jwt.IJWTAuth, skippers ...SkipperFunc) gin.HandlerFunc {
	return func(c *gin.Context) {
		if SkipHandler(c, skippers...) {
			c.Next()
			return
		}

		userID, err := a.ParseUserID(app.GetToken(c), false)
		if err != nil {
			app.ResError(c, err)
			return
		}
		wrapUserAuthContext(c, userID)
		c.Next()
	}
}
