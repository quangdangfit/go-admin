package wrapper

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/quangdangfit/go-admin/pkg/errors"
	gohttp "github.com/quangdangfit/go-admin/pkg/http"
)

// constants wrapper http
const (
	DataField    = "data"
	TraceIDField = "trace_id"
	StatusField  = "status"
	CodeField    = "code"
	MessageField = "message"
)

// GinHandlerFn gin handler function
type GinHandlerFn func(c *gin.Context) gohttp.Response

// Wrap return new gin.HandlerFunc by GinHandlerFn
func Wrap(fn GinHandlerFn) gin.HandlerFunc {
	return func(c *gin.Context) {
		// handle req
		res := fn(c)

		Translate(c, res)
	}
}

// Translate gohttp.Response to response
func Translate(c *gin.Context, res gohttp.Response) {
	result := gin.H{}
	if _, ok := res.Error.(errors.CustomError); ok {
		status := int(errors.GetType(res.Error))
		result[StatusField] = status
		result[MessageField] = errors.GetMsg(status)
		result[CodeField] = errors.GetCode(status)
	}

	// get data
	if res.Data != nil {
		result[DataField] = res.Data
	}

	c.JSON(http.StatusOK, result)
}
