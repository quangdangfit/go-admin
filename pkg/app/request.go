package app

import (
	"context"
	"strings"

	"github.com/gin-gonic/gin"
)

const (
	prefix = "gin-go"
	// UserIDKey
	UserIDKey        = prefix + "/user-id"
	ReqBodyKey       = prefix + "/req-body"
	ResBodyKey       = prefix + "/res-body"
	LoggerReqBodyKey = prefix + "/logger-req-body"
)

func GetToken(c *gin.Context) string {
	var token string
	auth := c.GetHeader("Authorization")
	prefix := "Bearer "
	if auth != "" && strings.HasPrefix(auth, prefix) {
		token = auth[len(prefix):]
	}
	return token
}

func GetUserID(c context.Context) string {
	userId := c.Value(UserIDKey)
	if userId == nil {
		return ""
	}
	return userId.(string)
}

// SetUserID
func SetUserID(c *gin.Context, userID string) {
	c.Set(UserIDKey, userID)
}
