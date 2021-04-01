package roles

import (
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"

	jwtMiddle "github.com/quangdangfit/go-admin/app/middleware/jwt"
	"github.com/quangdangfit/go-admin/pkg/utils"
)

// CheckAdmin middleware
func CheckAdmin() gin.HandlerFunc {
	return func(c *gin.Context) {
		var code string

		code = utils.Success
		token := c.GetHeader("Authorization")

		if token == "" {
			code = utils.InvalidParams
			c.JSON(http.StatusUnauthorized, utils.PrepareResponse(nil, "Unauthorized", code))

			c.Abort()
			return
		}

		_, err := jwtMiddle.ValidateToken(token)
		if err != nil {
			switch err.(*jwt.ValidationError).Errors {
			case jwt.ValidationErrorExpired:
				code = utils.ErrorAuthCheckTokenTimeout
			default:
				code = utils.ErrorAuthCheckTokenFail
			}
		}

		if code != utils.Success {
			c.JSON(http.StatusUnauthorized, utils.PrepareResponse(nil, "Unauthorized", code))

			c.Abort()
			return
		}

		c.Next()
	}
}
