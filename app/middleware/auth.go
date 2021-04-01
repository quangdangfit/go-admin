package middleware

import (
	"github.com/gin-gonic/gin"

	"github.com/quangdangfit/go-admin/pkg/app"
	gohttp "github.com/quangdangfit/go-admin/pkg/http"
	"github.com/quangdangfit/go-admin/pkg/http/wrapper"
	"github.com/quangdangfit/go-admin/pkg/jwt"
)

func wrapUserAuthContext(c *gin.Context, userID string) {
	app.SetUserID(c, userID)
	c.Request = c.Request.WithContext(c)
}

// UserAuthMiddleware User Auth Middleware
func UserAuthMiddleware(a jwt.IJWTAuth, skippers ...SkipperFunc) gin.HandlerFunc {
	return func(c *gin.Context) {
		if SkipHandler(c, skippers...) {
			c.Next()
			return
		}

		userID, err := a.ParseUserID(app.GetToken(c), false)
		if err != nil {
			wrapper.Translate(c, gohttp.Response{Error: err})
			c.Abort()
			return
		}
		wrapUserAuthContext(c, userID)
		c.Next()
	}
}
