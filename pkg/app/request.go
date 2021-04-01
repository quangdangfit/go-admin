package app

import (
	"context"
	"strings"

	"github.com/gin-gonic/gin"
)

// constants app request
const (
	prefix = "gin-go"
	// UserIDKey
	UserIDKey        = prefix + "/user-id"
	ReqBodyKey       = prefix + "/req-body"
	ResBodyKey       = prefix + "/res-body"
	LoggerReqBodyKey = prefix + "/logger-req-body"
)

// GetToken from header
func GetToken(c *gin.Context) string {
	var token string
	auth := c.GetHeader("Authorization")
	prefix := "Bearer "
	if auth != "" && strings.HasPrefix(auth, prefix) {
		token = auth[len(prefix):]
	}
	return token
}

// GetUserID get user id from context
func GetUserID(c context.Context) string {
	userID := c.Value(UserIDKey)
	if userID == nil {
		return ""
	}
	return userID.(string)
}

// SetUserID to context
func SetUserID(c *gin.Context, userID string) {
	c.Set(UserIDKey, userID)
}
